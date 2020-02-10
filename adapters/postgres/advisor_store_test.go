package postgres

import (
	"testing"
	"time"
)

func TestTimezones(t *testing.T) {
	t.Log(time.Now().Format(time.RFC3339))
}
