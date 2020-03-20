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

package debugz

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func GenerateStack() string {
	stackBuf := make([]byte, 1024*1024)
	size := runtime.Stack(stackBuf, true)
	return string(stackBuf[:size])
}

func GenerateLocalStack() string {
	stackBuf := make([]byte, 1024*10)
	size := runtime.Stack(stackBuf, false)
	return string(stackBuf[:size])
}

func DumpStack() {
	fmt.Println(GenerateStack())
}

func AddStackDumpHandler() {
	go func() {
		signalC := make(chan os.Signal, 1)
		signal.Notify(signalC, syscall.SIGQUIT)
		for range signalC {
			fmt.Printf("\n DUMPING STACK AS REQUESTED BY SIGQUIT \n\n%v\n", GenerateStack())
		}
	}()
}