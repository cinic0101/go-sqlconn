package sqlconn

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type DBConfig struct {
	Databases map[string]Database `yaml:"databases"`
}

type Database struct {
	Driver string
	Host string
	Database string
	User string
	Password string
}

func UnmarshalConfig(filename string) (*DBConfig, error) {
	c := &DBConfig{}

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(buf, &c); err != nil {
		return nil, err
	}

	return c, err
}