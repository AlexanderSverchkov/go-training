package files

import (
	"encoding/json"
	"os"
)

func ReadAndLoadFileInJson(fileName string) (map[string]string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	parsedJson := make(map[string]string)
	err = json.Unmarshal(data, &parsedJson)
	if err != nil {
		return nil, err
	}
	return parsedJson, nil
}

func ReadFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		return err
	}
	return nil
}
