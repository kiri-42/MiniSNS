package configs

import (
	"database/sql"
	"github.com/kelseyhightower/envconfig"
	"log"
	"sync"
)

var (
	conf Config
	once sync.Once
)

type Config struct {
	Server ServerConfig
	DB     DBConfig
}

type ServerConfig struct {
	Port int `default:"1323"`
}

type DBConfig struct {
	Driver     string `default:"mysql"`
	DataSource string `default:"root:@(db:3306)/app"`
}

func Get() Config {
	once.Do(func() {
		if err := envconfig.Process("server", &conf.Server); err != nil {
			log.Fatal(err.Error())
		}
		if err := envconfig.Process("db", &conf.DB); err != nil {
			log.Fatal(err.Error())
		}
	})
	return conf
}

func GetDB() (*sql.DB, error) {
	conf := Get()

	db, err := sql.Open(conf.DB.Driver, conf.DB.DataSource)
	if err != nil {
		return nil, err
	}

	return db, nil
}
