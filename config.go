package main

import (
    "io/ioutil"
    "log"
    "encoding/json"
)

func check(e error) {
    if e != nil {
        log.Panic(e)
        panic(e)
    }
}


type BotConfig struct {
    Token string `json:"token"`
    Hostname string `json:"hostname"`
    Port uint16 `json:"port"`
}


func getConfig() (BotConfig) {
    var config BotConfig

    data, err := ioutil.ReadFile("config.json")
    check(err)
    
    json_err := json.Unmarshal(data, &config)
    check(json_err)

    return config
}