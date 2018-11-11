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

// SetTheme 색상 테마 설정
func SetTheme(theme string) {
	switch theme {
	case "atom-one-dark":
	case "solarized":
		SetColorsBytes([]byte{241, 136, 136, 136, 166, 166, 125, 244, 240, 61, 33, 37, 61, 37, 241, 244, 245, 254, 230, 230})
	case "seoul256-dark":
	default:
	}
}
