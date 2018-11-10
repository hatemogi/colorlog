package colorlog

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	re "regexp"
)

var (
	숫자매처, _   = re.Compile("^[0-9]+$")
	실수매처, _   = re.Compile("^[0-9]*\\.[0-9]+")
	분수매처, _   = re.Compile("^[0-9]+/[0-9]+$")
	시간매처, _   = re.Compile("^[0-9]{2}:[0-9]{2}(:[0-9]{2})?")
	날짜매처, _   = re.Compile("^[0-9]{4}-[0-9]{2}-[0-9]{2}")
	단위매처, _   = re.Compile("^[0-9]+(µs|ms|%)$")
	IP매처, _   = re.Compile("^[0-9]+\\.[0-9]+\\.[0-9]+\\.[0-9]")
	헥사매처, _   = re.Compile("^(?i:[0-9a-f]{4,})$")
	UUID매처, _ = re.Compile("^(?i:[0-9a-f]{8}-[0-9a-f]{4})")
	도메인매처, _  = re.Compile("^(?i:([a-z]+(\\.[a-z])+)")
)

var (
	도입 = []byte{0x1b, 91, 51, 56, 58, 53, 58}
	리셋 = []byte{0x1b, 91, 48, 109}
)

const (
	기본색   byte = 0
	정수    byte = 1
	실수    byte = 2
	분수    byte = 3
	날짜    byte = 4
	시간    byte = 5
	IP    byte = 6
	도메인   byte = 7
	헥사    byte = 8
	UUID  byte = 9
	쌍따옴   byte = 10
	홑따옴   byte = 11
	중괄호   byte = 12
	대괄호   byte = 13
	괄호    byte = 14
	DEBUG byte = 15
	INFO  byte = 16
	WARN  byte = 17
	ERROR byte = 18
)

var 색상테이블 []byte

func 색칠(w []byte, color byte) []byte {
	return append(append(append(도입, fmt.Sprintf("%dm", 색상테이블[color])...), w...), 리셋...)
}

type 표시함수 func(타입 string, 단어 []byte) []byte

var 표시 표시함수 = 기본표시

func 단어꾸밈(w []byte) []byte {
	글자 := w[0]
	switch {
	case 글자 == ' ':
		return 색칠(표시("공백", w), 기본색)
	case (글자 >= '0' && 글자 <= '9'), (글자 >= 'a' && 글자 <= 'e'), (글자 >= 'A' && 글자 <= 'E'):
		switch {
		case 숫자매처.Match(w):
			return 색칠(표시("숫자", w), 정수)
		case IP매처.Match(w):
			return 색칠(표시("IP", w), IP)
		case 실수매처.Match(w):
			return 색칠(표시("실수", w), 실수)
		case 분수매처.Match(w):
			return 색칠(표시("분수", w), 분수)
		case 날짜매처.Match(w):
			return 색칠(표시("날짜", w), 날짜)
		case 시간매처.Match(w), 단위매처.Match(w):
			return 색칠(표시("시간", w), 시간)
		case 헥사매처.Match(w):
			return 색칠(표시("헥사", w), 헥사)
		case UUID매처.Match(w):
			return 색칠(표시("UUID", w), UUID)
		case bytes.Equal(w, []byte("DEBUG")):
			return 색칠(표시("레벨", w), DEBUG)
		case bytes.Equal(w, []byte("ERROR")):
			return 색칠(표시("레벨", w), ERROR)
		}
	case 글자 == '"':
		return 색칠(표시("인용", w), 쌍따옴)
	case 글자 == '\'':
		return 색칠(표시("인용", w), 홑따옴)
	case 글자 == '{':
		return 색칠(표시("{}", w), 중괄호)
	case 글자 == '[':
		return 색칠(표시("[]", w), 대괄호)
	case 글자 == '(':
		return 색칠(표시("괄호", w), 괄호)
	case 글자 == 'I', 글자 == 'W':
		switch {
		case bytes.Equal(w, []byte("INFO")):
			return 색칠(표시("레벨", w), INFO)
		case bytes.Equal(w, []byte("WARN")):
			return 색칠(표시("레벨", w), WARN)
		}
	}
	return 표시("그외", w)
}

var 닫힘문자 = map[byte]byte{'[': ']', '{': '}', '(': ')'}

func 줄처리(줄 []byte) []byte {
	const (
		닫힘 byte = 0
		일반 byte = 1
	)
	pos := 0
	출력 := []byte{}

	for pos < len(줄) {
		대상, idx := 줄[pos:], 0
		열림, 깊이 := 닫힘, 0
		var 이스케이프 byte
	단어찾기:
		for idx <= len(대상) {
			글자 := 대상[idx]
			idx++
			switch 열림 {
			case 닫힘:
				if idx >= len(대상) {
					break 단어찾기
				}
				switch 글자 {
				case ' ', ',', '=', ';':
					break 단어찾기
				case '[', '{', '(', '\'', '"':
					열림 = 글자
					깊이++
				default:
					열림 = 일반
				}
			case '[', '{', '(':
				switch 글자 {
				case 열림:
					깊이++
				case 닫힘문자[열림]:
					깊이--
				}
				if 깊이 == 0 || idx >= len(대상) {
					break 단어찾기
				}
			case '\'':
				if 글자 == '\'' || idx >= len(대상) {
					break 단어찾기
				}
			case '"':
				switch {
				case (이스케이프 == '\\' && 글자 == '"'):
					이스케이프 = 글자
				case (이스케이프 == '\\' && 글자 == '\\'):
					이스케이프 = ' '
				case 글자 == '"', idx >= len(대상):
					break 단어찾기
				default:
					이스케이프 = 글자
				}
			case 일반:
				if idx >= len(대상) {
					break 단어찾기
				}
				switch 글자 {
				case ' ', ',', '=', ';':
					idx--
					break 단어찾기
				}
			}
		}
		출력 = append(출력, 단어꾸밈(대상[:idx])...)
		pos += idx
	}
	return 출력
}

func 기본표시(타입 string, 토큰 []byte) []byte {
	return 토큰
}

func 디버그표시(타입 string, 토큰 []byte) []byte {
	return append(append([]byte(fmt.Sprintf("(%s ", 타입)), 토큰...), ')')
}

func init() {
	SetColors("000102030405060708090A0B0C0E0F1011121314")
}

// SetColors 종류별 출력할 색상 테이블 설정
func SetColors(table string) {
	dec, e := hex.DecodeString(table)
	if e == nil && len(dec) == 20 {
		색상테이블 = dec
	}
}

// SetDebug 디버그 출력설정
func SetDebug(debug bool) {
	if debug {
		표시 = 디버그표시
	} else {
		표시 = 기본표시
	}
}

// Run in에서 읽어서 out에 색칠 처리.
func Run(in io.Reader, out *os.File) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		out.Write(줄처리(scanner.Bytes()))
		out.WriteString("\r\n")
	}
}
