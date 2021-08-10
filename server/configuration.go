package server

import (
	"delivery/models"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func GetApplication() models.Global {

	var app models.Global

	filename, err := filepath.Abs("/home/lucas/go/src/deliveryApi/docker-compose.yml")
	if err != nil {
		fmt.Println(err)
	}
	yamlfile, _ := ioutil.ReadFile(filename)

	err = yaml.Unmarshal(yamlfile, &app)
	fmt.Println(err)

	return app
}

func GetPort() string {
	NewServer()
	x := GetApplication()
	conv := strings.Split(x.Services.Db.Ports[0], ":")
	return "host=localhost port=" + conv[0] + " user=admin dbname=product sslmode=disable password=123456"
}
