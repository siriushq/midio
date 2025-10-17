package lifecycle

import (
	"encoding/xml"
	"fmt"
	"testing"
)

// TestInvalidRules checks if Rule xml with invalid elements returns
// appropriate errors on validation
func TestInvalidRules(t *testing.T) {
	invalidTestCases := []struct {
		inputXML    string
		expectedErr error
	}{
		{ // Rule with ID longer than 255 characters
			inputXML: ` <Rule>
	                    <ID> babababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababab </ID>
	                    </Rule>`,
			expectedErr: errInvalidRuleID,
		},
		{ // Rule with empty ID
			inputXML: `<Rule>
							<ID></ID>
							<Filter><Prefix></Prefix></Filter>
							<Expiration>
								<Days>365</Days>
							</Expiration>
                            <Status>Enabled</Status>
	                    </Rule>`,
			expectedErr: nil,
		},
		{ // Rule with empty status
			inputXML: ` <Rule>
			                  <ID>rule with empty status</ID>
                              <Status></Status>
	                    </Rule>`,
			expectedErr: errEmptyRuleStatus,
		},
		{ // Rule with invalid status
			inputXML: ` <Rule>
			                  <ID>rule with invalid status</ID>
                              <Status>OK</Status>
	                    </Rule>`,
			expectedErr: errInvalidRuleStatus,
		},
	}

	for i, tc := range invalidTestCases {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			var rule Rule
			err := xml.Unmarshal([]byte(tc.inputXML), &rule)
			if err != nil {
				t.Fatal(err)
			}

			if err := rule.Validate(); err != tc.expectedErr {
				t.Fatalf("%d: Expected %v but got %v", i+1, tc.expectedErr, err)
			}
		})
	}
}
