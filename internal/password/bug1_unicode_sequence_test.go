package password

import "testing"

func TestBug1NonASCIIIsNotASCIISequence(t *testing.T) {
	if HasSequential([]rune("éêë")) {
		t.Fatal("non-ASCII runes must not be treated as an ASCII sequence")
	}
}
