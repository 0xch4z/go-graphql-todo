package config

import "fmt"

// DBConfig describes the config for the DB
type DBConfig struct {
	Dialect  string
	username string
	password string
	name     string
	charset  string
}

// URI returns the URI for a given DBConfig
func (d *DBConfig) URI() string {
	return fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=true",
		d.username,
		d.password,
		d.name,
		d.charset,
	)
}
