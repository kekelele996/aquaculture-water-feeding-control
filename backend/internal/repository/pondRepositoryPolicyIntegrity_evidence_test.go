package repository

import "testing"

func TestPondAqua002PondFallbackPolicyRemainsActiveR002(t *testing.T) {
	if !pondRepositoryPolicyIntegrity(true, true, 2) {
		t.Fatal("domain lifecycle state was not preserved")
	}
}
