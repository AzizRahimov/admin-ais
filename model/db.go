package model

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"log"
	"os"
	"time"
)


//type Application struct {
//
//	Db   Db
//}


type Config struct {
	Host string
	Port string
	Username string
	Password string
	DBName string
	SSLMode string
}

var configPath = "./config.json"

func LoadConfig()(config *Config){
	data,err:=ioutil.ReadFile(configPath)
	if err!=nil{
		log.Fatal(err)
	}
	config=&Config{}
	err=json.Unmarshal(data,&config)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("не пустой CONFIG", config)
	return config
}

func DBConnection() (*gorm.DB, error) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},

	)

	config := LoadConfig()
	fmt.Println(config)
	cfg := config
	///connStr := "postgres://"+ config.Db.User + ":" + config.Db.Pass + "@"+ config.Db.Host + "/"+ config.Db.Dbname
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	//


	return gorm.Open(postgres.Open(connStr), &gorm.Config{Logger: newLogger})

}

