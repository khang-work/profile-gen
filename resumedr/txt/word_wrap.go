package txt

import (
	"strings"
)

// wordWrap breaks long lines of the string input by word.
func wordWrap(text string, lineLength int) string {
	if lineLength <= 0 {
		return text // Return as is for invalid line length.
	}

	var result strings.Builder

	lines := strings.Split(text, "\n") // Splits the string by newline.
	newlineAtEnd := false

	if len(text) > 0 {
		lastChar := text[len(text)-1]
		if lastChar == '\n' {
			newlineAtEnd = true
		}
	}

	for i, line := range lines {
		words := strings.Fields(line) // Split the text into words.
		curLineLength := 0

		for j, word := range words {
			wordLen := len(word)

			// If adding the word exceeds the line length, wrap to the next line.
			if curLineLength+wordLen+1 > lineLength {
				if j > 0 {
					result.WriteString("\n")
					curLineLength = 0
				}
			} else if j > 0 {
				// Add a space before the word if it's not the first word on the line.
				result.WriteString(" ")
				curLineLength++
			}

			// Add actual word to the buffer.
			result.WriteString(word)
			curLineLength += wordLen
		}

		if i == len(lines)-1 {
			if newlineAtEnd {
				result.WriteString("\n")
			}
		} else {
			result.WriteString("\n")
		}
	}

	return result.String()
}
