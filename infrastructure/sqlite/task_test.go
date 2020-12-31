package sqlite

import (
	"testing"
)

func TestAgregates(t *testing.T) {

	tt := 3
	today, _, _, _, err := aggregates(tt)

	if err != nil {
		t.Errorf("Error %v", err)
	}

	t.Logf("Today %d \n", today)
}
