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

package nfx509

import (
	"crypto/x509"
	"encoding/pem"
	"github.com/pkg/errors"
	"io"
)

// MarshalToPem takes the list of x509 certs and writes them in pem format to the provided writer
func MarshalToPem(certs []*x509.Certificate, writer io.Writer) error {
	for _, cert := range certs {
		encodeErr := pem.Encode(writer, &pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw})
		if encodeErr != nil {
			return errors.Errorf("unexpected error while writing pem. no further attempt to marshall certificates will be done: %s", encodeErr)
		}
	}
	return nil
}
