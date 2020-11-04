package example

import (
	"bytes"
	"fmt"
	"github.com/dNaszta/go_cookbook/Ch01-io-interfaces/L01-interfaces/interfaces"
)

func main()  {
	in := bytes.NewReader([]byte("example"))
	out := &bytes.Buffer{}
	fmt.Print("stdout on Copy = ")
	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}

	fmt.Println("out bytes buffer = ", out.String())

	fmt.Print("stdout on PipeExample = ")
	if err := interfaces.PipeExample(); err != nil {
		panic(err)
	}
}
