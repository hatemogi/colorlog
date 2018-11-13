package colorlog

/*
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
*/

/*

solarized

base03    #002b36  8/4 brblack  234 	23
base02    #073642  0/4 black    235 	23
base01    #586e75 10/7 brgreen  240    102
base00    #657b83 11/7 bryellow 241    103
base0     #839496 12/6 brblue   244    145
base1     #93a1a1 14/4 brcyan   245    145
base2     #eee8d5  7/7 white    254    230
base3     #fdf6e3 15/7 brwhite  230    231
yellow    #b58900  3/3 yellow   136    178
orange    #cb4b16  9/3 brred    166    166
red       #dc322f  1/1 red      160    203
magenta   #d33682  5/5 magenta  125    169
violet    #6c71c4 13/5 brmagenta 61    104
blue      #268bd2  4/4 blue      33	    38
cyan      #2aa198  6/6 cyan      37     37
green     #859900  2/2 green     64    142

16-231:  6 × 6 × 6 cube (216 colors): 16 + 36 × r + 6 × g + b (0 ≤ r, g, b ≤ 5)

*/

// SetTheme 색상 테마 설정

func SetTheme(theme string) {
	switch theme {
	case "atom-one-dark":
	case "solarized":
		// 기본색 정수 실수 분수 날짜 시간 IP 도메인 헥사 UUID  쌍따옴 홑따옴 중괄호 대괄호 괄호 DEBUG INFO WARN ERROR
		SetColorsBytes([]byte{145, 142, 37, 104, 38, 169, 166, 203, 178, 169, 37, 166, 38, 142, 37, 145, 146, 178, 203})
	case "seoul256-dark":
	default:
	}
}
