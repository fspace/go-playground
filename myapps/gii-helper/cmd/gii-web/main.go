package main

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"playgo/myapps/gii-helper/cmd/gii-web/app"

	"github.com/gorilla/mux"
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
	//_ = conf
	//fmt.Printf("config: %#v", conf)

	//err := setupXxx()
	//if err != nil {
	//	return errors.Wrap(err,"setup Xxx")
	//}
	// ...
	// TODO 学习使用Mux 打造自己的路由

	//mux := http.NewServeMux()
	mux := mux.NewRouter()

	svr := app.Server{
		AppConfig: conf,
		Router:    mux,
	}
	// 构建路由配置
	svr.Routes()

	http.ListenAndServe(":8000", mux)

	return nil
}

func setupXxx() error {
	return nil
}
