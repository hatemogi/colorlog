package colorlog

/*

	기본색  byte = 0
	정수    byte = 1
	실수    byte = 2
	분수    byte = 3
	날짜    byte = 4
	시간    byte = 5
	IP      byte = 6
	헥사    byte = 7
	UUID    byte = 8
	쌍따옴  byte = 9
	홑따옴  byte = 10
	중괄호  byte = 11
	대괄호  byte = 12
	괄호    yte = 13
	DEBUG   byte = 14
	INFO    byte = 16
	ERROR   byte = 17

*/

// SetTheme 색상 테마 설정
func SetTheme(theme string) {
	switch theme {
	case "atom-one-dark":
	case "solarized-dark":
	default:
	}
}
