package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("test", 0777)
	os.MkdirAll("subdir/sub1/sub2", 0777)
	err := os.Remove("test")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("subdir")

	userFile := "test.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("これはテストです。\r\n")
		fout.Write([]byte("これはテストです。\r\n"))
	}
}
