package main

import "github.com/dNaszta/go_cookbook/Ch01-io-interfaces/L02-bytestrings/bytestrings"

func main()  {
	err := bytestrings.WorkWithBuffer()
	if err != nil {
		panic(err)
	}

	// each of these print to stdout
	bytestrings.SearchString()
	bytestrings.ModifyString()
	bytestrings.StringReader()
}
