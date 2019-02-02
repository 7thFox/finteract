package fin

import (
	"strconv"
	"strings"

	t "github.com/nsf/termbox-go"
)

func GetString() string {
	n := startLongInput()
	sb := &strings.Builder{}
	i := 0
	for c := GetChar(); c != '\n'; c = GetChar() {
		if c == backspace {
			i = inputBackspace(sb, n, i)
			continue
		}
		i = inputNormal(sb, c, n, i)
	}
	endLongInput()
	return sb.String()
}

func GetNum() int {
	n := startLongInput()
	sb := &strings.Builder{}
	i := 0
	for c := GetChar(); c != '\n'; c = GetChar() {
		if c == backspace {
			i = inputBackspace(sb, n, i)
			continue
		}
		if c < '0' || c > '9' {
			continue
		}
		i = inputNormal(sb, c, n, i)
	}
	endLongInput()
	id, _ := strconv.Atoi(sb.String())
	return id
}

func inputBackspace(sb *strings.Builder, n, i int) int {
	if i == 0 {
		return i
	}

	s := sb.String()[:sb.Len()-1]
	sb.Reset()
	sb.WriteString(s)
	i--
	t.SetCell(i, n, ' ', t.ColorWhite, t.ColorBlack)
	t.SetCursor(i, n)
	t.Flush()
	return i
}

func inputNormal(sb *strings.Builder, c rune, n, i int) int {
	sb.WriteRune(c)
	t.SetCell(i, n, c, t.ColorWhite, t.ColorBlack)
	t.SetCursor(i+1, n)

	t.Flush()
	return i + 1
}

func startLongInput() int {
	w, h := t.Size()
	for i := 0; i < w; i++ {
		t.SetCell(i, h-1, ' ', t.ColorWhite, t.ColorBlack)
	}
	t.SetCursor(0, h-1)
	t.Flush()
	return h - 1
}

func endLongInput() int {
	w, h := t.Size()
	for i := 0; i < w; i++ {
		t.SetCell(i, h-1, ' ', t.ColorDefault, t.ColorDefault)
	}
	t.HideCursor()
	t.Flush()
	return h - 1
}
