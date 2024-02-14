package dummy

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func SayHellWorld() {
	fmt.Fprint(out, "hello world")
}

func StringToWrite() {
	fmt.Fprint(out, stringToWrite)
}
