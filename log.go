package logger2

import (
	"fmt"
)

var Version string = "0.1.3"

func Log(mess string) {
	fmt.Println("[LOG] 2 " + mess)
}
