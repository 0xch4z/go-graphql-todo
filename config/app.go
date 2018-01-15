package config

// AppConfig represents config components necessary
// for reference throughout the app.
type AppConfig struct {
	DB     *DBConfig
	Crypto *CryptoConfig
}
