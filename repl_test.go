package main
import "testing"



func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		},{
			input:"HeLLo WoRLD",
			expected: []string{"hello", "world"},
		},{
			input: "    ",
			expected: []string{},
		},
		
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("lengths don't match: got %v (len %d), expected %v (len %d)", actual, len(actual), c.expected, len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Input word: %s does not match expected word: %s", word, expectedWord)
			}
		}
	}
}