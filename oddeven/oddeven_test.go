package oddeven_test

import (
	"go-examples/oddeven"

	"testing"
)

func TestOddEvenFind(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"When input is 1 should retuen odd", args{1}, "Odd"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := oddeven.Find(test.args.input)
			if actual != test.want {
				t.Errorf("Expected %s but it got %s", test.want, actual)
			}
		})
	}
}
