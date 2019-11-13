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

package pkcs11

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rsa"
	"encoding/asn1"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"github.com/miekg/pkcs11"
	"io"
	"math/big"
	"net/url"
	"strconv"
)

const EngineId = "pkcs11"

 //
 // engine supporting generic PKCS#11 HSM driver
 // possible key URLs:
 // - `pkcs11:/usr/lib/softhsm/libsofthsm2.so?slot=0&id=2171` - full driver path
 // - `pkcs11:softhsm2?slot=0&id=2171` - driver id, driver will be loaded according to following rules:
 //                       driver id converted to OS specific library file name (on *nix `lib${driver}.so`)
 //                       then loaded according to dynamic loader configuration (on *nix according to http://man7.org/linux/man-pages/man3/dlopen.3.html)
var Engine = &engine{}

type engine struct {
}

var contexts = map[string]*pkcs11.Ctx{}

var log = pfxlog.ContextLogger(EngineId)

type p11Signer struct {
	c *pkcs11.Ctx
	s pkcs11.SessionHandle
	h pkcs11.ObjectHandle
	m *pkcs11.Mechanism
	label string

	pub crypto.PublicKey
}

func (k *p11Signer) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error) {

	switch k.pub.(type) {
	case *ecdsa.PublicKey:
		return k.signECDSA(digest)
	case *rsa.PublicKey:
		return k.signRSA(digest, opts)

	default:
		return nil, fmt.Errorf("unsupported key")
	}
}

// copied from golang/crypto/rsa/rsa.go
var hashPrefixes = map[crypto.Hash][]byte{
	crypto.MD5:       {0x30, 0x20, 0x30, 0x0c, 0x06, 0x08, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x02, 0x05, 0x05, 0x00, 0x04, 0x10},
	crypto.SHA1:      {0x30, 0x21, 0x30, 0x09, 0x06, 0x05, 0x2b, 0x0e, 0x03, 0x02, 0x1a, 0x05, 0x00, 0x04, 0x14},
	crypto.SHA224:    {0x30, 0x2d, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x04, 0x05, 0x00, 0x04, 0x1c},
	crypto.SHA256:    {0x30, 0x31, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x01, 0x05, 0x00, 0x04, 0x20},
	crypto.SHA384:    {0x30, 0x41, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x02, 0x05, 0x00, 0x04, 0x30},
	crypto.SHA512:    {0x30, 0x51, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x03, 0x05, 0x00, 0x04, 0x40},
	crypto.MD5SHA1:   {}, // A special TLS case which doesn't use an ASN1 prefix.
	crypto.RIPEMD160: {0x30, 0x20, 0x30, 0x08, 0x06, 0x06, 0x28, 0xcf, 0x06, 0x03, 0x00, 0x31, 0x04, 0x14},
}

func (k *p11Signer) signRSA(digest []byte, opts crypto.SignerOpts) ([]byte, error) {
	switch opts.(type) {
	case *rsa.PSSOptions:
		return nil, fmt.Errorf("TODO")
	default: /* PKCS1-v1_5 */
		hash := opts.HashFunc()
		oid := hashPrefixes[hash]
		input := append(oid, digest...)
		mech := []*pkcs11.Mechanism{pkcs11.NewMechanism(pkcs11.CKM_RSA_PKCS, nil)}
		if err := k.c.SignInit(k.s, mech, k.h); err == nil {
			return k.c.Sign(k.s, input)
		} else {
			return nil, err
		}
	}
}

func (k *p11Signer) signECDSA(digest []byte) ([]byte, error) {

	mech := []*pkcs11.Mechanism{k.m}
	if err := k.c.SignInit(k.s, mech, k.h); err != nil {
		return nil, err
	}

	sigBytes, err := k.c.Sign(k.s, digest)
	if err != nil {
		return nil, err
	}

	var sig struct {
		R, S *big.Int
	}

	n := len(sigBytes) / 2
	sig.R = new(big.Int)
	sig.R.SetBytes(sigBytes[:n])
	sig.S = new(big.Int)
	sig.S.SetBytes(sigBytes[n:])

	bytes, err := asn1.Marshal(sig)
	return bytes, err
}

func (k *p11Signer) Public() crypto.PublicKey {
	return k.pub
}

