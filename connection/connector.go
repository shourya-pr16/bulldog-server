package connection

import (
	"fmt"
	"net"
	"strings"

	"github.com/shourya-pr16/bulldog-server/property"
)

func InitSocket() {
	conf := property.ReadProps()
	listener, err := net.Listen("tcp4", ":"+conf.Port)
	if err != nil {
		fmt.Errorf("Error in listening : %+v", err)
		panic("Unable to open socket on given port. Maybe already occupied.")
	}

	fmt.Printf("------- Server started listening on port: %s -------\n\n", conf.Port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error while accepting connection: %s\n", err)
			return
		}

		go handleConnection(conn)
	}
}

func handleConnection(connection net.Conn) {
	fmt.Println("Connection Established: ")
	for {
		buff := make([]byte, 1024)
		_, err := connection.Read(buff)
		if err != nil {
			fmt.Println(err)
			return
		}

		result := string(buff)

		if strings.Trim(result, "\n\r\x00") == "close" {
			connection.Close()
			fmt.Println("Connection Closed")
			return
		}
		fmt.Printf("%s", result)
	}
}
