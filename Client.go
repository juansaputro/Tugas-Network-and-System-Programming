package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func menu() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1. Send Massage")
		fmt.Println("2. Exit")
		scanner.Scan()
		ch := scanner.Text()
		if ch == "1" {
			sendMessageMenu()
		} else if ch == "2" {
			fmt.Println("Thankyou For Using This Program!")
			break
		}
	}
}

func sendMessageMenu() {
	scanner := bufio.NewScanner(os.Stdin)
	var message string
	for {
		fmt.Print("Please Insert Your Message: ")
		scanner.Scan()
		message = scanner.Text()
		if len(message) < 10 {
			fmt.Println("Message Cannot be less than 10 char")
		} else if strings.Contains(message, "kata kasar") {
			fmt.Println("No Bad Word Allowed In It")
		} else if strings.Compare(message, "Hello World test") == 0 {
			fmt.Println("Message Is Hello World Test")
		} else {
			break
		}

	}
	sendMessageToServer(message)
}

func sendMessageToServer(message string) {
	serverConn, err := net.DialTimeout("tcp", "127.0.0.1:191203", 3*time.Second)
	if err != nil {
		panic(err)
	}
	defer serverConn.Close()

	err = binary.Write(serverConn, binary.LittleEndian, uint32(len(message)))
	if err != nil {
		panic(err)
	}

	_, err = serverConn.Write([]byte(message))
	if err != nil {
		panic(err)
	}

	var size uint32
	err = binary.Read(serverConn, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}
	bytReply := make([]byte, size)
	serverConn.SetReadDeadline(time.Now().Add(10 * time.Second))
	_, err = serverConn.Read(bytReply)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Replied: %s\n", string(bytReply))

}

func main() {
	menu()
}
