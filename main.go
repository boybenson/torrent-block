package main

import (
	"example/torrent/internals/leech"
	"example/torrent/internals/seeder"
	"flag"
	"fmt"
)

func main() {
	mode := flag.String("mode", "leech", "Mode: 'seeder' or 'leech'")
	file := flag.String("file", "test.txt", "File to send (seeder mode)")
	addr := flag.String("addr", "127.0.0.1:9000", "Address to listen or connect to")

	flag.Parse()

	switch *mode {
	case "seeder":
		fmt.Println("Running in seeder mode...")
		if err := seeder.SendFile(*file, *addr); err != nil {
			fmt.Println("Error:", err)
		}
	case "leech":
		fmt.Println("Running in leech mode...")
		if err := leech.ReceiveFile(*addr); err != nil {
			fmt.Println("Error:", err)
		}
	default:
		fmt.Println("Invalid mode. Use 'seeder' or 'leech'.")
	}
}
