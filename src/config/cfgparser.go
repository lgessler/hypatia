package cfgparser

import (
  "github.com/go-yaml/yaml"
  "fmt"
  "os"
  "io/ioutil"
)

type ShowAttrs struct {
  StartTime string `yaml:"StartTime"`
  Duration string `yaml:"Duration"`
}

type DaysOfTheWeek struct {
  Monday []ShowAttrs `yaml:"Monday"`
  Tuesday []ShowAttrs `yaml:"Tuesday"`
  Wednesday []ShowAttrs `yaml:"Wednesday"`
  Thursday []ShowAttrs `yaml:"Thursday"`
  Friday []ShowAttrs `yaml:"Friday"`
  Saturday []ShowAttrs `yaml:"Saturday"`
  Sunday []ShowAttrs `yaml:"Sunday"`
}

type Config struct {
  StreamURL string `yaml:"StreamURL"`
  BaseSavePath string `yaml:"BaseSavePath"`
  Shows DaysOfTheWeek `yaml:"Shows"`
}

func GetConfig() Config {
  file, e := ioutil.ReadFile("./config.yaml")
  if e != nil {
    fmt.Printf("Error reading config.json--is it there?\n\n%v", e)
    os.Exit(1)
  }

  var config Config
  err := yaml.Unmarshal(file, &config)
  if err != nil {
    panic(err)
  }

  return config
}
