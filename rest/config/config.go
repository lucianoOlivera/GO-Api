package config

import (
	"fmt"
	"os"

	"github.com/eduardogpg/gonv"
)

type DataBaseconfig struct {
	Name      string
	Passwpord string
	Host      string
	Port      int
	DataBase  string
}

var database *DataBaseconfig

func init() {
	os.Setenv("USERNAME", "root")
	database = &DataBaseconfig{}
	database.Name = gonv.GetStringEnv("USERNAME", "root")
	database.Passwpord = gonv.GetStringEnv("PASSWORD", "")
	database.Host = gonv.GetStringEnv("HOST", "localhost")
	database.Port = gonv.GetIntEnv("PORT", 3306)
	database.DataBase = gonv.GetStringEnv("DATABASE", "project_go_web")
}

//user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
func (this *DataBaseconfig) url() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", this.Name, this.Passwpord, this.Host, this.Port, this.DataBase)
}

func GetUrl() string {
	return database.url()
}
