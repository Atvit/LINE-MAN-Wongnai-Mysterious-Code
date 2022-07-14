package mysterious

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMysterious_Decode(t *testing.T) {
	type testCase struct {
		name     string
		secret   string
		expected string
	}

	cases := []testCase{
		{
			name:     "decode aWFuZ25",
			secret:   "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K",
			expected: "iangnoW:NAM:ENIL:ta:su:nioJ",
		},
		{
			name:     "decode YXR0aGF",
			secret:   "YXR0aGF3aXQ6Ym9vbnJhc3NhbWVl",
			expected: "atthawit:boonrassamee",
		},
		{
			name:     "decode YmFzZTo",
			secret:   "YmFzZTo2NDpzdHJpbmc6ZW5jb2RlZA==",
			expected: "base:64:string:encoded",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			//Arrange
			m := NewMysterious()

			//Act
			sd, _ := m.Decode(c.secret)
			expected := c.expected

			//Assert
			assert.Equal(t, expected, sd)
		})
	}

	t.Run("bad base64 string", func(t *testing.T) {
		//Arrange
		m := NewMysterious()

		//Act
		_, err := m.Decode("bad_base64_string")

		//Assert
		assert.ErrorIs(t, err, ErrorIllegalData)
	})
}

func TestMysterious_Reverse(t *testing.T) {
	type testCase struct {
		name     string
		word     string
		expected string
	}

	cases := []testCase{
		{
			name:     "reverse LINE:MAN:Wongnai",
			word:     "LINE:MAN:Wongnai",
			expected: "iangnoW:NAM:ENIL",
		},
		{
			name:     "reverse atthawit:boonrassamee",
			word:     "atthawit:boonrassamee",
			expected: "eemassarnoob:tiwahtta",
		},
		{
			name:     "reverse radar:civic:radar",
			word:     "radar:civic:radar",
			expected: "radar:civic:radar",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			//Arrange
			m := NewMysterious()

			//Act
			reversed := m.Reverse(c.word)
			expected := c.expected

			//Assert
			assert.Equal(t, expected, reversed)
		})
	}
}
