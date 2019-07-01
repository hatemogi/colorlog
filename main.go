package main

import (
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
		colorlog.Reset(os.Stdout)
		os.Exit(0)
	}()
}

var (
	verbose bool
	version bool
	debug   bool
	theme   string
)

func init() {
	flag.BoolVar(&version, "v", false, "Show version and exit")
	flag.BoolVar(&debug, "d", false, "Turn on debug mode")
	flag.StringVar(&theme, "t", "", "set color theme")
	flag.Parse()
	종료시그널처리()
}

func main() {
	if version {
		fmt.Printf("colorlog %v\n", "0.0.3")
		os.Exit(0)
	}

	if len(theme) > 0 {
		colorlog.SetTheme(theme)
	} else if colors := os.Getenv("COLORLOG_COLORS"); len(colors) > 0 {
		colorlog.SetColors(colors)
	}
	colorlog.SetTheme("solarized")

	colorlog.SetDebug(debug)
	colorlog.Run(os.Stdin, os.Stdout)
}
