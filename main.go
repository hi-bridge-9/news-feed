package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/hi-bridge-9/news-feed/target"
)

func main() {
	path := flag.String("path", ".example_data/feed_list.json", "for read feed target list")
	bytes, err := ioutil.ReadFile(*path)
	if err != nil {
		panic(err)
	}

	var li []target.Info
	if err := json.Unmarshal(bytes, &li); err != nil {
		panic(err)
	}

	fmt.Println(li)
}
