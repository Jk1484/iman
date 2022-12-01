package configer

import (
	"encoding/json"
	"os"
)

func LoadConfig(path string, destination interface{}) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(&destination)
}
