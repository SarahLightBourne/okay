package main

import (
	"github.com/SarahLightBourne/okay/cli"
	"github.com/SarahLightBourne/okay/http"
)

func main() {
	storage, err := cli.GetStorageFromArgs()

	if err != nil {
		storage = cli.GetStorageFromCli()
	}

	app := http.NewApp(storage)
	app.RunHttp(":8021")
}
