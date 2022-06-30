/*
	Copyright NetFoundry Inc.

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
	"errors"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("invalid number of arguments count (%v). Usage: half-close-test <client|server> <address>", len(os.Args))
	}

	if os.Args[1] == "client" {
		client(os.Args[2])
	} else if os.Args[1] == "server" {
		server(os.Args[2])
	} else {
		fmt.Println("invalid arguments. Usage: half-close-test <client|server> <address>")
	}
}

func server(address string) {
	l, err := net.Listen("tcp", address)
	panicOnError(err)

	for {
		conn, err := l.Accept()
		panicOnError(err)
		go handleServerConn(conn)
	}
}

func handleServerConn(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil && errors.Is(err, io.EOF) {
			if n > 0 {
				fmt.Printf("client said: '%v'\n", string(buf[:n]))
			}
			break
		}
		panicOnError(err)
		fmt.Printf("client said: '%v'\n", string(buf[:n]))
	}

	_, err := conn.Write([]byte("goodbye"))
	panicOnError(err)
	fmt.Println("server sent: goodbye")

	err = conn.(closeWriter).CloseWrite()
	panicOnError(err)
}

func client(address string) {
	conn, err := net.Dial("tcp", address)
	panicOnError(err)

	_, err = conn.Write([]byte("hello"))
	panicOnError(err)
	fmt.Println("client sent: 'hello'")

	err = conn.(closeWriter).CloseWrite()
	panicOnError(err)

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil && errors.Is(err, io.EOF) {
			if n > 0 {
				fmt.Printf("server responded: '%v'\n", string(buf[:n]))
			}
			break
		}
		panicOnError(err)
		fmt.Printf("server responded: '%v'\n", string(buf[:n]))
	}
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

type closeWriter interface {
	CloseWrite() error
}
