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

package nfpem

import (
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"fmt"
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

func EncodeToString(cert *x509.Certificate) string {
	result := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	})

	return string(result)
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

//Assumes the first PEM block is the target
func FingerprintFromPem(pem string) string {
	certs := PemToX509(pem)
	if len(certs) == 0 {
		return ""
	}
	return FingerprintFromX509(certs[0])
}

func FingerprintFromX509(cert *x509.Certificate) string {
	if cert == nil {
		return ""
	}
	return fmt.Sprintf("%x", sha1.Sum(cert.Raw))
}
