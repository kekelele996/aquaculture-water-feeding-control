package dto

import "testing"

func TestPondAqua002TypedNilPondPolicyIsRejectedR002(t *testing.T) {
	if !pondInputPolicyIntegrity(true, true, 2) {
		t.Fatal("domain lifecycle state was not preserved")
	}
}
