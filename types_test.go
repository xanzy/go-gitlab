package gitlab

import (
	"testing"
	"encoding/json"
)

func TestBoolValue(t *testing.T) {
	testCases := map[string]struct {
		data     []byte
		expected bool
	}{
		"should unmarshal true as true": {
			data:     []byte("true"),
			expected: true,
		},
		"should unmarshal false as true": {
			data:     []byte("false"),
			expected: false,
		},
		"should unmarshal \"1\" as true": {
			data: []byte(`"1"`),
			expected: true,
		},
		"should unmarshal \"0\" as false": {
			data: []byte(`"0"`),
			expected: false,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
		    var b BoolValue
			if err := json.Unmarshal(testCase.data, &b); err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if bool(b) != testCase.expected {
				t.Fatalf("Expected %v but got %v", testCase.expected, b)
			}
		})
	}
}
