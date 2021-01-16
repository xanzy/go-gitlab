//
// Copyright 2021, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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
