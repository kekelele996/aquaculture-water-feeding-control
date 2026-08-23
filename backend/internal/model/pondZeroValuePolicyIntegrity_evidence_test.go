package model

import "testing"

func TestPondAqua002QuarantinePondZeroValueDoesNotPanicR002(t *testing.T) {
	if !pondZeroValuePolicyIntegrity(true, true, 2) {
		t.Fatal("domain lifecycle state was not preserved")
	}
}
