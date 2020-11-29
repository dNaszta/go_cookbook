package main

import "github.com/dNaszta/go_cookbook/Ch01-io-interfaces/L06-templates/templates"

func main() {
	if err := templates.RunTemplate(); err != nil {
		panic(err)
	}

	if err := templates.InitTemplates(); err != nil {
		panic(err)
	}

	if err := templates.HTMLDifferences(); err != nil {
		panic(err)
	}
}
