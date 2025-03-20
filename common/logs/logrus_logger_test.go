package log

import (
	"testing"
)

func TestLogrusLogger(t *testing.T) {

	tests := []struct {
		name string
	}{
		{
			name: "TestLogrusLogger",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log := NewLogrusLogger("logs", "test")
			log.Error(tt.name)
		})
	}
}
