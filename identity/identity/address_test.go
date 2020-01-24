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

package identity

import (
	"github.com/netfoundry/ziti-foundation/identity/certtools"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestParseAddrPem(t *testing.T) {
	key, err := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)

	assert.NoError(t, err)

	cert := &x509.Certificate{
		Subject: pkix.Name{
			CommonName: "Test Name",
		},
		SerialNumber: big.NewInt(169),

		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageClientAuth,
		},

		DNSNames: []string{
			"test.netfoundry.io",
		},
	}

	certDer, err := x509.CreateCertificate(rand.Reader, cert, cert, key.Public(), key)

	assert.NoError(t, err)

	certPem := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certDer,
	}

	pemBody := string(pem.EncodeToMemory(certPem))
	pemAddr := "pem:" + pemBody

	p, err := parseAddr(pemAddr)

	assert.NoError(t, err)
	assert.Equal(t, "pem", p.Scheme)
	assert.Equal(t, pemBody, p.Opaque)
}

func TestParseAddrEngine(t *testing.T) {
	if len(certtools.ListEngines()) < 1 {
		t.Skip("No engine tests")
	}

	// setup
	engine := certtools.ListEngines()[0]

	engineBody := "some-driver?slot=0"
	engineAddr := engine + ":" + engineBody

	p, err := parseAddr(engineAddr)

	assert.NoError(t, err)
	assert.Equal(t, engine, p.Scheme)
	assert.Equal(t, "some-driver", p.Opaque)
	assert.Equal(t, "0", p.Query().Get("slot"))
}
