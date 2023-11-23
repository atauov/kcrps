package dashboard

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1658),
		Subject: pkix.Name{
			Organization:  []string{"Anuar"},
			Country:       []string{"KZ"},
			Province:      []string{"ASTANA"},
			Locality:      []string{"Astana"},
			StreetAddress: []string{"Esil"},
			PostalCode:    []string{"010000"},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	pub := &priv.PublicKey

	// Sign the certificate
	certificate, _ := x509.CreateCertificate(rand.Reader, cert, cert, pub, priv)

	certBytes := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certificate})
	keyBytes := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})

	// Generate a key pair from your pem-encoded cert and key ([]byte).
	x509Cert, _ := tls.X509KeyPair(certBytes, keyBytes)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{x509Cert},
	}

	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, //1 MB ,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		TLSConfig:      tlsConfig,
	}
	return s.httpServer.ListenAndServeTLS("", "")
}
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
