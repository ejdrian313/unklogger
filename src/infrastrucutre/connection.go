package infrastrucutre

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

func NewConnection(connString string) (*gorm.DB, error) {
	return gorm.Open("postgres", connString)
}

type Config struct {
	Server struct {
		Port   string `yaml:"port"`
		Host   string `yaml:"host"`
		Secret string `yaml:"secret"`
	} `yaml:"server"`
	Database struct {
		Name     string `yaml:"name"`
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"database"`
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func ReadFile(cfg *Config) {
	f, err := os.Open("config.yaml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func ReadEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}
