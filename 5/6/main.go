package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	coon, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	defer coon.Close()

	// write
	r, err := coon.Do("SET", "test", "14")
	if err != nil {
		panic(err)
	}
	fmt.Println(r)

	// write
	s, err := redis.String(coon.Do("GET", "test"))
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
