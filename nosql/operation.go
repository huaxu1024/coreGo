package nosql

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func Set(key, value string) {
	_, err := Conn.Do("SET", key, value)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}

func SetEx(key, value, expir string) {
	_, err := Conn.Do("SET", key, value, "EX", expir)
	if err != nil {
		fmt.Println("redis set ex failed:", err)
	}
}

func Get(key string) string {
	username, err := redis.String(Conn.Do("GET", key))
	if err != nil {
		fmt.Println("redis get failed:", err)
		return ""
	}
	fmt.Printf("Get mykey: %v \n", username)
	return username
}

func Del(key string) {
	_, err := Conn.Do("DEL", key)
	if err != nil {
		fmt.Println("redis delete failed:", err)
	}
}

func Push(key string, nodes []string) {
	_, err := Conn.Do("lpush", key, "a","b","c")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}

func Lrange(key, left, right string) []interface{}{
	values, err := redis.Values(Conn.Do("lrange", key, left, right))
	if err != nil {
		fmt.Println("redis lrange failed:", err)
	}
	return values
}

