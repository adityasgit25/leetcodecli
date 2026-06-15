package render

import (
	"io"
	"os"

	"golang.org/x/term"
)

const FallbackWidth = 80

type WidthDetector func() (int, error)

func DetectTerminalWidth() (int, error) {
	return detectFileWidth(os.Stdout)
}

func DetectWriterWidth(writer io.Writer) WidthDetector {
	file, ok := writer.(*os.File)
	if !ok {
		return nil
	}

	return func() (int, error) {
		return detectFileWidth(file)
	}
}

func detectFileWidth(file *os.File) (int, error) {
	width, _, err := term.GetSize(int(file.Fd()))
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
