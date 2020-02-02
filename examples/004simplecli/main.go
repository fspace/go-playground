package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	//	"os"
	"path/filepath"
	//"runtime"
)

var (
	enableLog = flag.Bool("log", false, "enable log")
)

var (
	ErrOpenLogFile = errors.New("cannot open log file")
)

// AppConfig 应用相关的配置
type AppConfig struct {
}

func DefaultConfig() AppConfig {
	return AppConfig{}
}

func printError(err error) {
	fmt.Fprintln(os.Stderr, err)
}

func initConfig() AppConfig {
	var config AppConfig

	configDir, err := os.UserHomeDir() // UserConfigDir()
	//var err error
	if err != nil {
		printError(err)
		config = DefaultConfig()
	} else {
		configDir = filepath.Join(configDir, "ff")

		configFile := filepath.Join(configDir, "config.yaml")
		if system.IsExist(configFile) {
			b, err := ioutil.ReadFile(configFile)
			if err != nil {
				printError(err)
				config = DefaultConfig()
			} else {
				if err := yaml.Unmarshal(b, &config); err != nil {
					printError(err)
					config = DefaultConfig()
				}
			}
		} else {
			config = DefaultConfig()
		}
		config.ConfigDir = configDir
		config.ConfigFile = configFile
	}

	if *enableLog && configDir != "" {
		config.Log.Enable = *enableLog
		if config.Log.File == "" {
			config.Log.File = filepath.Join(configDir, "ff.log")
		}
	}

	//if *enableTree {
	//	config.EnableTree = *enableTree
	//}

	//if config.OpenCmd == "" {
	//	switch runtime.GOOS {
	//	case "darwin":
	//		system.OpenCmd = "open"
	//	case "linux":
	//		system.OpenCmd = "xdg-open"
	//	}
	//} else {
	//	system.OpenCmd = config.OpenCmd
	//}
	return config
}

func initLogger(config gui.Config) error {
	var logWriter io.Writer
	var err error
	if config.Log.Enable {
		logWriter, err = os.OpenFile(os.ExpandEnv(config.Log.File),
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

		if err != nil {
			return fmt.Errorf("%s: %s", ErrOpenLogFile, err)
		}
		log.SetFlags(log.Lshortfile)
	} else {
		// don't print log
		logWriter = ioutil.Discard
	}

	log.SetOutput(logWriter)
	return nil
}

func run() int {
	flag.Parse()

	config := initConfig()
	if err := initLogger(config); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	if err := gui.New(config).Run(); err != nil {
		return 1
	}

	return 0
}

func main() {
	os.Exit(run())
}

/**
@see https://github.com/skanehira/ff/blob/master/main.go
*/
