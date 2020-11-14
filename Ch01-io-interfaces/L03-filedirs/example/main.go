package main

import "github.com/dNaszta/go_cookbook/Ch01-io-interfaces/L03-filedirs/filedirs"

func main() {
	if err := filedirs.Operate(); err != nil {
		panic(err)
	}

	if err := filedirs.CapitalizerExample(); err != nil {
		panic(err)
	}
}