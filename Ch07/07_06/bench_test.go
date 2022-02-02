package nlp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var benchText = "Don't communicate by sharing memory, share memory by communicating."

func BenchmarkTokenize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tokens := Tokenize(benchText)
		require.Equal(b, 10, len(tokens))
	}
}

func TestTokenize(t *testing.T) {
	text := "Who's on first?"
	expected := []string{"who", "s", "on", "first"}
	tokens := Tokenize(text)
	require.Equal(t, expected, tokens)
}
