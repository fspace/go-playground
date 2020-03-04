package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	"github.com/spf13/viper"
	"log"
	"os"

	// Import godotenv
	"github.com/joho/godotenv"
)

func main() {
	os.Exit(realMain())
}
func realMain() (exitCode int) {
	app := cli.App("fuction-demo", "demo for function type")
	// --------------------------------------------------------------------------------------- />
	//			## cmd 配置  如果命令行分散在各个包或者库中 此处会是集成点
	// Declare command, which is invocable with "uman info"
	app.Command("bs", "basic syntax", cli.ActionCommand(basics))
	app.Command("de", "use GoDotEnv lib to load the .env file ", cli.ActionCommand(GoDotEnv))
	app.Command("vip", "use Viper lib to load the .env file ", cli.ActionCommand(Viper))

	// ---------------------------------------------------------------------------------------------
	// With the app configured, execute it, passing in the os.Args array
	app.Run(os.Args)

	return
}

// ==============================================================================

// ==============================================================================

// use os package to get the env variable which is already set
func envVariable(key string) string {

	// set env variable using os package
	os.Setenv(key, "gopher")

	// return the env variable using os package
	return os.Getenv(key)
}
func basics() {
	// os package
	value := envVariable("name")

	fmt.Printf("os package: %s = %s \n", "name", value)
}

func GoDotEnv() {
	// godotenv package
	dotenv := goDotEnvVariable("STRONGEST_AVENGER")

	fmt.Printf("godotenv : %s = %s \n", "STRONGEST_AVENGER", dotenv)
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

/**

 */

// use viper package to read .env file
// return the value of the key
func viperEnvVariable(key string) string {

	// SetConfigFile explicitly defines the path, name and extension of the config file.
	// Viper will use this and not check any of the config paths.
	// .env - It will search for the .env file in the current directory
	viper.SetConfigFile(".env")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	value, ok := viper.Get(key).(string)

	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}

// use viper package to load/read the config file or .env file and
// return the value of the key
func viperConfigVariable(key string) string {

	// name of config file (without extension)
	viper.SetConfigName("config")
	// look for config in the working directory
	viper.AddConfigPath(".")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	value, ok := viper.Get(key).(string)

	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}

func Viper() {
	// viper package read .env
	viperenv := viperEnvVariable("STRONGEST_AVENGER")

	fmt.Printf("viper : %s = %s \n", "STRONGEST_AVENGER", viperenv)
}
