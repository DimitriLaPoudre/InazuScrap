package player

import (
	"encoding/json"
	"os"
)

func Load(path string) ([]Player, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	list := []Player{}
	if err := json.NewDecoder(file).Decode(&list); err != nil {
		return nil, err
	}

	return list, nil
}

func Save(path string, list []Player) error {
	file, err := os.OpenFile(path,
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644)
	if err != nil {
		return err
	}
	defer file.Close()
	defer file.Sync()

	if err := json.NewEncoder(file).Encode(list); err != nil {
		return err
	}

	return nil
}
