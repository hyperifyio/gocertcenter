// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package managers

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io"
	"math/big"
	"net"
	"net/http"
	"os"

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

	// ParsePKCS8PrivateKey wraps up a call to x509.ParsePKCS8PrivateKey
	ParsePKCS8PrivateKey(der []byte) (any, error)

	// ParsePKCS1PrivateKey wraps up a call to x509.ParsePKCS1PrivateKey
	ParsePKCS1PrivateKey(der []byte) (*rsa.PrivateKey, error)

	// ParseECPrivateKey wraps up a call to x509.ParseECPrivateKey
	ParseECPrivateKey(der []byte) (*ecdsa.PrivateKey, error)

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

	// EncodePEMToMemory wraps a call to pem.EncodeToMemory
	//  - b *pem.Block:
	// Returns []byte or nil
	EncodePEMToMemory(b *pem.Block) []byte

	// DecodePEM wraps a call to pem.Decode
	//  - data []byte:
	DecodePEM(data []byte) (p *pem.Block, rest []byte)
}

type IServerManager interface {
	Serve(l net.Listener) error
	Shutdown() error
}

type ISwaggerManager interface {
	GenerateAndExposeOpenapi() error

	AddRoute(method string, path string, handler http.HandlerFunc, schema swagger.Definitions) (*mux.Route, error)
}

// IFileManager implements a file system manager
type IFileManager interface {

	// ReadBytes reads bytes from a file
	//   - fileName string: The file where to read
	//
	// Returns the bytes read or nil
	ReadBytes(fileName string) ([]byte, error)

	// SaveBytes saves bytes to a file, including creating any parent directories.
	//   - fileName string: The file where to save
	//   - data []byte: The data to save
	//   - filePerms os.FileMode: Permissions for file
	//   - dirPerms os.FileMode: Permissions for directories
	//
	// Returns nil or error
	SaveBytes(
		fileName string,
		data []byte,
		filePerms, dirPerms os.FileMode,
	) error
}
