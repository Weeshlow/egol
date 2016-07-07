package conf

var config *Conf

// Conf represents all the ingest runtime flags passed to the binary.
type Conf struct {
	SimID        string
	FrameMS      int64
	NumOrganisms uint64
	Port         string
	RedisHost    string
	RedisPort    string
	PublicDir    string
}

// SaveConf saves the parsed conf.
func SaveConf(c *Conf) {
	config = c
}

// GetConf returns a copy of the parsed conf.
func GetConf() Conf {
	return *config
}
