package storage

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/charmbracelet/log"
)

type JsonStorage struct {
	filename string
	mutex    sync.Mutex
	data     map[string]string
}

func NewJsonStorage(filename string) (*JsonStorage, error) {

	jsonStorage := &JsonStorage{
		filename: filename,
		mutex:    sync.Mutex{},
		data:     make(map[string]string),
	}

	if err := jsonStorage.readFromFile(); err != nil {
		return jsonStorage, err
	}

	return jsonStorage, nil
}

func (js *JsonStorage) readFromFile() error {
	js.mutex.Lock()
	content, err := os.ReadFile(js.filename)
	js.mutex.Unlock()

	if err != nil {
		js.writeToFile()
		js.data = make(map[string]string)
		return nil
	}

	if err := json.Unmarshal(content, &js.data); err != nil {
		return err
	}

	return nil
}

func (js *JsonStorage) writeToFile() {
	js.mutex.Lock()
	defer js.mutex.Unlock()

	file, err := os.Create(js.filename)

	if err != nil {
		log.Fatal("JsonStorage.writeToFile: file creation", "err", err)
	}

	defer file.Close()
	jsonData, err := json.MarshalIndent(js.data, "", "  ")

	if err != nil {
		log.Fatal("JsonStorage.writeToFile: json encoding", "err", err)
	}

	file.Write(jsonData)
}

func (js *JsonStorage) Get(key string) (string, error) {
	value, ok := js.data[key]

	if !ok {
		return value, errors.New("not found")
	}

	return value, nil
}

func (js *JsonStorage) Set(key string, value string) {
	js.data[key] = value
	js.writeToFile()
}

func (js *JsonStorage) Delete(key string) error {
	_, ok := js.data[key]

	if !ok {
		return errors.New("not found")
	}

	delete(js.data, key)
	js.writeToFile()
	return nil
}
