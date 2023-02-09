package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/ipld/go-car"
	"github.com/ipld/go-car/util"
)

var rootsFlag = flag.Bool("roots", false, "Display CAR roots")

func printfErr(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
}

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		printfErr("Usage: ./carls file.car")
	}

	f, err := os.Open(flag.Args()[0])
	if err != nil {
		printfErr(err.Error())
		os.Exit(1)
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	h, err := car.ReadHeader(buf)
	if err != nil {
		printfErr(err.Error())
		os.Exit(1)
	}

	if *rootsFlag {
		for _, root := range h.Roots {
			fmt.Println(root)
		}
		return
	}

	for {
		c, _, err := util.ReadNode(buf)
		if err == io.EOF {
			return
		}
		if err != nil {
			printfErr(err.Error())
			os.Exit(1)
		}
		fmt.Println(c)
	}
}
