package crypto

import "testing"

func TestEncryptDecrypt_roundTrip(t *testing.T) {
	key := DeriveKey("dev-secret")
	ct, err := Encrypt(key, "AKIAEXAMPLE")
	if err != nil {
		t.Fatal(err)
	}
	pt, err := Decrypt(key, ct)
	if err != nil {
		t.Fatal(err)
	}
	if pt != "AKIAEXAMPLE" {
		t.Fatalf("got %q", pt)
	}
}
