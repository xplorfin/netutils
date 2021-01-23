package tls

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/big"
	"net"
	"strings"
	"time"
)

func MakeSubject() pkix.Name {
	return pkix.Name{
		Organization:  []string{"Acme, Inc"},
		Locality:      []string{"Nowhere"},
		Province:      []string{"USA"},
		StreetAddress: []string{"18 main street"},
		PostalCode:    []string{"014611"},
	}
}

// params for creating a tls cert
type TlsParams struct {
	// Comma-separated hostnames and IPs to generate a certificate for
	Host string
	// Creation date formatted as Jan 1 15:04:05 2011
	ValidFrom time.Time
	// Duration that certificate is valid for
	ValidFor time.Duration
	// whether this cert should be its own Certificate Authority
	IsCa bool
	// Size of RSA key to generate. Ignored if EcdsaCurve is set
	RsaBits int
	// ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521
	EcdsaCurve string
	// Generate an Ed25519 key
	Ed25519 bool
}

// copy a tls params struct
func (t *TlsParams) Copy() TlsParams {
	return TlsParams{
		Host:       t.Host,
		ValidFrom:  t.ValidFrom,
		ValidFor:   t.ValidFor,
		IsCa:       t.IsCa,
		RsaBits:    t.RsaBits,
		EcdsaCurve: t.EcdsaCurve,
		Ed25519:    t.Ed25519,
	}
}

var (
	Day   = time.Hour * 24
	Week  = Day * 7
	Month = Week * 4
	Year  = Month * 12
)

// default params for creating a tls cert
// a function for immutability
func DefaultTlsParams() TlsParams {
	return TlsParams{
		Host:       "localhost",
		ValidFrom:  time.Now(),
		ValidFor:   Year,
		IsCa:       false,
		RsaBits:    2048,
		EcdsaCurve: "",
		Ed25519:    false,
	}
}

// params returned by certificate generator
type TlsCert struct {
	CertType string
	// certificate.pem (public key)
	PublicKey string
	// key.pem (private key)
	PrivateKey string
	// Certificate
	Certificate *x509.Certificate
	// private key
	Key interface{}
	// parameters (
	Params TlsParams
}

// generate a dca certificate from a rootc
func (t TlsCert) MakeDca() (TlsCert, error) {
	if !t.Certificate.IsCA {
		return TlsCert{}, fmt.Errorf("certificate must be a root certificate to issue a dca")
	}
	priv, err := generatePrivateKey(t.Params)
	if err != nil {
		return TlsCert{}, err
	}

	dcaTemplate := x509.Certificate{
		SerialNumber:          generateSerialNumber(),
		Subject:               MakeSubject(),
		NotBefore:             t.Params.ValidFrom,
		NotAfter:              t.Params.ValidFrom.Add(t.Params.ValidFor),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
		MaxPathLenZero:        false,
		MaxPathLen:            1,
	}
	AddHostToTemplate(t.Params.Host, &dcaTemplate)
	generatedCert, err := genTlsCert(t.Params, &dcaTemplate, t.Certificate, publicKey(priv), t.Key)
	if err != nil {
		return TlsCert{}, err
	}
	generatedCert.Key = priv
	return generatedCert, nil

}

// generate a dca certificate from a rootc
func (t TlsCert) MakeServerCertificate() (TlsCert, error) {
	if !t.Certificate.IsCA {
		return TlsCert{}, fmt.Errorf("certificate must be a root certificate to issue a dca")
	}
	serialNumber := generateSerialNumber()
	priv, err := generatePrivateKey(t.Params)
	if err != nil {
		return TlsCert{}, err
	}

	serverTemplate := x509.Certificate{
		SerialNumber:   serialNumber,
		NotBefore:      t.Params.ValidFrom,
		NotAfter:       t.Params.ValidFrom.Add(t.Params.ValidFor),
		KeyUsage:       x509.KeyUsageCRLSign,
		ExtKeyUsage:    []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IsCA:           false,
		MaxPathLenZero: true,
	}
	AddHostToTemplate(t.Params.Host, &serverTemplate)
	generatedCert, err := genTlsCert(t.Params, &serverTemplate, t.Certificate, publicKey(priv), t.Key)
	if err != nil {
		return TlsCert{}, err
	}
	generatedCert.Key = priv
	return generatedCert, nil
}

