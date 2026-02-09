package seeder

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
)

func SendFile(filePath, listenAddr string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	defer ln.Close()

	fmt.Println("Seeder listening on", listenAddr)
	conn, err := ln.Accept()
	if err != nil {
		return err
	}
	defer conn.Close()

	filename := info.Name()
	binary.Write(conn, binary.BigEndian, uint32(len(filename)))
	conn.Write([]byte(filename))

	binary.Write(conn, binary.BigEndian, uint64(info.Size()))

	written, err := io.Copy(conn, file)
	if err != nil {
		return err
	}

	fmt.Printf("Sent %d bytes\n", written)
	return nil
}
