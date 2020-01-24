/*
	Copyright 2019 NetFoundry, Inc.

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

package nfpem

import (
	"crypto/x509"
	"encoding/pem"
)

// An initial limit of how deep a valid cert chain might be.
const certDepth = 4

func DecodeAll(pemBytes []byte) []*pem.Block {
	var blocks []*pem.Block
	if len(pemBytes) < 1 {
		return blocks
	}
	b, rest := pem.Decode(pemBytes)

	for b != nil {
		blocks = append(blocks, b)
		b, rest = pem.Decode(rest)
	}
	return blocks
}

func PemToX509(pem string) []*x509.Certificate {

	pemBytes := []byte(pem)
	certs := make([]*x509.Certificate, 0)
	for _, block := range DecodeAll(pemBytes) {
		xcerts, err := x509.ParseCertificate(block.Bytes)
		if err == nil && xcerts != nil {
			certs = append(certs, xcerts)
		}
	}
	return certs
}

