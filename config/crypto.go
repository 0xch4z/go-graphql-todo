package config

import (
	"crypto/rsa"
	"io/ioutil"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	privateKeyPath = "keys/app.rsa"
	publicKeyPath  = "keys/app.rsa.pub"
)

// CryptoConfig describes the crypto config
type CryptoConfig struct {
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}

func getCryptoConfig() *CryptoConfig {
	pubBytes, pubErr := ioutil.ReadFile(publicKeyPath)
	privBytes, prbErr := ioutil.ReadFile(privateKeyPath)

	pub, pukErr := jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	priv, prkErr := jwt.ParseRSAPrivateKeyFromPEM(privBytes)

	fatalConfigError(pubErr, prbErr, pukErr, prkErr)

	return &CryptoConfig{pub, priv}
}
