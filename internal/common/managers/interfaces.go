// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package managers

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"io"
	"math/big"
	"net"
	"net/http"

	swagger "github.com/davidebianchi/gswagger"
	"github.com/gorilla/mux"
)

// IRandomManager describes operations to create random values
type IRandomManager interface {
	CreateBigInt(max *big.Int) (*big.Int, error)
}

// ICertificateManager describes operations to manage x509 certificates.
// This is intended to wrap low level external library operations for easier
// testing by using mocks. Any higher level operations shouldn't be implemented
// inside it.
type ICertificateManager interface {

	// GetRandomManager returns a random number manager
	GetRandomManager() IRandomManager

	// CreateCertificate wraps up a call to x509.CreateCertificate
	//  - rand io.Reader
	//  - template  *x509.Certificate
	//  - parent *x509.Certificate
	//  - publicKey *rsa.PublicKey, *ecdsa.PublicKey or ed25519.PublicKey
	//  - privateKey crypto.Signer with a supported publicKey
	// Returns a new certificate in DER format []byte or an error
	CreateCertificate(rand io.Reader, template, parent *x509.Certificate, publicKey, privateKey any) ([]byte, error)

	// ParseCertificate wraps up a call to x509.ParseCertificate to parse a single certification
	//  - der []byte: ASN.1 DER data
	// Returns *x509.Certificate or an error
	ParseCertificate(der []byte) (*x509.Certificate, error)

	// MarshalPKCS1PrivateKey wraps up a call to x509.MarshalPKCS1PrivateKey
	//  - key *rsa.PrivateKey: RSA private key
	// Returns PKCS #1, ASN.1 DER form []byte, e.g. "RSA PRIVATE KEY" PEM block or an error
	MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte

	// MarshalECPrivateKey wraps up a call to x509.MarshalECPrivateKey
	//  - key *ecdsa.PrivateKey
	// Returns SEC 1, ASN.1 DER form []byte, e.g. "EC PRIVATE KEY" PEM block or an error
	MarshalECPrivateKey(key *ecdsa.PrivateKey) ([]byte, error)

	// MarshalPKCS8PrivateKey wraps up a call to x509.MarshalPKCS8PrivateKey
	//  - key *rsa.PrivateKey, *ecdsa.PrivateKey, ed25519.PrivateKey (not a pointer), or *ecdh.PrivateKey
	// Returns PKCS #8, ASN.1 DER form []byte e.g. "PRIVATE KEY" PEM block or an error
	MarshalPKCS8PrivateKey(key any) ([]byte, error)
}

type IServerManager interface {
	Serve(l net.Listener) error
	Shutdown() error
}

type ISwaggerManager interface {
	GenerateAndExposeOpenapi() error

	AddRoute(method string, path string, handler http.HandlerFunc, schema swagger.Definitions) (*mux.Route, error)
}
