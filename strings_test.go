package slackbot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBrackets(t *testing.T) {
	t.Parallel()
	t.Run("normal test", func(t *testing.T) {
		result := addBrackets("test")
		assert.Equal(t, "(test)", result)
	})

	t.Run("error test", func(t *testing.T) {
		result := addBrackets("")
		assert.Equal(t, "", result)
	})
}

func TestEncloseString(t *testing.T) {
	t.Parallel()
	t.Run("normal test", func(t *testing.T) {
		result := encloseString("test", "*")
		assert.Equal(t, "*test*", result)
	})

	t.Run("error test", func(t *testing.T) {
		result := encloseString("", "*")
		assert.Equal(t, "", result)
	})
}

func TestEncloseSubstring(t *testing.T) {
	t.Parallel()
	t.Run("normal test", func(t *testing.T) {
		result := encloseSubstring("test", "es", "*")
		assert.Equal(t, "t*es*t", result)
	})

	t.Run("error test", func(t *testing.T) {
		result := encloseSubstring("", "es", "*")
		assert.Equal(t, "", result)

		result = encloseSubstring("test", "", "*")
		assert.Equal(t, "test", result)
	})
}

func TestBoldString(t *testing.T) {
	t.Parallel()
	t.Run("normal test", func(t *testing.T) {
		result := boldString("test")
		assert.Equal(t, "*test*", result)
	})
}

func TestItalicString(t *testing.T) {
	t.Parallel()
	t.Run("normal test", func(t *testing.T) {
		result := italicString("test")
		assert.Equal(t, "_test_", result)
	})
}

func TestBoldSubstring(t *testing.T) {
	t.Parallel()
	t.Run("normal test", func(t *testing.T) {
		result := boldSubstring("test", "es")
		assert.Equal(t, "t*es*t", result)
	})
}

func TestSelectString(t *testing.T) {
	t.Parallel()
	t.Run("true test", func(t *testing.T) {
		result := selectString(true, "test1", "test2")
		assert.Equal(t, "test1", result)
	})

	t.Run("false test", func(t *testing.T) {
		result := selectString(false, "test1", "test2")
		assert.Equal(t, "test2", result)
	})
}

func TestContainsString(t *testing.T) {
	t.Parallel()
	t.Run("normal test", func(t *testing.T) {
		result := containsString([]string{"test1", "test2"}, "test1")
		assert.True(t, result)
	})

	t.Run("not found test", func(t *testing.T) {
		result := containsString([]string{"test1", "test2"}, "test3")
		assert.False(t, result)
	})

	t.Run("error test", func(t *testing.T) {
		result := containsString([]string{"test1", "test2"}, "")
		assert.False(t, result)
	})
}
