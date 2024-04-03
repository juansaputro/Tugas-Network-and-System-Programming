package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

func server() {
	listener, err := net.Listen("tcp", "127.0.0.1:191203")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go HandleServerConnection(clientConn)
	}
}

func HandleServerConnection(client net.Conn) {
	var size uint32
	err := binary.Read(client, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}
	bytMsg := make([]byte, size)
	client.SetReadDeadline(time.Now().Add(10 * time.Second))
	_, err = client.Read(bytMsg)
	if err != nil {
		panic(err)
	}
	strMsg := string(bytMsg)
	fmt.Printf("Received: %s\n", strMsg)

	var reply string
	if strings.HasSuffix(strMsg, ".zip") {
		reply = "File Has Been Received"
	} else if strings.Contains(strMsg, ".") {
		reply = "Only Zip File Can Be Uploaded"
	} else {
		reply = "Message Has Been Received"
	}

	err = binary.Write(client, binary.LittleEndian, uint32(len(reply)))
	if err != nil {
		panic(err)
	}
	client.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = client.Write([]byte(reply))
	if err != nil {
		panic(err)
	}
}
