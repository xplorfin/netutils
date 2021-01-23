package tls

import (
	"crypto/x509"
	"fmt"
)

// verify a dca certificate against it's parent
func VerifyDCA(root, dca *x509.Certificate) (bool, error) {
	roots := x509.NewCertPool()
	roots.AddCert(root)
	opts := x509.VerifyOptions{
		Roots: roots,
	}

	if _, err := dca.Verify(opts); err != nil {
		return false, fmt.Errorf("failed to verify certificate: " + err.Error())
	}
	return true, nil
}

// verify a server certificate against it's chain
func VerifyLow(root, DCA, child *x509.Certificate) (bool, error) {
	roots := x509.NewCertPool()
	inter := x509.NewCertPool()
	roots.AddCert(root)
	inter.AddCert(DCA)
	opts := x509.VerifyOptions{
		Roots:         roots,
		Intermediates: inter,
	}

	if _, err := child.Verify(opts); err != nil {
		return false, fmt.Errorf("failed to verify certificate: " + err.Error())
	}
	return true, nil
}

// verify a server certificate against it's chain
func VerifyLowNoDca(root, child *x509.Certificate) (bool, error) {
	roots := x509.NewCertPool()
	inter := x509.NewCertPool()
	roots.AddCert(root)
	opts := x509.VerifyOptions{
		Roots:         roots,
		Intermediates: inter,
	}

	if _, err := child.Verify(opts); err != nil {
		return false, fmt.Errorf("failed to verify certificate: " + err.Error())
	}
	return true, nil
}
