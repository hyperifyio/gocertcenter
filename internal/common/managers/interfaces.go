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
	GetRandomManager() IRandomManager
	CreateCertificate(rand io.Reader, template, parent *x509.Certificate, publicKey, privateKey any) ([]byte, error)
	ParseCertificate(certBytes []byte) (*x509.Certificate, error)
	MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte
	MarshalECPrivateKey(key *ecdsa.PrivateKey) ([]byte, error)
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
