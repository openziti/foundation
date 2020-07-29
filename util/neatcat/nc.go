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
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/foundation/util/info"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io"
	"net"
	"os"
	"strconv"
)

func init() {
	pfxlog.Global(logrus.InfoLevel)
	pfxlog.SetPrefix("github.com/openziti/")
}

func init() {
	root.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging")
	root.PersistentFlags().BoolVarP(&useUDP, "useUDP", "u", false, "Use UDP instead of TCP")
	root.PersistentFlags().StringVar(&logFormatter, "log-formatter", "", "Specify log formatter [json|pfxlog|text]")
}

var root = &cobra.Command{
	Use:   "neatcat <host> <port>",
	Short: "NeatCat",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose {
			logrus.SetLevel(logrus.DebugLevel)
		}

		switch logFormatter {
		case "pfxlog":
			logrus.SetFormatter(pfxlog.NewFormatterStartingToday())
		case "json":
			logrus.SetFormatter(&logrus.JSONFormatter{})
		case "text":
			logrus.SetFormatter(&logrus.TextFormatter{})
		default:
			// let logrus do its own thing
		}
	},
	Args: cobra.MinimumNArgs(2),
	Run:  runFunc,
}

var verbose bool
var useUDP bool
var logFormatter string

func main() {
	if err := root.Execute(); err != nil {
		fmt.Printf("error: %s", err)
	}
}

func runFunc(cmd *cobra.Command, args []string) {

	log := pfxlog.Logger()

	host := args[0]
	port, err := strconv.Atoi(args[1])

	if err != nil {
		log.WithError(err).Fatal("unable to parse port")
	}

	network := "tcp"
	if useUDP {
		network = "udp"
	}

	conn, err := net.Dial(network, fmt.Sprintf("%v:%v", host, port))
	if err != nil {
		log.WithError(err).Fatalf("unable to dial %v:%v:%v", network, host, port)
	}

	pfxlog.Logger().Debug("connected")
	go Copy(conn, os.Stdin)
	Copy(os.Stdout, conn)
}

func Copy(writer io.Writer, reader io.Reader) {
	buf := make([]byte, info.MaxPacketSize)
	bytesCopied, err := io.CopyBuffer(writer, reader, buf)
	pfxlog.Logger().Infof("Copied %v bytes", bytesCopied)
	if err != nil {
		pfxlog.Logger().Errorf("error while copying bytes (%v)", err)
	}
}

// CopyAndLog does what io.Copy does but with additional logging
func CopyAndLog(context string, writer io.Writer, reader io.Reader) {
	buf := make([]byte, info.MaxPacketSize)

	var bytesRead, totalBytesRead, bytesWritten, totalBytesWritten int
	var readErr, writeErr error

	for {
		bytesRead, readErr = reader.Read(buf)
		totalBytesRead += bytesRead
		if bytesRead > 0 {
			bytesWritten, writeErr = writer.Write(buf[:bytesRead])
			totalBytesWritten += bytesWritten
			if writeErr != nil {
				pfxlog.Logger().WithError(writeErr).Error("Write failure on copy")
			}
		}

		if readErr != nil && readErr != io.EOF {
			pfxlog.Logger().WithError(readErr).Error("Read failure on copy")
		}

		if readErr != nil || writeErr != nil {
			return
		}

		_, _ = fmt.Fprintf(os.Stderr, "%v: Read %v (%v total), Wrote %v (%v total)\n",
			context, bytesRead, totalBytesRead, bytesWritten, totalBytesWritten)
	}
}
