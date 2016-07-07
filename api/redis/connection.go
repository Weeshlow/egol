package redis

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"

	"github.com/garyburd/redigo/redis"
)

// Compress compresses the provided bytes.
func Compress(data []byte) ([]byte, error) {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	_, err := w.Write(data)
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	return b.Bytes()[0:], nil
}

// Decompress decompresses the provided bytes.
func Decompress(data []byte) ([]byte, error) {
	b := bytes.NewBuffer(data[0:])
	r, err := gzip.NewReader(b)
	if err != nil {
		return nil, err
	}
	data, err = ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	err = r.Close()
	if err != nil {
		return nil, err
	}
	return data[0:], nil
}

// Connection represents a single connection to a redis server.
type Connection struct {
	conn   redis.Conn
	expiry int
}

// NewConnection instantiates and returns a new redis connection.
func NewConnection(host, port string, expirySeconds int64) *Connection {
	return &Connection{
		conn:   getConnection(host, port),
		expiry: int(expirySeconds),
	}
}

// Get when given a string key will return a byte slice of data from redis.
func (r *Connection) Get(key string) ([]byte, error) {
	bytes, err := redis.Bytes(r.conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return Decompress(bytes)
}

// Set will store a byte slice under a given key in redis.
func (r *Connection) Set(key string, value []byte) error {
	compressed, err := Compress(value)
	if err != nil {
		return err
	}
	if r.expiry > 0 {
		_, err = r.conn.Do("SET", key, compressed, "NX", "EX", r.expiry)
	} else {
		_, err = r.conn.Do("SET", key, compressed)
	}
	return err
}

// Exists returns whether or not a key exists in redis.
func (r *Connection) Exists(key string) (bool, error) {
	return redis.Bool(r.conn.Do("Exists", key))
}

// Close closes the redis connection.
func (r *Connection) Close() {
	r.conn.Close()
}
