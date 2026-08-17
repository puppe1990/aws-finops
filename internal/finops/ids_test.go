package finops

import "testing"

func TestValidAWSAccountID(t *testing.T) {
	if !ValidAWSAccountID("111111111111") {
		t.Fatal("expected 12 digits to be valid")
	}
	if ValidAWSAccountID("840") {
		t.Fatal("short id should be invalid")
	}
	if ValidAWSAccountID("") {
		t.Fatal("empty id should be invalid")
	}
}

func TestSeedAWSAccountID_readsEnvOnly(t *testing.T) {
	t.Setenv(SeedAccountEnv, "111111111111")
	if got := SeedAWSAccountID(); got != "111111111111" {
		t.Fatalf("got %q", got)
	}
}

func TestSeedAWSAccountID_emptyByDefault(t *testing.T) {
	t.Setenv(SeedAccountEnv, "")
	if got := SeedAWSAccountID(); got != "" {
		t.Fatalf("default seed must be empty, got %q", got)
	}
}
