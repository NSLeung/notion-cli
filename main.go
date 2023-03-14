package main

import (
    "fmt"
    "io/ioutil"
    "net/http"

    "gopkg.in/yaml.v2"
)

type Config struct {
    Token      string `yaml:"token"`
}

func (c *Config) readConfig() *Config {
    cfgFile, err := ioutil.ReadFile("./config.yaml")
    if err != nil {
        panic(err)
    }
    err = yaml.Unmarshal(cfgFile, c)
    if err != nil {
	panic(err)
    }
    return c
}
func main() {
    var c Config
    c.readConfig()
    // fmt.Printf("%+v", c)
    url := "https://api.notion.com/v1/pages/1f2c887cd57b4f71b2e50318e9ecb2d2"

    req, _ := http.NewRequest("GET", url, nil)

    req.Header.Add("accept", "application/json")
    req.Header.Add("Notion-Version", "2022-06-28")
    var bearer = "Bearer " + c.Token
    req.Header.Add("Authorization", bearer)
    res, _ := http.DefaultClient.Do(req)

    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)

    fmt.Println(res)
    fmt.Println(string(body))
    /*
    req, err := http.NewRequest("GET", config.URL, nil)
    if err != nil {
        panic(err)
    }
    req.SetBasicAuth(config.Username, config.Token)
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(body))
    */
}

