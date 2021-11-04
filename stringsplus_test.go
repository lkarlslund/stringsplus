package stringsplus

import "testing"

func TestEqualFoldHasSuffix(t *testing.T) {
	type args struct {
		s      string
		suffix string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EqualFoldHasSuffix(tt.args.s, tt.args.suffix); got != tt.want {
				t.Errorf("EqualFoldHasSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}
