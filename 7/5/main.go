package main

import (
	"fmt"
	"os"
)

func main() {
	// ディレクトリの操作
	os.Mkdir("test", 0777)
	os.MkdirAll("subdir/sub1/sub2", 0777)
	err := os.Remove("test")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("subdir")

	// ファイルへの書き込み
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

	// ファイルの読み込み
	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
