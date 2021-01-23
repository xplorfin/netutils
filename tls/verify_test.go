package tls

import (
	"testing"
)

func TestChain(t *testing.T) {
	rootParams := DefaultTlsParams()
	rootParams.IsCa = true
	cert, err := MakeCertificate(rootParams)
	if err != nil {
		t.Error(err)
	}
	rootCert := cert.Certificate

	dca, err := cert.MakeDca()
	if err != nil {
		t.Error(err)
	}
	dcaIsValid, err := VerifyDCA(rootCert, dca.Certificate)
	if !dcaIsValid {
		t.Error(err)
	}

	serverCert, err := dca.MakeServerCertificate()
	if err != nil {
		t.Error(err)
	}
	isValid, err := VerifyLow(rootCert, dca.Certificate, serverCert.Certificate)
	if !isValid {
		t.Error(err)
	}
}
