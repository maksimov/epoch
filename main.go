package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/araddon/dateparse"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage:")
		fmt.Println("	epoch <timestamp or date>")
		fmt.Println()
		fmt.Println("Parameters:")
		fmt.Println("	timestamp - A Unix timestamp in second or milliseconds")
		fmt.Println("	date - A date in most common formats, e.g. 2019-04-16 12:05:00")
		fmt.Println()
		fmt.Println("Examples:")
		fmt.Println("	epoch 1555416300")
		fmt.Println("	epoch 1555416300000")
		fmt.Println("	epoch 2019-04-16 12:05:00")
		fmt.Println("	epoch 04/16/2019 12:05:00")
		fmt.Println("	epoch 2019-04-16 15:05:00 +0300 UTC")
		fmt.Println()
		os.Exit(1)
	}

	input := strings.Join(os.Args[1:], " ")
	t, err := dateparse.ParseAny(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println("Milliseconds: ", t.Unix()*1000)
	fmt.Println("         GMT: ", t.UTC())
	fmt.Println("       Local: ", t)
}
