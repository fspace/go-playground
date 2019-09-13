package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
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
	fmt.Printf("config: %#v", conf)

	// 测试是否可以再次使用配置对象：
	var conf2 = struct {
		Contacts []struct {
			Name  string
			Email string `required:"true"`
		}
	}{}
	if err := conf.Configure(&conf2); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("config2: %#v", conf2)

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
