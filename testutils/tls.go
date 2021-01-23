package testutils

import (
	"github.com/Flaque/filet"
	"github.com/xplorfin/netutils/tls"
	"testing"
)

// create a temporary certificate key/value pair that's deleted after the test that's delted
func TemporaryCert(t *testing.T) (certFile string, keyFile string) {
	dir := filet.TmpDir(t, "")
	cert, err := tls.MakeCertificateDefault()
	if err != nil {
		t.Error(err)
	}
	certFileLoc := filet.TmpFile(t, dir, cert.PublicKey)
	keyFileLoc := filet.TmpFile(t, dir, cert.PrivateKey)
	return certFileLoc.Name(), keyFileLoc.Name()
}

// create a cert chain we validate against
func TemporaryCertInChain(t *testing.T) (chainFile string, serverFile string, keyFile string) {
	chainDir := filet.TmpDir(t, "")
	params := tls.DefaultTlsParams()
	params.IsCa = true
	chain, err := tls.MakeCertificate(params)
	if err != nil {
		t.Error(err)
	}
	chainFileLoc := filet.TmpFile(t, chainDir, chain.PublicKey)

	dir := filet.TmpDir(t, "")
	serverCert, err := chain.MakeServerCertificate()
	if err != nil {
		t.Error(err)
	}
	serverCertLoc := filet.TmpFile(t, dir, serverCert.PublicKey)
	keyFileLoc := filet.TmpFile(t, dir, serverCert.PrivateKey)
	return chainFileLoc.Name(), serverCertLoc.Name(), keyFileLoc.Name()
}
