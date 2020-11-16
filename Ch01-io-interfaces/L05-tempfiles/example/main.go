package main

import "github.com/dNaszta/go_cookbook/Ch01-io-interfaces/L05-tempfiles/tempfiles"

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
