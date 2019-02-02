package fin

import t "github.com/nsf/termbox-go"

const backOption = '\000'
const backspace = '\001'

func GetChar() rune {
	e := t.PollEvent()
	if e.Ch == 'q' || e.Ch == 'Q' || e.Key == t.KeyEsc {
		return backOption
	}
	if e.Key == t.KeyEnter {
		return '\n'
	}
	if e.Key == t.KeyBackspace || e.Key == t.KeyBackspace2 {
		return backspace
	}
	return e.Ch
}
