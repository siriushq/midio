package lifecycle

import (
	"encoding/xml"
	"fmt"
	"testing"
)

// TestUnsupportedFilters checks if parsing Filter xml with
// unsupported elements returns appropriate errors
func TestUnsupportedFilters(t *testing.T) {
	testCases := []struct {
		inputXML    string
		expectedErr error
	}{
		{ // Filter with And tags
			inputXML: ` <Filter>
						<And>
							<Prefix>key-prefix</Prefix>
						</And>
						</Filter>`,
			expectedErr: errXMLNotWellFormed,
		},
		{ // Filter with Tag tags
			inputXML: ` <Filter>
						<Tag>
							<Key>key1</Key>
							<Value>value1</Value>
						</Tag>
						</Filter>`,
			expectedErr: nil,
		},
		{ // Filter with Prefix tag
			inputXML: ` <Filter>
							<Prefix>key-prefix</Prefix>
						</Filter>`,
			expectedErr: nil,
		},
		{ // Filter without And and multiple Tag tags
			inputXML: ` <Filter>
							<Prefix>key-prefix</Prefix>
							<Tag>
								<Key>key1</Key>
								<Value>value1</Value>
							</Tag>
							<Tag>
								<Key>key2</Key>
								<Value>value2</Value>
							</Tag>
						</Filter>`,
			expectedErr: errInvalidFilter,
		},
		{ // Filter with And, Prefix & multiple Tag tags
			inputXML: ` <Filter>
							<And>
							<Prefix>key-prefix</Prefix>
							<Tag>
								<Key>key1</Key>
								<Value>value1</Value>
							</Tag>
							<Tag>
								<Key>key2</Key>
								<Value>value2</Value>
							</Tag>
							</And>
						</Filter>`,
			expectedErr: nil,
		},
		{ // Filter with And and multiple Tag tags
			inputXML: ` <Filter>
							<And>
							<Prefix></Prefix>
							<Tag>
								<Key>key1</Key>
								<Value>value1</Value>
							</Tag>
							<Tag>
								<Key>key2</Key>
								<Value>value2</Value>
							</Tag>
							</And>
						</Filter>`,
			expectedErr: nil,
		},
		{ // Filter without And and single Tag tag
			inputXML: ` <Filter>
							<Prefix>key-prefix</Prefix>
							<Tag>
								<Key>key1</Key>
								<Value>value1</Value>
							</Tag>
						</Filter>`,
			expectedErr: errInvalidFilter,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			var filter Filter
			err := xml.Unmarshal([]byte(tc.inputXML), &filter)
			if err != nil {
				t.Fatalf("%d: Expected no error but got %v", i+1, err)
			}
			err = filter.Validate()
			if err != tc.expectedErr {
				t.Fatalf("%d: Expected %v but got %v", i+1, tc.expectedErr, err)
			}
		})
	}
}
