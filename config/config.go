package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
)

var (
	//GRPCPort - network port for grpc
	GRPCPort string
	//GRPCNetwork - network for grpc
	GRPCNetwork string
	//DBName - database name
	DBName string
	//DBUserName - dbusername
	DBUserName string
)

const (
	modeKey     = "MODE"
	devMode     = "development"
	grpcPort    = "GRPCPORT"
	grpcNetwork = "GRPCNETWORK"
	dbname      = "DBNAME"
	dbusername  = "DBUSERNAME"
)

//GetEnvironment - this function returns mode string of the os environment or "development" mode if empty or not defined
func GetEnvironment() string {
	var env string
	if env = os.Getenv(modeKey); env == "" {
		return devMode
	}
	return env
}

//IsDevelopmentEnv - this function try to get mode environment and check it is development
func IsDevelopmentEnv() bool { return GetEnvironment() == devMode }

//LoadEnvironment - function to load env file and get all required variables from the os environment
func LoadEnvironment() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = godotenv.Load(path.Join(wd, strings.ToLower(GetEnvironment())+".env"))
	if err != nil {
		log.Warning(err.Error())
	}
	GRPCPort = os.Getenv(grpcPort)
	if GRPCPort == "" {
		log.Fatal("grpc port must be set")
	}
	GRPCNetwork = os.Getenv(grpcNetwork)
	if GRPCNetwork == "" {
		log.Fatal("grpc network must be set")
	}
	DBName = os.Getenv(dbname)
	if DBName == "" {
		log.Fatal("db name must be set")
	}
	DBUserName = os.Getenv(dbusername)
	if DBUserName == "" {
		log.Fatal("db user must be set")
	}
}
