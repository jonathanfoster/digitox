package oauth

import (
	"crypto/rsa"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var (
	// DefaultSigningKey is the default token signing key used if none provided.
	DefaultSigningKey *rsa.PrivateKey
	// DefaultVerifyingKey is the default token verifying key used if none provided.
	DefaultVerifyingKey *rsa.PublicKey
)

func init() {
	var err error

	DefaultSigningKey, err = jwt.ParseRSAPrivateKeyFromPEM(defaultSigningKeyBytes)
	if err != nil {
		panic(fmt.Sprintf("error parsing default signing key: %s", err.Error()))
	}

	DefaultVerifyingKey, err = jwt.ParseRSAPublicKeyFromPEM(defaultVerifyingKeyBytes)
	if err != nil {
		panic(fmt.Sprintf("error parsing default verifying key: %s", err.Error()))
	}
}

var defaultSigningKeyBytes = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAkKkzTwlMCHMfDaKwoxMew01tx0TrN2SLl8sE6UG2g3u/P4W6
UqBP5yv+9X82FndDbxpASxRucRhO/PAROPNEo6dU/2FNARtjxoMVK9MTjtLOZKDP
OoV6i9B7EkMVD9Zk8PRzCKoAk8GByWLC2jDuQq6YvW+bIhBcwjz+e2qvZkm06/50
hhcQQE98aaCpE6u6sfFFEx1gfZteR9qj8UEj/XFV/nyRCF9Li4OB2FPyhUg2/1Aa
0DkIwFWqQURB8msNaOECw0hPT0jq0tahuKYqGw8mB9zYlTlHFDq21HIybJ9Z02t5
eE3PJ2gtZO3o4YWSAjjAveNV7gOgBnTgUIyFrQIDAQABAoIBADIy5OEqYr4T5NTA
ffc47VXsiom5ur3wIBi+lKe06/bYfFc4up1tkAyyUbkzObu2CyqEu4bSQjjwrIhN
bkyK7miz6mTsiOI3dPowBqq8hm7rbD+zJfYy14GpCOwfZzGlvkV3Lmv1QloDrlwJ
73/zttpg6BPkpLq/XtDwhYaiUNd3esetchx95C1BNzZwApBLj+Fo4Sb5kDCoRPjK
m6iZwretW3Awz+vVhNDgyL3/LwUHpl7+fCTtwDwc2Gea/NU+88gjEucoqYxb3A5g
GnYcmFPUy0i45NPIN/B1IHaN5+XczgEfJ3aOnWvJMrl+o1xbmnutw0aTLuVDe2Zu
7omrToECgYEAwJIzQvUL5g2FkE3qibFvI5b4146X25bzvi2iwZFIUqs7MIb1eBcS
rSvQ/hkibIOajrVn3u7S9GWDNU9iDsU0nYCQLwwmR06rfvas9AFOqWIDiDsj664R
tvIk7KpbRLmrjYAwCkx2wFlr84vhvrHszU2537JtqXegPT0Xda+9JhECgYEAwE8q
M90YCPSomOFMQNXRsdFdkwEOxTJifarfJ2Ms5ht7BlvED/9anVsy+v3ZXzxuSfNE
E2hvdFp+p/SGRdCplOhC1aBasIa75lpTf8OOALqy4OUlvN/Nq4v/q9Hmf0A3UT7B
IoQui5HHBWJAXR9gUVq7XKO8NtMpaAh9IIHbGd0CgYAsuAbFbshjlRJGL4HeleC8
QAvrasajDMvvhwN7tfQ4lmD5ZO3OBHWm1z0CNO4Eiw8yQrgrUgSVEpnEoHmh+nO8
e6V/929QMdmrczc4trEArq0pTqqJyXN9q3+dofXt4LwQ8Qq26YjOJDXoabxznzfh
eUJHy1Sh/RCuB+jRwIzJMQKBgQCSfHr4MKT7RWobshpElNr7aTCvJrIakgumD/+V
4By2Vx56NHJ/gRKEJJEL0UvAGKcmG3Cym+2yIrCxvTh+e7iBXf5y/Ye3SZpdmFZc
TxifA5f4aXQ6j/v5fVXOir/aFI9Ois/RPAC8fdmtBy9h/+F9dvCbW3mmBWlX/odZ
uLCt3QKBgDGGWwmSLrydxSNHdF3sCVV31w5ixjkgSpE8Ifk5is6IQmgl+I+CKKXz
Y4MplOmBzKlQcYuh5qP8ywjiICWp+EbyfX0OiN81qNAmxGT8e0x95Vi+p0bm8KDo
S9aJe/2fMg1/eM3KlcKBnmhqHfEcNiIXbdyjSXqi9qyhKkrFxMx3
-----END RSA PRIVATE KEY-----`)

var defaultVerifyingKeyBytes = []byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAkKkzTwlMCHMfDaKwoxMe
w01tx0TrN2SLl8sE6UG2g3u/P4W6UqBP5yv+9X82FndDbxpASxRucRhO/PAROPNE
o6dU/2FNARtjxoMVK9MTjtLOZKDPOoV6i9B7EkMVD9Zk8PRzCKoAk8GByWLC2jDu
Qq6YvW+bIhBcwjz+e2qvZkm06/50hhcQQE98aaCpE6u6sfFFEx1gfZteR9qj8UEj
/XFV/nyRCF9Li4OB2FPyhUg2/1Aa0DkIwFWqQURB8msNaOECw0hPT0jq0tahuKYq
Gw8mB9zYlTlHFDq21HIybJ9Z02t5eE3PJ2gtZO3o4YWSAjjAveNV7gOgBnTgUIyF
rQIDAQAB
-----END PUBLIC KEY-----`)
