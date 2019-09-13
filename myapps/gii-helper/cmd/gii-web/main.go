package main

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"playgo/myapps/gii-helper/cmd/gii-web/app"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

func run() error {

	conf, err := app.LoadConfig("./config/app.yaml")
	if err != nil {
		return errors.Wrap(err, "LoadConfig")
	}
	_ = conf
	fmt.Println(conf)

	//err := setupXxx()
	//if err != nil {
	//	return errors.Wrap(err,"setup Xxx")
	//}
	// ...
	svr := http.Server{}
	_ = svr

	return nil
}

func setupXxx() error {
	return nil
}
