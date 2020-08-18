package gitlab

import (
	"encoding/json"
	"testing"
)

func TestBoolValue(t *testing.T) {
	testCases := []struct {
		name     string
		data     []byte
		expected bool
	}{
		{
			name:     "should unmarshal true as true",
			data:     []byte("true"),
			expected: true,
		},
		{
			name:     "should unmarshal false as false",
			data:     []byte("false"),
			expected: false,
		},
		{
			name:     "should unmarshal true as true",
			data:     []byte(`"true"`),
			expected: true,
		},
		{
			name:     "should unmarshal false as false",
			data:     []byte(`"false"`),
			expected: false,
		},
		{
			name:     "should unmarshal \"1\" as true",
			data:     []byte(`"1"`),
			expected: true,
		},
		{
			name:     "should unmarshal \"0\" as false",
			data:     []byte(`"0"`),
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
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
