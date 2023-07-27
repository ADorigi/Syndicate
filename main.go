package main

import (
	"fmt"

	"github.com/adorigi/syndicate/systemInfo"
)

func main() {
	info := systemInfo.CollectStats()
	fmt.Println(info)
}