func (*engine) LoadKey(key *url.URL) (crypto.PrivateKey, error) {
	log.WithField("url", key).Debug("loading key")

	driver := key.Path
	if driver == "" {
		if key.Opaque == "" {
			return nil, fmt.Errorf("driver not specified for PKCS#11 engine, see docs")
		}

		driver = "lib" + key.Opaque + ".so"
	}

	ctx, err := getContext(driver)
	if err != nil {
		return nil, err
	}

	opts := key.Query()

	slot := opts.Get("slot")
	var slotId uint
	if slot == "" { // use the first slot returned by token
		if slots, err := ctx.GetSlotList(true); err != nil {
			return nil, err
		} else {
			slotId = slots[0]
			log.Warnf("slot not specified, using first slot reported by the driver (%d)", slotId)
		}
	} else {
		id, _ := strconv.Atoi(slot)
		slotId = uint(id)
	}

	session, err := ctx.OpenSession(slotId, pkcs11.CKF_SERIAL_SESSION | pkcs11.CKF_RW_SESSION)
	if err != nil {
		return nil, err
	}
	pin := opts.Get("pin")
	if pin != "" {
		err = ctx.Login(session, pkcs11.CKU_USER, pin)
		if err != nil && err != pkcs11.Error(pkcs11.CKR_USER_ALREADY_LOGGED_IN) {
			return nil, err
		}
	}

	// find the key
	id := []byte{0}
	if keyId := opts.Get("id"); keyId != "" {
		id, _ = hex.DecodeString(keyId)
	}

	if len(id) == 0 {
		id = []byte{0}
	}

	keyHandle, err := findHandle(ctx, session, pkcs11.CKO_PRIVATE_KEY, id)
	if err != nil {
		return nil, err
	}

	pubHandle, err := findHandle(ctx, session, pkcs11.CKO_PUBLIC_KEY, id)
	if err != nil {
		return nil, err
	}

	keyType, err := getObjectUintAttr(ctx, session, keyHandle, pkcs11.CKA_KEY_TYPE)
	if err != nil {
		return nil, err
	}

	mechId, err := getObjectUintAttr(ctx, session, keyHandle, pkcs11.CKA_ALLOWED_MECHANISMS)
	if err != nil && err != pkcs11.Error(pkcs11.CKR_ATTRIBUTE_TYPE_INVALID) {
		return nil, err
	}
	log.WithField("sign mechanism", mechId).Debug("found signing mechanism")

	var pubKey crypto.PublicKey
	var signMech *pkcs11.Mechanism
	switch keyType {
	case pkcs11.CKK_ECDSA:
		pubKey, err = loadECDSApub(ctx, session, pubHandle)
		signMech, err = getECDSAmechanism(ctx, slotId, pubKey.(*ecdsa.PublicKey))

	case pkcs11.CKK_RSA:
		signMech = pkcs11.NewMechanism(pkcs11.CKM_RSA_PKCS, nil)
		pubKey, err = loadRSApub(ctx, session, pubHandle)
	default:
		return nil, fmt.Errorf("unsupported key type (%d)", keyType)
	}

	signer := &p11Signer{c: ctx, s: session, h: keyHandle, m: signMech, pub: pubKey}
	signer.label, _ = getObjectStringAttr(ctx, session, keyHandle, pkcs11.CKA_LABEL)

	return signer, nil
}

func (*engine) Id() string {
	return EngineId
}

func getContext(driver string) (*pkcs11.Ctx, error) {
	if ctx, ok := contexts[driver]; ok {
		return ctx, nil
	}

	ctx := pkcs11.New(driver)
	err := ctx.Initialize()
	if err != nil {
		return nil, err
	}
	contexts[driver] = ctx
	return ctx, nil
}

// this mapping is lifted from standard golang crypto/x509.go
var oid2curve = map[string]elliptic.Curve{
	"1.3.132.0.33":        elliptic.P224(),
	"1.2.840.10045.3.1.7": elliptic.P256(),
	"1.3.132.0.34":        elliptic.P384(),
	"1.3.132.0.35":        elliptic.P521(),
}

func getECDSAmechanism(ctx *pkcs11.Ctx, slot uint, pubKey *ecdsa.PublicKey) (*pkcs11.Mechanism, error) {

	var signMech uint
	switch pubKey.Curve.Params().BitSize {
	case 512:
		signMech = pkcs11.CKM_ECDSA_SHA512
	case 384:
		signMech = pkcs11.CKM_ECDSA_SHA384
	case 256:
		signMech = pkcs11.CKM_ECDSA_SHA256
	case 224:
		signMech = pkcs11.CKM_ECDSA_SHA224
	default:
		return nil, fmt.Errorf("unexpected key size curve(%s)", pubKey.Curve.Params().Name)
	}

	prefered := []uint{
		signMech,
		pkcs11.CKM_ECDSA, // fallback -- softhsm only reports this mechanism as available
	}

	for _, m := range prefered {
		mech := []*pkcs11.Mechanism {
			pkcs11.NewMechanism(m, nil),
		}

		_, err := ctx.GetMechanismInfo(slot, mech)
		if err == nil {
			return mech[0], nil
		} else {
			log.WithError(err).Warnf("failed to get mechanism info [%x]", m)
		}
	}

	return nil, fmt.Errorf("token does not support ECDSA sign mechanisms")
}


