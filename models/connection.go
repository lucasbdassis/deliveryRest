package models

type Connection struct {
	Ports interface{} `yaml:"services"`
}

type Global struct {
	Services Db `yaml:"services"`
}

type Db struct {
	Db Ports `yaml:"db"`
}

type Ports struct {
	Ports []string `yaml:"ports"`
}
