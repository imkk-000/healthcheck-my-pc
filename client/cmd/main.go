package main

import (
	"bytes"
	"encoding/json"
	"healthcheck/client/model"
	"healthcheck/client/util"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/process"
)

var config = &model.Config{}

func main() {
	rawConfig, err := ioutil.ReadFile("config/settings.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(rawConfig, config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("load config:", config)

	for {
		time.Sleep(time.Duration(config.Minute) * time.Minute)

		procs, err := process.Processes()
		if err != nil {
			log.Fatal(err)
		}

		procs, err = util.GetProcessByName(procs, config.Target)
		if err != nil {
			log.Fatal(err)
		}

		data, err := json.Marshal(util.GetInfo(procs))
		if err != nil {
			log.Fatal(err)
		}

		resp, err := http.Post(config.URL, "application/json", bytes.NewBuffer(data))
		if err != nil {
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		log.Println(resp.StatusCode, body)
	}
}
