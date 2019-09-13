package app

//import "github.com/go-ozzo/ozzo-config"
import (
	"github.com/jinzhu/configor"
)

type Config struct {
	DbUser string // = "root"
	DbPass string // = ""
	DbHost string //= "127.0.0.1"
	DbPort int    //= 3306

	APPName string `default:"app name"`
	// 原始配置对象
	// Raw *config.Config
	RawConfigor *configor.Configor
}

var DefaultConfig = Config{
	DbUser: "root",
	DbPass: "",
	DbHost: "127.0.0.1",
	DbPort: 3306,
}

func (conf Config) Validate() error {
	//return validation.ValidateStruct(&config,
	//	validation.Field(&config.DSN, validation.Required),
	//	validation.Field(&config.JWTSigningKey, validation.Required),
	//	validation.Field(&config.JWTVerificationKey, validation.Required),
	//)
	return nil
}

// LoadConfig loads configuration from the given list of paths and populates it into the Config variable.
// The configuration file(s) should be named as app.yaml.
func LoadConfig(configPaths ...string) (*Config, error) {
	conf := &Config{}
	// create a Config object
	//c := config.New()
	//
	//// load from one or multiple JSON, YAML, or TOML files.
	//// file formats are determined by their extensions: .json, .yaml, .yml, .toml
	//c.Load(configPaths...)
	//conf := &Config{}
	//
	//// *conf = DefaultConfig
	//
	//c.Configure(conf)
	//
	//conf.Raw = c
	// configor.Load(conf, configPaths...)
	c := configor.New(nil)
	err := c.Load(conf, configPaths...)
	if err != nil {
		return nil, err
	}
	conf.RawConfigor = c
	//	configor.Load(conf, "./config/app.yaml")
	// fmt.Printf("config: %#v", conf)
	return conf, nil
}
