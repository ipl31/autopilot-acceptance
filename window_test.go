package window

import "testing"

func TestContains(t *testing.T) {
	for _, tt := range []struct {
		name               string
		first, last, value int
		want               bool
	}{
		{name: "lower bound", first: -3, last: 3, value: -3, want: true},
		{name: "interior", first: -3, last: 3, value: 0, want: true},
		{name: "below", first: -3, last: 3, value: -4},
		{name: "above", first: -3, last: 3, value: 4},
		{name: "reversed bounds", first: 3, last: -3, value: 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.first, tt.last, tt.value); got != tt.want {
				t.Errorf("Contains(%d, %d, %d) = %t, want %t", tt.first, tt.last, tt.value, got, tt.want)
			}
		})
	}
}
