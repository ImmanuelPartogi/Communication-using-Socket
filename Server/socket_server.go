package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Membuat koneksi database
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/communication")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Membuat listener socket
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	fmt.Println("Server started. Listening on port 8000...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn, db)
	}
}

func handleConnection(conn net.Conn, db *sql.DB) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Printf("Received message: %s\n", message)

		// Simpan pesan ke database
		_, err := db.Exec("INSERT INTO chat (message) VALUES (?)", message)
		if err != nil {
			log.Println(err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
