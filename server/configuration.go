package server

import (
	"delivery/models"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func GetApplication() models.Global {

	var app models.Global


	filename, err := filepath.Abs("./docker-compose.yml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(filepath.Dir(";"))
	yamlfile, _ := ioutil.ReadFile(filename)

	err = yaml.Unmarshal(yamlfile, &app)


	return app
}

func GetPort() string {
	NewServer()

	x := GetApplication()
	conv := strings.Split(x.Services.Db.Ports[0], ":")
	return "host=localhost port=" + conv[0] + " user=admin dbname=product sslmode=disable password=123456"
}
