package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func GetJson(file string, buf *Users) (*Users, error) {
	jsonFile, err := os.Open(file)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(byteValue, buf)
	return buf, nil
}

func SaveJson(file string, buf *Users) error {
	result, _ := json.Marshal(buf)
	err := ioutil.WriteFile(file, result, 0644)
	if err != nil {
		return err
	}
	return nil
}

