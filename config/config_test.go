package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDBConfig(t *testing.T) {
	fileName, portNumber := ParseArgs()
	assert.NotNil(t, fileName)
	assert.NotNil(t, portNumber)
	assert.FileExists(t, fileName)
	assert.IsType(t, 3, portNumber)
	config, err := GetDbConfig(fileName)
	assert.NotEmpty(t, config.Driver, "Driver")
	assert.NotEmpty(t, config.Port, "Port")
	assert.NotEmpty(t, config.Host, "Host")
	assert.NotEmpty(t, config.User, "User")
	assert.NotEmpty(t, config.Password, "Password")
	assert.NotEmpty(t, config.DBName, "DBName")
	assert.NotEmpty(t, config.Sslmode, "Sslmode")
	assert.NotEmpty(t, config.Timeout, "Timeout")
	assert.NoError(t, err)
	_, err = GetDbConfig("./somefile.txt")
	assert.Error(t, err)
	_, err = GetDbConfig("./test_bad_format.toml")
	assert.Error(t, err)
}
