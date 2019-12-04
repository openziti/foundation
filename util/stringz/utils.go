/*
	Copyright 2019 Netfoundry, Inc.

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

package stringz

func Contains(a []string, val string) bool {
	_, found := FirstIndexOf(a, val)
	return found
}

func ContainsAll(a []string, values ...string) bool {
	for _, val := range values {
		if !Contains(a, val) {
			return false
		}
	}
	return true
}

func FirstIndexOf(a []string, val string) (int, bool) {
	for idx, s := range a {
		if s == val {
			return idx, true
		}
	}
	return -1, false
}

func Difference(a, b []string) []string {
	mb := map[string]bool{}
	for _, x := range b {
		mb[x] = true
	}
	var ab []string
	for _, x := range a {
		if _, ok := mb[x]; !ok {
			ab = append(ab, x)
		}
	}
	return ab
}

func ToStringSlice(a [][]byte) []string {
	result := make([]string, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = string(a[i])
	}
	return result
}

func OrEmpty(val *string) string {
	if val == nil {
		return ""
	}
	return *val
}