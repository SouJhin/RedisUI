package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"changeme/internal/define"
)

func ConnectionList() ([]*define.Connection, error) {
	nowPath, _ := os.Getwd()
	data, err := ioutil.ReadFile(nowPath + string(os.PathSeparator) + define.Configname)
	if errors.Is(err, os.ErrNotExist) {
		return []*define.Connection{{
			Identity: "1",
			Name:     "1",
			Addr:     "1",
			Port:     "2121",
			Username: "111",
			Password: "1290843879",
		}}, nil
	}
	conf := new(define.Config)
	err = json.Unmarshal(data, conf)
	if err != nil {
		fmt.Printf("err =====> 🚀🚀🚀 %v\n", err)
		return nil, err
	}
	return conf.Connections, nil
}
