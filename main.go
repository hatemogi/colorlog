package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
)

func 매치(re string, s string) bool {
	m, _ := regexp.MatchString(re, s)
	return m
}

func 색상(s string, color int) string {
	return fmt.Sprintf("\x1b[38:5:%dm%s\x1b[0m", color, s)
}

type 표시함수 func(타입 string, 단어 string) string

func 구분(w string, 표시 표시함수) string {
	if w == "" {
		return 표시("공백", " ")
	} else if 매치("^[0-9]+$", w) {
		return 색상(표시("숫자", w), 87)
	} else if 매치("^\".*\"$", w) || 매치("^'.*'$", w) {
		return 색상(표시("인용", w), 154)
	} else if 매치("^[0-9]+\\.[0-9]+\\.[0-9]+\\.[0-9]+(:[0-9]+)?$", w) {
		return 색상(표시("IP", w), 39)
	} else if 매치("^[0-9]{4}-[0-9]{2}-[0-9]{2}", w) {
		return 색상(표시("시간", w), 39)
	} else if strings.HasPrefix(w, "{") || strings.HasPrefix(w, "[") {
		return 색상(표시("JSON", w), 214)
	} else if w == "INFO" || w == "DEBUG" || w == "ERROR" || w == "WARN" {
		return 색상(표시("레벨", w), 76)
	}
	return 표시("그외", w)
}

func 줄처리(line string, 표시 표시함수) string {
	words := strings.Split(line, " ")
	var output []string
	for _, word := range words {
		output = append(output, 구분(word, 표시))
	}
	return strings.Join(output, " ")
}

func 기본표시(타입 string, 토큰 string) string {
	return 토큰
}

func 디버그표시(타입 string, 토큰 string) string {
	return fmt.Sprintf("(%s %s)", 타입, 토큰)
}

func main() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(0)
	}()

	scanner := bufio.NewScanner(os.Stdin)
	var 표시 표시함수 = 기본표시
	if len(os.Args) > 1 && os.Args[1] == "-d" {
		표시 = 디버그표시
	}
	for scanner.Scan() {
		fmt.Println(줄처리(scanner.Text(), 표시))
	}
}
