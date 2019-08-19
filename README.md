### Parser log files


#### Example
```go
package main

import (
	"fmt"
	"gologio/src"
	"log"
)

func main() {
	config := gologio.Config{}
	gologio.LoadConfig("config.yml", &config)

	data, err := gologio.Run(config, "php")

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for val := range data {
		fmt.Printf("%v", val)
	}
}
```