// generate a private key of a given type specified in params
func generatePrivateKey(params TlsParams) (priv interface{}, err error) {
	switch params.EcdsaCurve {
	case "":
		if params.Ed25519 {
			_, priv, err = ed25519.GenerateKey(rand.Reader)
		} else {
			priv, err = rsa.GenerateKey(rand.Reader, params.RsaBits)
		}
	case P224:
		priv, err = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case P256:
		priv, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case P384:
		priv, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case P521:
		priv, err = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		return priv, fmt.Errorf("unrecognized elliptic curve: %q", params.EcdsaCurve)
	}
	if err != nil {
		return priv, fmt.Errorf("failed to generate private key: %v", err)
	}
	return priv, nil
}

// generate a tls cert from params
func genTlsCert(params TlsParams, template, parent *x509.Certificate, publicKey interface{}, privateKey interface{}) (cert TlsCert, err error) {
	derBytes, err := x509.CreateCertificate(rand.Reader, template, parent, publicKey, privateKey)

	if err != nil {
		return TlsCert{}, fmt.Errorf("failed to create certificate: %v", err)
	}

	certOut := bytes.NewBufferString("")
	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return TlsCert{}, fmt.Errorf("failed to write data to cert.pem: %v", err)
	}

	xCert, err := x509.ParseCertificate(derBytes)
	if err != nil {
		return TlsCert{}, err
	}

	keyOut := bytes.NewBufferString("")
	privBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return TlsCert{}, fmt.Errorf("unable to marshal private key: %v", err)
	}
	if err := pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		return TlsCert{}, fmt.Errorf("failed to write data to key.pem: %v", err)
	}

	return TlsCert{
		PublicKey:   certOut.String(),
		PrivateKey:  keyOut.String(),
		Certificate: xCert,
		Key:         privateKey,
		Params:      params,
	}, nil
}

// generate a serial number for use in certs
func generateSerialNumber() *big.Int {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		panic(fmt.Errorf("failed to generate serial number: %v", err))
	}
	return serialNumber
}

func AddHostToTemplate(host string, template *x509.Certificate) {
	hosts := strings.Split(host, ",")
	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}
}

func MakeCertificateDefault() (cert TlsCert, err error) {
	return MakeCertificate(DefaultTlsParams())
}

// create a public key from a private key
func publicKey(priv interface{}) interface{} {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	case ed25519.PrivateKey:
		return k.Public().(ed25519.PublicKey)
	default:
		return nil
	}
}

// curve types for rsa keys
const (
	P224 = "P224"
	P256 = "P256"
	P384 = "P384"
	P521 = "P521"
)

// Create a tls certificate
// adapted from https://golang.org/src/crypto/tls/generate_cert.go
func MakeCertificate(params TlsParams) (cert TlsCert, err error) {
	if len(params.Host) == 0 {
		return cert, errors.New("missing required host parameter")
	}

	priv, err := generatePrivateKey(params)
	if err != nil {
		return TlsCert{}, err
	}

	// ECDSA, ED25519 and RSA subject keys should have the DigitalSignature
	// KeyUsage bits set in the x509.Certificate template
	keyUsage := x509.KeyUsageDigitalSignature
	// Only RSA subject keys should have the KeyEncipherment KeyUsage bits set. In
	// the context of TLS this KeyUsage is particular to RSA key exchange and
	// authentication.
	if _, isRSA := priv.(*rsa.PrivateKey); isRSA {
		keyUsage |= x509.KeyUsageKeyEncipherment
	}

	template := x509.Certificate{
		SerialNumber: generateSerialNumber(),
		Subject:      MakeSubject(),
		NotBefore:    params.ValidFrom,
		NotAfter:     params.ValidFrom.Add(params.ValidFor),

		KeyUsage:              keyUsage,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	AddHostToTemplate(params.Host, &template)

	if params.IsCa {
		template.IsCA = true
		template.KeyUsage |= x509.KeyUsageCertSign
		template.KeyUsage |= x509.KeyUsageCRLSign
		template.MaxPathLen = 2
		template.Subject.CommonName = "Root CA"
	}

	return genTlsCert(params, &template, &template, publicKey(priv), priv)
}

// verify a tls key parir is valid
func VerifyCertificate(cert TlsCert) (isValid bool, err error) {
	_, err = tls.X509KeyPair([]byte(cert.PublicKey), []byte(cert.PrivateKey))
	return err == nil, err
}

func VerifyKeyPair(rsaPublicKey, rsaPrivateKey string) (isValid bool, err error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(rsaPrivateKey))
	if err != nil {
		return isValid, err
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(rsaPublicKey))
	if err != nil {
		return isValid, err
	}
	return key.PublicKey.Equal(pubKey), err
}
