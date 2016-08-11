package utilities

import (
	"testing"
)

func TestToInt(t *testing.T) {
	// act
	_, intValue := ToInt("996")
	// assert
	if intValue != 996 {
		t.Error("Error Converting string into Integer")
	}
}