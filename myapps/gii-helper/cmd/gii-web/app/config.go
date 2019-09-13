package app

import (
	// 也可以试试 go-ini  用ini做配置文件  简单程序不用考虑太多
	"github.com/jinzhu/configor"
)

type Config struct {
	DbUser string // = "root"
	DbPass string // = ""
	DbHost string //= "127.0.0.1"
	DbPort int    //= 3306

	APPName string `default:"app name"`
	// 原始配置对象 OriginalConfigor
	// Raw *config.Config
	configor    *configor.Configor // TODO 这里保存原始配置对象是为了以后用 可以考虑改为提供Load|Populate|Configure 方法
	configPaths []string
}

// Configure populate the cfg object from the original configuration
// 借此 是否可以实现 分次提取配置？
func (c *Config) Configure(cfg interface{}) error {
	return configor.Load(cfg, c.configPaths...)
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

	*conf = DefaultConfig

	c := configor.New(nil)
	err := c.Load(conf, configPaths...)
	if err != nil {
		return nil, err
	}
	// 保存起来重复可以重复使用哦
	conf.configor = c
	conf.configPaths = configPaths

	//	configor.Load(conf, "./config/app.yaml")
	// fmt.Printf("config: %#v", conf)
	return conf, nil
}
