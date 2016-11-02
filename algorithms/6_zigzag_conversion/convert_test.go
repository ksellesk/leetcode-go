
package main

import (
	"testing"
)

func Test_convert(t *testing.T) {

	if convert("qwer", 3) != "qwre" {
		t.Error("Test case 0 failed")
	}

}
