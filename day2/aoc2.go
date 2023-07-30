package day2

import (
	//"fmt"
	//"strconv"
	"fmt"
	"strings"

	"golang.design/x/clipboard"
)

func main() {

	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

}

func data() {
	const rock, paper, scissors, loss, draw, win int = 1, 2, 3, 0, 3, 6 // rock, paper, scissors, loss, draw, win r, p, s, l, d ,w

	content := clipboard.Read(clipboard.FmtText)
	//var keys []string
	if content == nil {
		panic("No clipboard content")
	}
	sets := string(content)
	setsMap := strings.Split(sets, "\r\n\r\n")
	for _, vals := range setsMap {
		fmt.Println(vals)
	}
}
