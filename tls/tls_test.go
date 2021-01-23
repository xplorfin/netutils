package tls

import (
	"github.com/brianvoe/gofakeit/v5"
	"github.com/davecgh/go-spew/spew"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/r3labs/diff/v2"
	"sync"
	"testing"
)

var allBoolOptions = []bool{true, false}

func TestMakeCertificateDefault(t *testing.T) {
	cert, err := MakeCertificateDefault()
	if err != nil {
		t.Error(err)
	}

	VerifyKeyPairTest(t, cert.PublicKey, cert.PrivateKey)
}

func TestCopy(t *testing.T) {
	for _, param := range MakeTestParams() {
		change, err := diff.Diff(param, param.Copy())
		if err != nil {
			t.Error(err)
		}
		if len(change) > 0 {
			t.Errorf("expected copy to be equal to original, got %s", spew.Sdump(change))
		}
	}
}

func TestMakeCertificate(t *testing.T) {
	testParams := MakeTestParams()
	sliceLength := len(testParams)
	var wg sync.WaitGroup
	wg.Add(sliceLength)
	for i := 0; i < sliceLength; i++ {
		go func(i int) {
			defer wg.Done()
			val := testParams[i]
			cert, err := MakeCertificate(val)
			if err != nil {
				t.Error(err)
			}
			isValid, err := VerifyCertificate(cert)
			if !isValid || err != nil {
				t.Error(err)
			}
		}(i)
	}
	wg.Wait()
}

func TestInvalidParams(t *testing.T) {
	params := DefaultTlsParams()
	params.Host = ""
	_, err := MakeCertificate(params)
	if err == nil {
		t.Error(err)
	}

	params = DefaultTlsParams()
	params.RsaBits = -1
	_, err = MakeCertificate(params)
	if err == nil {
		t.Error(err)
	}

	params = DefaultTlsParams()
	params.EcdsaCurve = gofakeit.BitcoinAddress()
	_, err = MakeCertificate(params)
	if err == nil {
		t.Error(err)
	}
}

// test invalid cases.
func TestVerifyKeyPairInvalid(t *testing.T) {
	// make sure keys that don't match return false
	isValid, _ := VerifyKeyPair(MakeTestCertificate(t).PublicKey, MakeTestCertificate(t).PrivateKey)
	if isValid {
		t.Errorf("expected error when trying to match two different keys")
	}
	// make sure switched private key and public key don't match
	validCert := MakeTestCertificate(t)
	isValid, _ = VerifyKeyPair(validCert.PrivateKey, validCert.PublicKey)
	if isValid {
		t.Errorf("expected error when private and public key are switched")
	}

	// test bs private key
	isValid, _ = VerifyKeyPair(gofakeit.HipsterParagraph(40, 10, 5, " "), validCert.PrivateKey)
	if isValid {
		t.Errorf("expected error when private key invalid")
	}

	// test bs public key
	isValid, _ = VerifyKeyPair(validCert.PrivateKey, gofakeit.HipsterParagraph(40, 10, 5, " "))
	if isValid {
		t.Errorf("expected error when private key invalid")
	}

}

// make sure a valid key pair supplied by const is valid
func TestVerifyKeyPairValid(t *testing.T) {
	isValid, err := VerifyKeyPair(rsaPublicKey, rsaPrivateKey)
	if err != nil {
		t.Error(err)
	}
	if !isValid {
		t.Errorf("expected cert with public key %s and private key %s to be valid", rsaPublicKey, rsaPrivateKey)
	}
}

// all combos here should create valid certs
// there's go to be a better way to do this
func MakeTestParams() (testParams []TlsParams) {
	for _, host := range []string{gofakeit.DomainName(), gofakeit.IPv4Address(), gofakeit.IPv6Address()} {
		domainNameParams := DefaultTlsParams()
		domainNameParams.Host = host
		// try out each of the curves
		for _, curve := range []string{"", P224, P256, P384, P521} {
			// make sure we're not mutating top level params
			curveParams := domainNameParams.Copy()
			curveParams.EcdsaCurve = curve
			for _, isCa := range allBoolOptions {
				// make sure we're not mutating top level params
				isCaParams := curveParams.Copy()
				isCaParams.IsCa = isCa
				for _, ed25519 := range allBoolOptions {
					ed25519Params := isCaParams.Copy()
					ed25519Params.Ed25519 = ed25519
					for _, bits := range []int{DefaultTlsParams().RsaBits, tools.RandomInt(60, 100)} {
						bitParams := ed25519Params.Copy()
						bitParams.RsaBits = bits
						testParams = append(testParams, ed25519Params)
					}
				}
			}
		}
	}

	return testParams
}

