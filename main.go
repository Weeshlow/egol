package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"runtime"
	"syscall"
	"time"

	"github.com/streamrail/concurrent-map"
	log "github.com/unchartedsoftware/plog"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/unchartedsoftware/egol/api/conf"
	"github.com/unchartedsoftware/egol/api/middleware"
	"github.com/unchartedsoftware/egol/api/redis"
	"github.com/unchartedsoftware/egol/api/sim"
	"github.com/unchartedsoftware/egol/api/util"
	"github.com/unchartedsoftware/egol/api/ws"
)

const (
	websocketRoute = "/connect"
	numFamilyTypes = 5
	organismCount  = 50
)

var (
	exit      = make(chan bool)
	organisms map[string]*sim.Organism
	redisConn *redis.Connection
	clients   cmap.ConcurrentMap
	config    *conf.Conf
	families  []*sim.Attributes
)

func handleMessage(client *ws.Client) ws.RequestHandler {
	return func(conn *ws.Connection, msg []byte) {
		// parse the tile request
		_, err := ws.NewMessage(msg)
		if err != nil {
			// parsing error, send back a failure response
			err := fmt.Errorf("Unable to parse message: %s", string(msg))
			// log error
			log.Warn(err)
			err = conn.Send(&ws.Message{
				Success: false,
			})
			if err != nil {
				log.Warn(err)
			}
			return
		}
		// send response
		err = conn.Send(&ws.Message{
			Success: true,
		})
		if err != nil {
			log.Warn(err)
		}
	}
}

func shouldExit() bool {
	select {
	case <-exit:
		return true
	default:
		// non blocking
		return false
	}
}

func initializeSim() {
	families = make([]*sim.Attributes, numFamilyTypes)
	organisms = make(map[string]*sim.Organism)
	for i := 0; i < numFamilyTypes; i++ {
		families[i] = &sim.Attributes{
			Family:         uint32(i),
			Offense:        0.01 + (rand.Float64() * 0.02),
			Defense:        0.01 + (rand.Float64() * 0.02),
			Agility:        0.01 + (rand.Float64() * 0.02),
			Reproductivity: math.Min(0.1, math.Max(0.9, rand.Float64())),
			// coordniate based
			Speed:      0.01 + (rand.Float64() * 0.05),
			Range:      0.01 + (rand.Float64() * 0.03),
			Perception: 0.1 + (rand.Float64() * 0.1),
		}
	}
	// Initialize organisms. Add random variation from family
	for i := 0; i < organismCount; i++ {
		family := families[i%numFamilyTypes]
		organism := sim.NewOrganism(family)
		organisms[organism.ID] = organism
	}
}

func store(suffix string, iteration int64, data interface{}) error {
	key := fmt.Sprintf("%s-%d-%s", config.SimID, iteration, suffix)
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = redisConn.Set(key, bytes)
	if err != nil {
		return err
	}
	return nil
}

func loop() {
	iteration := int64(0)

	for {
		// get timestamp
		stamp := util.Timestamp()

		// check if should exit
		if shouldExit() {
			break
		}

		// apoply constraints and determine AI input for each organism
		updates := sim.Iterate(organisms)

		//Every 20 iterations spawn new organisms around edge
		if iteration % 30 == 0 {
			count := rand.Intn(organismCount / 2)

			for i := 0; i < count; i++ {
				family := families[i%numFamilyTypes]
				organism := sim.NewOrganism(family)

				x := rand.Float64();
				y := rand.Float64();
				if (rand.Float64() < 0.5) {
					if x < 0.5 {
						x = 0 + - rand.Float64()*0.02;
					} else {
						x = 1 - rand.Float64()*0.02;
					}
				} else {
					if y < 0.5 {
						y = 0 + - rand.Float64()*0.02;
					} else {
						y = 1 - rand.Float64()*0.02;
					}	
				}
				organism.State.Position = mgl32.Vec3{
					float32(x),
					float32(y),
					0.0,
				}
				updates[organism.ID] = &sim.Update{
					ID:         organism.ID,
					State:      organism.State,
					Attributes: organism.Attributes,
				}
			}
		}

		// apply updates to the state before next iteration
		for key, update := range updates {
			if organisms[key] == nil {
				organisms[key] = &sim.Organism{
					ID:         update.ID,
					State:      update.State,
					Attributes: update.Attributes,
				}
			} else {
				organisms[key].Update(update)
			}
		}

		// write out current state
		err := store("state", iteration, organisms)
		if err != nil {
			log.Error(err)
			time.Sleep(time.Duration(1000) * time.Millisecond)
			continue
		}

		// write out updates
		err = store("update", iteration, updates)
		if err != nil {
			log.Error(err)
			time.Sleep(time.Duration(1000) * time.Millisecond)
			continue
		}

		log.Info("Iteration: ", iteration)

		for iter := range clients.Iter() {
			client, ok := iter.Val.(*ws.Client)
			if !ok {
				log.Error(err)
				continue
			}
			// broadcast update to connected clients
			if !client.New {
				err := client.Conn.Send(&ws.Message{
					Type:    "update",
					Data:    updates,
					Success: true,
				})
				if err != nil {
					log.Error(err)
				}
			}
			// broadcast state to new clients
			if client.New {
				err := client.Conn.Send(&ws.Message{
					Type:    "state",
					Data:    organisms,
					Success: true,
				})
				if err != nil {
					log.Error(err)
				}
				client.New = false
			}
		}

		// wait
		now := util.Timestamp()
		elapsed := now - stamp
		if elapsed < config.FrameMS {
			time.Sleep(time.Duration(config.FrameMS-elapsed) * time.Millisecond)
		}

		// increment the iteration
		iteration++
	}
	exit <- true
}

func onWSConnection(w http.ResponseWriter, r *http.Request) {
	log.Info("New connection")
	// create client
	client := ws.NewClient()
	// create connection
	conn, err := ws.NewConnection(w, r, handleMessage(client))
	if err != nil {
		log.Warn(err)
		return
	}
	client.Conn = conn
	clients.Set(client.ID, client)
	// listen for requests and respond
	err = conn.ListenAndRespond()
	if err != nil {
		log.Info(err)
	}
	// clean up internals
	conn.Close()
	clients.Remove(client.ID)
	log.Info("Connection lost")
}

func server() {
	// create server
	app := web.New()
	// mount logger middleware
	app.Use(middleware.Log)
	// mount gzip middleware
	app.Use(middleware.Gzip)
	// add websocket route
	app.Get(websocketRoute, onWSConnection)
	// add greedy route last
	app.Get("/*", http.FileServer(http.Dir(config.PublicDir)))
	// catch kill signals for graceful shutdown
	graceful.AddSignal(syscall.SIGINT, syscall.SIGTERM)
	// start server
	err := graceful.ListenAndServe(":"+config.Port, app)
	if err != nil {
		log.Error(err)
	}
	// exit loop
	exit <- true
	// wait for acknowledgement
	<-exit
}

func main() {

	rand.Seed(time.Now().UnixNano())

	// sets the maximum number of CPUs that can be executing simultaneously
	runtime.GOMAXPROCS(runtime.NumCPU())

	// parse flags into config struct
	config = conf.ParseCommandLine()

	// initialize the sim
	initializeSim()

	// get redis connection
	redisConn = redis.NewConnection(config.RedisHost, config.RedisPort, 0)

	// create clients map
	clients = cmap.New()

	// start server
	go loop()

	server()
}
