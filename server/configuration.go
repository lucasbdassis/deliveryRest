package server

import (
	"delivery/models"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)



func GetApplication() (models.Global) {

	var app models.Global

	filename , err := filepath.Abs("/home/lucas/go/src/deliveryApi/docker-compose.yml")
	if err != nil {
		fmt.Println("nao consegui")
	}
	yamlfile , _ := ioutil.ReadFile(filename)



	err = yaml.Unmarshal(yamlfile, &app)
	fmt.Println(err)

	return app
}
