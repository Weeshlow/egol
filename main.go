package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"syscall"
	"encoding/json"
	"time"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/streamrail/concurrent-map"
	log "github.com/unchartedsoftware/plog"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"

	"github.com/unchartedsoftware/egol/api/conf"
	"github.com/unchartedsoftware/egol/api/middleware"
	"github.com/unchartedsoftware/egol/api/redis"
	"github.com/unchartedsoftware/egol/api/sim"
	"github.com/unchartedsoftware/egol/api/util"
	"github.com/unchartedsoftware/egol/api/ws"
)

const (
	websocketRoute = "/connect"
)

var (
	exit      = make(chan bool)
	organisms []*sim.Organism
	redisConn *redis.Connection
	clients   cmap.ConcurrentMap
	config    *conf.Conf
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

func loop() {
	iteration := 0

	for {
		// get timestamp
		stamp := util.Timestamp()

		// check if should exit
		if shouldExit() {
			break
		}

		// apply constraints to each organism
		sim.ApplyConstraints(organisms)

		// determine AI input for each organism
		updates := sim.Iterate(organisms)

		// write out current state
		stateID := fmt.Sprintf("%s-%d-state", config.SimID, iteration)
		stateBytes, err := json.Marshal(organisms)
		if err != nil {
			log.Error(err)
			continue
		}
		err = redisConn.Set(stateID, stateBytes)
		if err != nil {
			log.Error(err)
			continue
		}

		// write out delta
		updateID := fmt.Sprintf("%s-%d-update", config.SimID, iteration)
		updateBytes, err := json.Marshal(updates)
		if err != nil {
			log.Error(err)
			continue
		}
		err = redisConn.Set(updateID, updateBytes)
		if err != nil {
			log.Error(err)
			continue
		}

		log.Info("Iteration: ", iteration)
		log.Info("organisms: ", organisms)
		log.Info("updates: ", updates)

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
					Data:    updateBytes,
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
					Data:    stateBytes,
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

	organisms = make([]*sim.Organism, 3)

	// debug states for 3 organisms

	var aliveState = sim.State{
		Type: "alive",
	}

	var defaultAttributes = sim.Attributes{}

	organisms[0] = &sim.Organism{
		ID:    0,
		State: &aliveState,
		Attributes: &defaultAttributes,
		Position: mgl32.Vec3{0.0, 1.0, 0.0},
	}

	organisms[1] = &sim.Organism{
		ID:    0,
		State: &aliveState,
		Attributes: &defaultAttributes,
		Position: mgl32.Vec3{0.5, 1.5, 2.0},
	}

	organisms[2] = &sim.Organism{
		ID:    0,
		State: &aliveState,
		Attributes: &defaultAttributes,
		Position: mgl32.Vec3{3.0, 3.0, 3.0},
	}

	// get redis connection
	redisConn = redis.NewConnection(config.RedisHost, config.RedisPort, 0)

	// create clients map
	clients = cmap.New()

	// start server
	go loop()

	server()
}
