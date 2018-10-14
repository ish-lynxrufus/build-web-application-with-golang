package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("test", 0777)
	os.MkdirAll("subdir/sub1/sub2", 0777)
	err := os.Remove("dummy")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("subdir")
}
