package render

import (
	"os"

	"golang.org/x/term"
)

const FallbackWidth = 80

type WidthDetector func() (int, error)

func DetectTerminalWidth() (int, error) {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	return width, err
}

func ResolveWidth(detector WidthDetector) int {
	if detector == nil {
		return FallbackWidth
	}

	width, err := detector()
	if err != nil || width <= 0 {
		return FallbackWidth
	}

	return width
}
