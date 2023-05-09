package flag

import (
	"crypto/tls"
)

func init() {
	// support official IANA names as well on go1.14
	ciphers["TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256"] = tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256
	ciphers["TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256"] = tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256
}
