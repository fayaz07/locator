package console

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func boxHeader(width int) string {
	return fmt.Sprintf("%s%s%s", BOX_TOP_LEFT, strings.Repeat(BOX_HORIZONTAL, width-2), BOX_TOP_RIGHT)
}

func boxFooter(width int) string {
	return fmt.Sprintf("%s%s%s", BOX_BOTTOM_LEFT, strings.Repeat(BOX_HORIZONTAL, width-2), BOX_BOTTOM_RIGHT)
}

func boxEmptyLine(width int) string {
	return fmt.Sprintf("%s%s%s", BOX_VERTICAL, strings.Repeat(" ", width-2), BOX_VERTICAL)
}

func boxBarWithText(width int, text string) string {
	blankSpaces := strings.Repeat(" ", (width-len(text))/2)
	return fmt.Sprintf("%s%s%s%s%s", BOX_VERTICAL, blankSpaces, text, blankSpaces[1:], BOX_VERTICAL)
}

func InBox(width int, text string) {
	Println(boxHeader(width))

	emptyBoxLine := boxEmptyLine(width)
	Println(emptyBoxLine)

	Println(boxBarWithText(width, text))

	Println(emptyBoxLine)

	Println(boxFooter(width))
}

func Print(value string) {
	fmt.Print(color.BlackString(value))
}

func Println(value string) {
	fmt.Println(color.BlackString(value))
}

func Info(value string) {
	fmt.Println(color.CyanString(value))
}

func InfoTS(value string, tabSpaces int) {
	fmt.Println(
		fmt.Sprintf("%s%s", strings.Repeat("\t", tabSpaces), color.CyanString(value)),
		color.CyanString(value),
	)
}

func InfoS(value string, spaces int) {
	fmt.Println(
		fmt.Sprintf("%s%s", strings.Repeat(" ", spaces), color.CyanString(value)),
		color.CyanString(value),
	)
}

func Error(value string) {
	fmt.Println(color.RedString(value))
}
