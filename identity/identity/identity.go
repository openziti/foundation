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
	"github.com/netfoundry/ziti-foundation/identity/certtools"
	"github.com/netfoundry/ziti-foundation/util/tlz"
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
)

type Identity interface {
	Cert() *tls.Certificate
	ServerCert() *tls.Certificate
	CA() *x509.CertPool
	ServerTLSConfig() *tls.Config
	ClientTLSConfig() *tls.Config
}

type IdentityConfig struct {
	Key        string `json:"key" yaml:"key" mapstructure:"key"`
	Cert       string `json:"cert" yaml:"cert" mapstructure:"cert"`
	ServerCert string `json:"server_cert,omitempty" yaml:"server_cert,omitempty" mapstructure:"server_cert,omitempty"`
	ServerKey  string `json:"server_key,omitempty" yaml:"server_key,omitempty" mapstructure:"server_key,omitempty"`
	CA         string `json:"ca,omitempty" yaml:"ca,omitempty" mapstructure:"ca"`
}

type ID struct {
	IdentityConfig

	cert       *tls.Certificate
	serverCert *tls.Certificate
	ca         *x509.CertPool
}

func (id *ID) Cert() *tls.Certificate {
	return id.cert
}

func (id *ID) ServerCert() *tls.Certificate {
	return id.serverCert
}

func (id *ID) CA() *x509.CertPool {
	return id.ca
}

func (i *ID) ServerTLSConfig() *tls.Config {
	if i.serverCert == nil {
		return nil
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{*i.ServerCert()},
		RootCAs:      i.CA(),
		ClientAuth:   tls.RequireAnyClientCert,
		MinVersion:   tlz.GetMinTlsVersion(),
		CipherSuites: tlz.GetCipherSuites(),
	}
	tlsConfig.BuildNameToCertificate()
	return tlsConfig
}

func (i *ID) ClientTLSConfig() *tls.Config {
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{*i.Cert()},
		RootCAs:      i.CA(),
	}
	tlsConfig.BuildNameToCertificate()

	return tlsConfig
}

func LoadIdentity(cfg IdentityConfig) (Identity, error) {
	id := &ID{
		cert: &tls.Certificate{},
	}

	var err error
	id.cert.PrivateKey, err = LoadKey(cfg.Key)
	if err != nil {
		return nil, err
	}

	if idCert, err := loadCert(cfg.Cert); err != nil {
		return id, err
	} else {
		id.cert.Certificate = make([][]byte, len(idCert))
		for i, c := range idCert {
			id.cert.Certificate[i] = c.Raw
		}
		id.cert.Leaf = idCert[0]
	}

	// Server Cert is optional
	if cfg.ServerCert != "" {
		if svrCert, err := loadCert(cfg.ServerCert); err != nil {
			return id, err
		} else {
			serverKey := id.cert.PrivateKey
			if cfg.ServerKey != "" {
				serverKey, err = LoadKey(cfg.ServerKey)
				if err != nil {
					return nil, err
				}
			}
			id.serverCert = &tls.Certificate{
				PrivateKey:  serverKey,
				Certificate: make([][]byte, len(svrCert)),
				Leaf:        svrCert[0],
			}
			for i, c := range svrCert {
				id.serverCert.Certificate[i] = c.Raw
			}
		}
	}

	// CA bundle is optional
	if cfg.CA != "" {
		if id.ca, err = loadCABundle(cfg.CA); err != nil {
			return id, err
		}
	}

	return id, nil
}

func LoadKey(keyAddr string) (crypto.PrivateKey, error) {
	if keyUrl, err := parseAddr(keyAddr); err != nil {
		return nil, err
	} else {

		switch keyUrl.Scheme {
		case "pem":
			return certtools.LoadPrivateKey([]byte(keyUrl.Opaque))
		case "file", "":
			return certtools.GetKey(nil, keyUrl.Path, "")
		default:
			// engine key format: "{engine_id}:{engine_opts} see specific engine for supported options
			return certtools.GetKey(keyUrl, "", "")
			//return nil, fmt.Errorf("could not load key, location scheme not supported (%s) or address not defined (%s)", keyUrl.Scheme, keyAddr)
		}
	}
}

func loadCert(certAddr string) ([]*x509.Certificate, error) {
	if certUrl, err := parseAddr(certAddr); err != nil {
		return nil, err
	} else {
		switch certUrl.Scheme {
		case "pem":
			return certtools.LoadCert([]byte(certUrl.Opaque))
		case "file", "":
			return certtools.LoadCertFromFile(certUrl.Path)
		default:
			return nil, fmt.Errorf("could not load cert, location scheme not supported (%s) or address not defined (%s)", certUrl.Scheme, certAddr)
		}
	}
}

func loadCABundle(caAddr string) (*x509.CertPool, error) {
	if caUrl, err := parseAddr(caAddr); err != nil {
		return nil, err
	} else {
		pool := x509.NewCertPool()
		var bundle []byte
		switch caUrl.Scheme {
		case "pem":
			bundle = []byte(caUrl.Opaque)

		case "file", "":
			if bundle, err = ioutil.ReadFile(caUrl.Path); err != nil {
				return nil, err
			}

		default:
			return nil, fmt.Errorf("NO valid Cert location specified")
		}
		pool.AppendCertsFromPEM(bundle)
		return pool, nil
	}

}
