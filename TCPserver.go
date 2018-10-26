package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func jsoncompil() []byte {

	rand.Seed(time.Now().UTC().UnixNano())
	var bytes int
	bytes = rand.Intn(100)

	mapVar1 := map[string]interface{}{"Name": "Ivan", "Age": bytes}
	mapVar2 := map[string]interface{}{"Name": "Andrey", "Age": 32}
	mapVar3 := map[string]interface{}{"Name": "Mikhail", "Age": 23}
	mapVar4 := map[string]map[string]interface{}{"Person": mapVar3, "Person_friends": {"Brother": mapVar1, "Employer": mapVar2}}

	data, _ := json.Marshal(mapVar4)
	return data
}

func main() {

	listener, err := net.Listen("tcp", ":4545")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("server is listening...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		input := make([]byte, (1024 * 4))
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error", err)
			break
		}

		source := string(input[0:n])
		target := jsoncompil()

		fmt.Println(source, "- command recieved")
		conn.Write([]byte(target))
	}
}
