package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hatemogi/colorlog/src"
)

func 종료시그널처리() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(0)
	}()
}

var (
	verbose bool
	version bool
	debug   bool
	colors  string
)

func init() {
	flag.BoolVar(&version, "v", false, "Show version and exit")
	flag.BoolVar(&debug, "d", false, "Turn on debug mode")
	flag.Parse()
	종료시그널처리()
}

func main() {
	if version {
		fmt.Printf("colorlog %v\n", "0.0.3")
		os.Exit(0)
	}
	if colors := os.Getenv("CL_COLORS"); len(colors) == 32 {
		table, e := hex.DecodeString(colors)
		if e == nil {
			colorlog.SetColors(table)
		}
	}
	colorlog.SetDebug(debug)
	colorlog.Run(os.Stdin, os.Stdout)
}
