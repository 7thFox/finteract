package fin

import (
	"fmt"

	t "github.com/nsf/termbox-go"
)

func Print(x, y int, parts ...interface{}) int {
	msg := fmt.Sprint(parts...)
	for _, c := range msg {
		if c == '\n' {
			y++
			x = 0
			continue
		}
		t.SetCell(x, y, c, t.ColorDefault, t.ColorDefault)
		x++
	}
	t.Flush()
	return y + 1
}

func Printf(x, y int, f string, parts ...interface{}) int {
	msg := fmt.Sprintf(f, parts...)
	return Print(x, y, msg)
}

func Clear() int {
	t.Clear(t.ColorDefault, t.ColorDefault)
	return 0
}
