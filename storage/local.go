package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type LocalStorage struct {
	StoragePath string
}

func (this LocalStorage) Store(cluster Cluster) error {
	data, err := json.Marshal(cluster)
	if err != nil {
		return err
	}
	return os.WriteFile(this.StoragePath, data, os.ModePerm)
}

func (this LocalStorage) Read() (Cluster, error) {
	res := Cluster{}
	data, err := ioutil.ReadFile(this.StoragePath)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}
