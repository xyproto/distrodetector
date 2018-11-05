package distrodetector

import (
	"fmt"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	dd := os.Getenv("DISTRODETECT")
	if dd != "" {
		detected := New().String()
		if detected != dd {
			t.Fatalf("%s should be %s!", detected, dd)
		}
	}
	fmt.Println(New())
}
