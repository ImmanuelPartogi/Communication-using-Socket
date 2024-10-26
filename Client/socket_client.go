package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		_, err := fmt.Fprintf(conn, "%s\n", message)
		if err != nil {
			log.Println(err)
			return
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
