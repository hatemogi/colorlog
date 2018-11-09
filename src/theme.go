package colorlog

// SetTheme 색상 테마 설정
func SetTheme(theme string) {
	switch theme {
	case "atom-one-dark":
		SetColors([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	case "solarized-dark":
		SetColors([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	default:
		SetColors([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	}
}
