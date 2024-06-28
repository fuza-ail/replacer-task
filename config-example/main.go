package main

import (
	"flag"
	"fmt"
)

func main () {
	key := flag.String("key", "", "key argument")
	value := flag.String("value", "", "value argument")
	fmt.Println(key, value)
}