func loadECDSApub(ctx *pkcs11.Ctx, session pkcs11.SessionHandle, ph pkcs11.ObjectHandle) (*ecdsa.PublicKey, error) {
	templ := []*pkcs11.Attribute {
		{Type: pkcs11.CKA_EC_PARAMS},
		{Type: pkcs11.CKA_EC_POINT},
	}
	if attrs, err := ctx.GetAttributeValue(session, ph, templ); err != nil {
		return nil, err
	} else {
		var oid asn1.ObjectIdentifier
		rest, err := asn1.UnmarshalWithParams(attrs[0].Value, &oid, "")
		log.Debugf("EC oid[%s], rest: %v, err: %v", oid, rest, err)
		curve, found := oid2curve[oid.String()]
		if !found {
			return nil, fmt.Errorf("elliptic curve not found for oid[%s]", oid)
		}

		var pointBytes []byte
		extra, err := asn1.Unmarshal(attrs[1].Value, &pointBytes)
		if err != nil {
			return nil, fmt.Errorf("elliptic curve point is invalid ASN.1")
		}

		if len(extra) > 0 {
			// We weren't expecting extra data
			return nil, fmt.Errorf("unexpected data found when parsing elliptic curve point")
		}
		x,y := elliptic.Unmarshal(curve, pointBytes)

		return &ecdsa.PublicKey{
			Curve: curve,
			X: x,
			Y: y,
		}, nil
	}
}

func loadRSApub(ctx *pkcs11.Ctx, session pkcs11.SessionHandle, ph pkcs11.ObjectHandle) (crypto.PublicKey, error) {
	templ := []*pkcs11.Attribute {
		{Type: pkcs11.CKA_MODULUS},
		{Type: pkcs11.CKA_PUBLIC_EXPONENT},
	}
	if attrs, err := ctx.GetAttributeValue(session, ph, templ); err != nil {
		return nil, err
	} else {

		result := &rsa.PublicKey{
			N: new(big.Int),
		}

		result.N.SetBytes(attrs[0].Value)

		exp := new(big.Int)
		exp.SetBytes(attrs[1].Value)
		if exp.BitLen() > 32 || exp.Sign() < 1{
			return nil, fmt.Errorf("unxpected RSA exponent value")
		}
		result.E = int(exp.Uint64())

		return result, nil
	}
}

func findHandle(ctx *pkcs11.Ctx, session pkcs11.SessionHandle, cls uint, id []byte) (pkcs11.ObjectHandle, error) {
	// find the key
	query := make([]*pkcs11.Attribute, 1)
	query[0] = pkcs11.NewAttribute(pkcs11.CKA_CLASS, cls)

	if id != nil {
		query = append(query, pkcs11.NewAttribute(pkcs11.CKA_ID, id))
	}

	err := ctx.FindObjectsInit(session, query)
	if err != nil {
		return 0, err
	}
	defer func() {
		if e := ctx.FindObjectsFinal(session); e != nil {
			log.Warnf("error: FindObjectFinal(): %v", e)
		}
	}()

	objs, _, err := ctx.FindObjects(session, 1)
	if err != nil {
		return 0, err
	}
	if len(objs) == 0 {
		return 0, fmt.Errorf("key not found")
	}
	return objs[0], nil
}

func getObjectStringAttr(ctx *pkcs11.Ctx, session pkcs11.SessionHandle, obj pkcs11.ObjectHandle, attr_id uint) (string, error) {
	v, err := getObjectAttribute(ctx, session, obj, attr_id)
	return string(v), err
}

func getObjectUintAttr(ctx *pkcs11.Ctx, session pkcs11.SessionHandle, obj pkcs11.ObjectHandle, attr_id uint) (uint, error) {
	v, err := getObjectAttribute(ctx, session, obj, attr_id)
	u, _ := binary.Uvarint(v)
	return uint(u), err
}

func getObjectAttribute(ctx *pkcs11.Ctx, session pkcs11.SessionHandle, obj pkcs11.ObjectHandle, attr_id uint) ([]byte, error) {
	templ := []*pkcs11.Attribute {
		{Type: attr_id},
	}
	attrs, err := ctx.GetAttributeValue(session, obj, templ)
	if err == pkcs11.Error(pkcs11.CKR_ATTRIBUTE_TYPE_INVALID) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return attrs[0].Value, nil
}