package gologio

import (
	"bufio"
	"errors"
	"gopkg.in/yaml.v3"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func LoadConfig(configPath string, configStruct *Config) {
	parseConfig(readConfig(configPath), configStruct)
}

func Run(config Config, name string) (data <-chan string, err error) {
	switch {
	case "apache" == name:
		return readFile(config.Apache.Path)
	case "nginx" == name:
		return readFile(config.Nginx.Path)
	case "php_fpm" == name:
		return readFile(config.PhpFpm.Path)
	case "php" == name:
		return readFile(config.Php.Path)
	case "mysql" == name:
		return readFile(config.Mysql.Path)
	}
	return nil, errors.New("Unknown name.")
}

func readFile(path string) (data <-chan string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	f := bufio.NewReader(file)
	c := make(chan string)

	go func() {
		for {
			readLine, errRead := f.ReadString('\n')
			if errRead == io.EOF {
				break
			}
			c <- readLine
		}
		close(c)
		file.Close()
	}()

	return c, nil
}

func readConfig(configPath string) []byte {
	data, err := ioutil.ReadFile(configPath)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return data
}

func parseConfig(configData []byte, configStruct *Config) {
	err := yaml.Unmarshal(configData, &configStruct)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
