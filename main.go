package main

import (
  "encoding/json"
  "flag"
  "fmt"
  "github.com/cnnrznn/channel"
  "io/ioutil"
  "time"
)

type Config struct {
    Peers []string `json:"peers"`
}

func checkError(err error) {
    if nil != err {
        fmt.Println(err)
    }
}

func main() {
    var err error

    id := flag.Int("id", -1, "Index of this process in the config")
    confFn := flag.String("conf", "config.json", "Config file for the network")
    flag.Parse()

    // Load the JSON config here
    config := Config{}

    confData, err := ioutil.ReadFile(*confFn)
    checkError(err)

    err = json.Unmarshal([]byte(confData), &config)
    checkError(err)

    // Instantiate the channel
    ch := channel.Channel{Id: *id,
                          Peers: config.Peers}
    fmt.Println(ch)

    // TODO run consensus protocol
}