// valid rsa public/private key pair from https://bit.ly/3hZTmYc
const (
	rsaPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEoQIBAAKCAQBt6fZgG7bHYXSAno1Tf0mPZeciBhL19s7AjnFVlH+HXDs81SSf
lyFCH87qm2BArXHXZs2bId0znguMlSoTBVRiQgY2mUwRIhj8a7aAMcFNPmS2cTwm
5L2JvybVDF7okt2GD6rVyYTiD+h1ekw97rNUqNUDK+e76KweQvliuYtx/xfQrhSH
0wz2hR6AnDGm28sTheN4oGrdR0lp7k3u1KFWcREp6LiXigphb7XMpd2UAxsIDz5i
9pkRXcDFEatBpKl6VgR1S57kB3HVxP1QE0pKQF/stuhU6fIkfliXnVV/CxG1poM5
Dvz7nQn0jjpGXE7F9AHPFF79y0fZKUY4yqmDAgMBAAECggEASuAX427du8Mq3zej
ZN8QWGx94NKsdfxU6h3fHQKVQbvV6uH0GfeVQ3txtKZ2EnlVVPyGUgjmrHQcv/8Z
c22tB2ac3vLdO8qzpLkn0PqUCS/Y4eQPqxsl90FNjdLokJ0D5YCkyxCFwo49uFHA
wGvspF6DBa8SJRMitVbAlr15PgUyFfX5oC1gOBWj9mK+5QQGaiz9HRNZyBuowP9x
VLIrZLtlLM14CvwdU6V79DSWe1GtXc6zb6S4bONq24b79Ub9FV+9RQrD5K6rmuim
lRvlF5SX51KE/wKIKUx4TuVzjGG0PYpSHhnnIUniWu/p9XjvSL2Vqr7/S/GOFk9y
ZPrpSQKBgQDXYyPJFJE0T4DJe0fE+R/zj0pjML4WrGja1A3HyZZ3e3JpsGAbToUs
X8VwT9yXoiZzj6xt6OXSo45T3IrhcuWxd1CMvj1fgxoYlOzsq2XHDwwVzq5N2Ke3
P6OD+W2DtY352RYXJ2KtPoNZOkejSHDiswm8XJr6fOKRdsSRT/e6zwKBgQCCo5EZ
B3kNW5u2de0+OlW/wM+m6r32wC3oYVcBn5Jf1NH/U31h01NAO8I1PRyEbhi4j5dz
xXCQMqV+RAy/A7TnQS4nPl0UqCXA6s7HVXg2eLdMzp6Jm1iah1bN+E+4utqM1hk8
Dgasmgq1iKsUuMisvU4uKis2miswHt76JA5DDQKBgGjNm68PK+xpNwBS1UQ5+Fsa
ARcss4Hy2H6KKj5pj6aJ0c0tfkYrOc+yti6FHZBG3TDj2wIMDjAlV27k5Er5Dl0A
8pfZRaHA+CS36mTqrXZjkvzVeaj1X/5hn93qs2ggInpNMFuJ1ZD41w7Gte70o8Eb
XwRhhyOVOuWPBeyzHZavAoGAGTQ9djq+3Bjkfdtanjra+FfWuDlp1QVW1hKRmrqS
nvKMYVpWQl1nHmlpGqRjsBkdo93wNmHNScS7sRSn8OJiMIuev+uEQcv/HK0wn7yZ
qMi5dJQYeiwCeC3MTYiuuNE0AR/9VlzOZNaDYmqvtxu/e7Q6NSXlmG8+Ddam5lO2
fLECgYBNB6h49zPWtk/Ot9wts07oSc4Zj+17cBiV/nnQ2H65rSeRnC4ObYS6nc/A
9uUi5wQcb0qKwOh9LHhPjeY3NfsroUknOLSHf+zQfbF1E2S1i/I5SbRKJO+Mxq75
h382gDA/J7uXS3eFe6AaDfNUUxH6NVUc4nzy+XsL6r/6D+L5gA==
-----END RSA PRIVATE KEY-----`

	rsaPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBITANBgkqhkiG9w0BAQEFAAOCAQ4AMIIBCQKCAQBt6fZgG7bHYXSAno1Tf0mP
ZeciBhL19s7AjnFVlH+HXDs81SSflyFCH87qm2BArXHXZs2bId0znguMlSoTBVRi
QgY2mUwRIhj8a7aAMcFNPmS2cTwm5L2JvybVDF7okt2GD6rVyYTiD+h1ekw97rNU
qNUDK+e76KweQvliuYtx/xfQrhSH0wz2hR6AnDGm28sTheN4oGrdR0lp7k3u1KFW
cREp6LiXigphb7XMpd2UAxsIDz5i9pkRXcDFEatBpKl6VgR1S57kB3HVxP1QE0pK
QF/stuhU6fIkfliXnVV/CxG1poM5Dvz7nQn0jjpGXE7F9AHPFF79y0fZKUY4yqmD
AgMBAAE=
-----END PUBLIC KEY-----`
)

// wrapper around tls cert creator
func MakeTestCertificate(t *testing.T) (cert TlsCert) {
	cert, err := MakeCertificateDefault()
	if err != nil {
		t.Error(err)
	}
	return cert
}

// verify an dsa pub private key pair during a test
func VerifyKeyPairTest(t *testing.T, rsaPublicKey, rsaPrivateKey string) {
	isValid, err := VerifyKeyPair(rsaPublicKey, rsaPrivateKey)
	if err != nil {
		t.Error(err)
	}
	if !isValid {
		t.Error(err)
	}
}
