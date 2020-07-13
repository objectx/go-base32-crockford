package base32crockford

import (
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	fmt.Println("Hello")

	plain, err := Decode([]byte{0, 1, 2, 3, 4})
	if err != nil {
		t.Errorf("error => %s\n", err.Error())
	} else {
		t.Errorf("plain => %s\n", plain)
	}
}
