package cli

import (
	"errors"
	"flag"
	"log"
	"os"

	"github.com/SarahLightBourne/okay/http"
	"github.com/SarahLightBourne/okay/storage"

	"github.com/charmbracelet/huh"
)

func GetStorageFromArgs() (http.Storage, error) {
	var help bool
	var storageStr string
	var filename string

	flag.BoolVar(&help, "help", false, "Help")
	flag.StringVar(&storageStr, "storage", "", "Choose storage")
	flag.StringVar(&filename, "filename", "data.json", "Filename for persisted storage")

	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if storageStr == "" {
		return nil, errors.New("no storage specified")
	}

	if storageStr == "memory" {
		return storage.NewMemoryStorage(), nil
	}

	if storageStr == "json" {
		jsonStorage, err := storage.NewJsonStorage(filename)

		if err != nil {
			log.Fatal("invalid json")
		}

		return jsonStorage, nil
	}

	return nil, errors.New("unknown storage")
}

func GetStorageFromCli() http.Storage {
	var storageStr string

	huh.NewSelect[string]().
		Title("Select storage").
		Description("use -help for automation").
		Options(
			huh.NewOption("In-memory storage", "memory"),
			huh.NewOption("Json storage", "json"),
		).Value(&storageStr).Run()

	if storageStr == "memory" {
		return storage.NewMemoryStorage()
	}

	var filename string
	huh.NewInput().Title("Enter filename").Value(&filename).Run()

	jsonStorage, err := storage.NewJsonStorage(filename)

	if err != nil {
		log.Fatal("invalid json")
	}

	return jsonStorage
}
