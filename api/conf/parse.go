package conf

import (
	"flag"

	"github.com/unchartedsoftware/egol/api/util"
)

const (
	defaultFrameMs       = 100
	defaultOrganismCount = 1000
	defaultPort          = "8080"
	defaultRedisHost     = "localhost"
	defaultRedisPort     = "6379"
	defaultPublicDir     = "./build/public"
)

// ParseCommandLine parses the commandline arguments and returns a Conf object.
func ParseCommandLine() *Conf {
	simID := flag.String("sim-id", util.RandID(), "Simulation ID")
	frameMS := flag.Int64("frame-ms", defaultFrameMs, "Number of milliseconds per frame")
	numOrganisms := flag.Int64("num-organisms", defaultOrganismCount, "Number of organisms to simulate")
	port := flag.String("port", defaultPort, "Port to run the server on")
	redisHost := flag.String("redis-host", defaultRedisHost, "Redis host")
	redisPort := flag.String("redis-port", defaultRedisPort, "Redis port")
	publicDir := flag.String("public-dir", defaultPublicDir, "Public directory to serve from")
	// parse the flags
	flag.Parse()
	// Set and save config
	config := &Conf{
		SimID:        *simID,
		FrameMS:      *frameMS,
		NumOrganisms: uint64(*numOrganisms),
		Port:         *port,
		RedisHost:    *redisHost,
		RedisPort:    *redisPort,
		PublicDir:    *publicDir,
	}
	SaveConf(config)
	return config
}
