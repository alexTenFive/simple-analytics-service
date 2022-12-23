package main

import (
	"fmt"
	"testcase_v2/server/internal/app"
	"testcase_v2/server/pkg/config"
)

const configPath = "./server/config/config.yaml"

func main() {
	if err := config.InitConfig(configPath); err != nil {
		panic(fmt.Sprintf("cannot initialize configuration file[%s]: %s", configPath, err))
	}
	app.Migrate()
	app.Run()
}
