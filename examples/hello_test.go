package examples

import (
    "testing"
)


func TestHello(t *testing.T) {
    res := Hello()
	if res != "Ola..." {
		t.Errorf(`Hello() = (%s) but expected 'Hello...')`, res)
	}
}
