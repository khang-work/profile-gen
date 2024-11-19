package txt

import "testing"

func TestWordWrap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		text       string
		lineLength int
		expected   string
	}{
		{
			name:       "Basic wrap within limit",
			text:       "Lorem ipsum dolor sit amet",
			lineLength: 10,
			expected:   "Lorem\nipsum\ndolor sit\namet",
		},
		{
			name:       "Longer line length",
			text:       "Lorem ipsum dolor sit amet",
			lineLength: 20,
			expected:   "Lorem ipsum dolor\nsit amet",
		},
		{
			name:       "Line length shorter than longest word",
			text:       "Supercalifragilisticexpialidocious is a long word",
			lineLength: 5,
			expected:   "Supercalifragilisticexpialidocious\nis a\nlong\nword",
		},
		{
			name:       "Empty text",
			text:       "",
			lineLength: 10,
			expected:   "",
		},
		{
			name:       "Zero line length",
			text:       "Lorem ipsum",
			lineLength: 0,
			expected:   "Lorem ipsum",
		},
		{
			name:       "Negative line length",
			text:       "Lorem ipsum",
			lineLength: -10,
			expected:   "Lorem ipsum",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Create local variable to be compatible with older Go versions.
			tt := tt

			actual := wordWrap(tt.text, tt.lineLength)
			if actual != tt.expected {
				t.Errorf("got %q, expected %q", actual, tt.expected)
			}
		})
	}
}
