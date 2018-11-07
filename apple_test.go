package distrodetector

import (
	"testing"
)

//func TestApple(t *testing.T) {
//	version := "10.12"
//	codename, err := codenameFromApple(version)
//	if err != nil {
//		t.Fatal(err)
//	}
//	correctCodename := "Sierra"
//	if codename != correctCodename {
//		t.Fatalf("Codename for %s should be %s, not %s!", version, correctCodename, codename)
//	}
//}

func TestLookupTable(t *testing.T) {
	version := "10.14"
	codename := AppleCodename(version)
	correctCodename := "Mojave"
	if codename != correctCodename {
		t.Fatalf("Codename for %s should be %s, not %s!", version, correctCodename, codename)
	}
}
