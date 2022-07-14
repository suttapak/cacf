package configs

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//function create gorm database recive *viper.Viper  return *gorm.DB
func CreatePostgres(config *viper.Viper) *gorm.DB {
	//get env config
	url := config.GetString("POSTGRES_URL")
	//create gorm db
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
