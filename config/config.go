package config

import (
	"errors"
	"flag"
	"strings"

	"github.com/spf13/viper"
)

// DBConfig needs to be exported as it is the return type of GetDbConfig, which is also exported
type DBConfig struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	Sslmode  string
	Timeout  int
}

// ParseArgs needs to be exported as it is called from main.go
func ParseArgs() (string, int) {
	var fileName string
	// Parse the postrgres configuration file name and path. if not deifned the default is "postgresql.cfg" from /cmd
	flag.StringVar(&fileName, "file", "../config/postgresql.toml", "Path of postgresql config file to be parsed.")
	var portNumber int
	// Parse the port number that the server uses to listen and serve. If none is defined the default is 8080
	flag.IntVar(&portNumber, "port", 8080, "Port on which the server will listen and serve.")
	flag.Parse()
	return fileName, portNumber
}

// GetDbConfig needs to be exported as it is called from outside of the config package
func GetDbConfig(fileName string) (DBConfig, error) {

	// prepare the path, filname and extension variables for viper
	s := strings.Split(fileName, ".")
	if s[len(s)-1] != "toml" {
		var ErrWrongFormat = errors.New("err: Unexpected extension. File must be .toml")
		var d = DBConfig{}
		return d, ErrWrongFormat
	}
	s = strings.Split(fileName, "/")
	path := s[0]
	for i := 1; i < len(s)-1; i++ {
		path = path + "/" + s[i]
	}
	s = strings.Split(s[len(s)-1], ".")
	name := s[0]

	// read configuration file
	viper.SetConfigName(name)
	viper.SetConfigType("toml")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		var d = DBConfig{}
		return d, err
	}

	// db configuration
	Driver := viper.GetString("DRIVER")
	Host := viper.GetString("HOST")
	DBPort := viper.GetInt("PORT")
	DBName := viper.GetString("DBNAME")
	User := viper.GetString("USER")
	Password := viper.GetString("PASSWORD")
	SSLMode := viper.GetString("SSLMODE")
	Timeout := viper.GetInt("Timeout")

	// Assign the gathered values to the configStruct struct of type sqlDBTx
	configStruct := DBConfig{
		Driver:   Driver,
		Host:     Host,
		Port:     DBPort,
		User:     User,
		Password: Password,
		DBName:   DBName,
		Sslmode:  SSLMode,
		Timeout:  Timeout,
	}
	return configStruct, nil

}
