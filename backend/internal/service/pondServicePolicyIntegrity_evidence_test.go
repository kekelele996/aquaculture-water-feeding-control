package service

import "testing"

func TestPondAqua002PondMetadataWriteInitializesMapR002(t *testing.T) {
	if !pondServicePolicyIntegrity(true, true, 2) {
		t.Fatal("domain lifecycle state was not preserved")
	}
}
