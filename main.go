//Package inmemorydb creates a tcp session with a map to accept and store key-value pairs required for immediate execution
//syntax to use is to
//SET "key" "value"
//GET "key"
//DEL "key"
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		switch fs[0] {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Println(v)
		case "SET":
			if len(fs) != 3 {
				fmt.Println("Invalid syntax")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
		case "DEL":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Println("Invalid entry, Read instructions")

		}

	}
}
