package config

import (
	"fmt"
	_ "fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBHOST string  `env:"DB_HOST"`
	DBPORT string  `env:"DB_PORT"`
	DBUSER string  `env:"DB_USER"`
	DBPASSWORD string `env:"DB_PASSWORD"`
	DBNAME string `env:"DB_NAME"`
	DBSSLMODE string `env:"DB_SSLMODE"`
	PORT string `env:"PORT"`
}

func LoadConfig() *DBConfig{
	
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Cant get env from %v", err)
	}
  var cfg_object DBConfig;
	cfg := &cfg_object;
  
	if err = env.Parse(cfg) ; err != nil { 
     log.Fatalf("Fail to get env %v", err)
	 }

	fmt.Println(cfg)
	return cfg
}



