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

package identity

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/big"
	"os"
	"testing"
)

func mkCert(cn string, dns []string) (crypto.Signer, *x509.Certificate) {
	key, _ := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)


	cert := &x509.Certificate{
		Subject: pkix.Name{
			CommonName: cn,
		},
		SerialNumber: big.NewInt(169),

		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageClientAuth,
		},

		DNSNames: dns,
	}
	return key, cert
}

func TestLoadIdentityWithPEM(t *testing.T) {
	// setup
	key, cert := mkCert("Test Name", []string{"test.netfoundry.io"})

	keyDer, _ := x509.MarshalECPrivateKey(key.(*ecdsa.PrivateKey))
	keyPem := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: keyDer,
	}

	certDer, err := x509.CreateCertificate(rand.Reader, cert, cert, key.Public(), key)
	certPem := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certDer,
	}

	cfg := Config{
		Key:  "pem:" + string(pem.EncodeToMemory(keyPem)),
		Cert: "pem:" + string(pem.EncodeToMemory(certPem)),
	}

	id, err := LoadIdentity(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, id.Cert())
	assert.NotNil(t, id.Cert().Leaf)
	assert.Equal(t, key.Public(), id.Cert().Leaf.PublicKey)

}

func TestLoadIdentityWithPEMChain(t *testing.T) {
	// setup
	parentKey, parentCert := mkCert("Parent", []string{})
	parentDer, _ := x509.CreateCertificate(rand.Reader, parentCert, parentCert, parentKey.Public(), parentKey)

	key, cert := mkCert("Test Child", []string{"client.netfoundry.io"})
	certDer, _ := x509.CreateCertificate(rand.Reader, cert, parentCert, key.Public(), parentKey)

	keyDer, _ := x509.MarshalECPrivateKey(key.(*ecdsa.PrivateKey))
	keyPem := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: keyDer,
	}

	parentPem := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: parentDer,
	}

	certPem := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certDer,
	}

	cfg := Config{
		Key:  "pem:" + string(pem.EncodeToMemory(keyPem)),
		Cert: "pem:" + string(pem.EncodeToMemory(certPem)) + string(pem.EncodeToMemory(parentPem)),
	}

	id, err := LoadIdentity(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, id.Cert())
	assert.Equal(t, 2, len(id.Cert().Certificate))
	assert.NotNil(t, id.Cert().Leaf)
	assert.Equal(t, id.Cert().Leaf.Subject.CommonName, "Test Child")
	assert.Equal(t, key.Public(), id.Cert().Leaf.PublicKey)

}

func TestLoadIdentityWithFile(t *testing.T) {
	// setup
	key, _ := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)

	keyDer, _ := x509.MarshalECPrivateKey(key)
	keyPem := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: keyDer,
	}

	keyFile, _ := ioutil.TempFile(os.TempDir(), "test-key")

	defer os.Remove(keyFile.Name())

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
	certPem := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certDer,
	}

	certFile, _ := ioutil.TempFile(os.TempDir(), "test-cert")
	defer os.Remove(certFile.Name())

	pem.Encode(keyFile, keyPem)
	pem.Encode(certFile, certPem)

	cfg := Config{
		Key:  "file://" + keyFile.Name(),
		Cert: "file://" + certFile.Name(),
	}

	id, err := LoadIdentity(cfg)
	assert.NoError(t, err)
	assert.Equal(t, key, id.Cert().PrivateKey)

}