package godisco

import (
	"testing"
)

func TestDeactivateUser(t *testing.T) {
	client := newTestClient()
	err := DeactivateUser(client, 1)

	if err != nil {
		t.Errorf("expected: %v actual: %v", nil, err)
	}

}
