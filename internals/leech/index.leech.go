package leech

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
)

func ReceiveFile(serverAddr string) error {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Read filename length
	var nameLen uint32
	binary.Read(conn, binary.BigEndian, &nameLen)

	// Read filename
	nameBuf := make([]byte, nameLen)
	io.ReadFull(conn, nameBuf)
	filename := string(nameBuf)

	// Read file size
	var fileSize uint64
	binary.Read(conn, binary.BigEndian, &fileSize)

	out, err := os.Create("received_" + filename)
	if err != nil {
		return err
	}
	defer out.Close()

	written, err := io.CopyN(out, conn, int64(fileSize))
	if err != nil {
		return err
	}

	fmt.Printf("Received %d bytes\n", written)
	return nil
}
