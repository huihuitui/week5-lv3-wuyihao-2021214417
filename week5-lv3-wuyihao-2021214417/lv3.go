package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

func UserInfoWrite() {
	for _, user := range Usermap {
		byte, _ := json.Marshal(user)
		byte = append(byte, '\n')
		fmt.Println(string(byte))
		_, err := datafile.Write(byte)
		if err != nil {
			log.Println(err)
		}
	}
}
func ReadInfo() {
	all, err := io.ReadAll(datafile)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(bytes.NewReader(all))
	for {
		readstring, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}
		err = json.Unmarshal([]byte(readstring), &Usermap)
		if err != nil {
			panic(err)
		}
	}
}
