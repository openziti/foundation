/*
	Copyright NetFoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package main

import (
	"fmt"
	"github.com/openziti/foundation/util/info"
	"io"
	"net"
	"os"
	"strconv"
)

func main() {
	port := 9087
	serverType := "tcp"
	if len(os.Args) > 1 {
		var err error
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Failed to parse port argument: %v\n", os.Args[1])
			panic(err)
		}
	}
	if len(os.Args) > 2 {
		serverType = os.Args[2]
	}

	fmt.Printf("Starting echo server on port %v of type %v", port, serverType)

	if serverType == "tcp" {
		tcpEcho(port)
	} else if serverType == "udp" {
		udpEcho(port)
	} else {
		panic(fmt.Errorf("unknown network type: %v", serverType))
	}
}

func udpEcho(port int) {
	packetConn, err := net.ListenPacket("udp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		panic(err)
	}

	buf := make([]byte, info.MaxPacketSize)
	fmt.Printf("Running on port: %v, Buffer size: %v\n", port, len(buf))

	var bytesReadTotal, bytesWrittenTotal int

	for {
		bytesRead, sourceAddr, err := packetConn.ReadFrom(buf)
		bytesReadTotal += bytesRead
		if err != nil {
			panic(err)
		}

		bytesWritten, err := packetConn.WriteTo(buf[0:bytesRead], sourceAddr)
		if err != nil {
			panic(err)
		}
		bytesWrittenTotal += bytesWritten
		fmt.Printf("%v bytes read (%v total) from %v, %v bytes (%v total) echoed\n",
			bytesRead, bytesReadTotal, sourceAddr, bytesWritten, bytesWrittenTotal)
	}
}

func tcpEcho(port int) {
	listenSocket, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listenSocket.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Printf("new connection: %v\n", conn.LocalAddr())
		go echoConn(conn)
	}
}

func echoConn(conn net.Conn) {
	defer fmt.Printf("exited conn %v", conn.LocalAddr())
	buf := make([]byte, info.MaxPacketSize)

	var bytesReadTotal, bytesWrittenTotal int

	for {
		bytesRead, err := conn.Read(buf)
		bytesReadTotal += bytesRead
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}

		bytesWritten, err := conn.Write(buf[0:bytesRead])
		if err != nil {
			panic(err)
		}
		bytesWrittenTotal += bytesWritten
		fmt.Printf("%v bytes read (%v total), %v bytes (%v total) echoed\n",
			bytesRead, bytesReadTotal, bytesWritten, bytesWrittenTotal)
	}
}
