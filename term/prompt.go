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

package term

import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"os"
	"strings"
	"syscall"
)

func PromptPassword(prompt string, allowEmpty bool) (string, error) {
	prompting := true
	password := ""
	for prompting {
		fmt.Print(prompt)

		//Be aware that debugging this on windows in GoLand will cause this to error w/ an invalid
		//stdin handler. To debug compile w/o optimizations and attach to the process
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		println("")
		if err != nil {
			return "", err
		}

		password = strings.TrimSpace(string(bytePassword))

		if !allowEmpty && password == "" {
			fmt.Println("\nError: a value must be set")
		} else {
			prompting = false
		}
	}

	return password, nil
}

// Prompt will output the given prompt and return the line entered by the user.
//
//	The line will be trimmed of white-space.
func Prompt(prompt string) (string, error) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	text = strings.TrimSpace(text)

	return text, nil
}
