package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/kardianos/osext"
	"gopkg.in/yaml.v2"
)

type Jobs map[string][]string

type TazFile struct {
	Image   string
	Preapre []string `yaml:"prepare"`
	Process Jobs     `yaml:"process"`
}

func main() {
	currentDir, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadFile(currentDir + string(filepath.Separator) + "taz.yml")
	t := TazFile{}
	err = yaml.Unmarshal([]byte(data), &t)
	if err == nil {
		fmt.Println(t)
	} else {
		fmt.Println("error", err)
	}
	fmt.Println("done")
}
