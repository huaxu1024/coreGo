package nosql

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
var Conn redis.Conn
var Err error

func init() {
	Conn, Err = redis.Dial("tcp", "dev.docker.redis01:6380", redis.DialPassword("T0HX9ZzXwJOfwkb"))
	if Err != nil {
		fmt.Println("Connect to redis error", Err)
		return
	}
}