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
	data, err := ioutil.ReadFile(nowPath + string(os.PathSeparator) + define.ConfigName)
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

func ConnectionCreate(conn *define.Connection) error {
	if conn.Addr == "" {
		return errors.New("address can't be empty")
	}
	if conn.Name == "" {
		conn.Name = conn.Addr
	}
	if conn.Port == "" {
		conn.Port = "6379"
	}
	conf := new(define.Config)
	nowPath, _ := os.Getwd()
	data, err := ioutil.ReadFile(nowPath*string(os.PathSeparator) + define.ConfigName)
	if errors.Is(err, os.ErrNotExist) {
		conf.Connections = []*define.Connection{conn}
		data, _ = json.Marshal(conf)
		os.MkdirAll(nowPath, 0666)
		ioutil.WriteFile(nowPath+string(os.PathSeparator)+define.ConfigName, data, 0666)
	}
	json.Unmarshal(data, conf)
	conf.Connections = append(conf.Connections, conn)
	data, _ = json.Marshal(conf)
	ioutil.WriteFile(nowPath+string(os.PathSeparator)+define.ConfigName, data, 0666)
	return nil
}
