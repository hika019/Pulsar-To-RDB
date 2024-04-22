package main

import "fmt"

func main() {
	conf, err := readConfig()

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(conf)
}
