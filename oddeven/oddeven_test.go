package oddeven_test

import (
	"go-examples/oddeven"

	"testing"

	"github.com/stretchr/testify/assert"
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
		{"When input is 2 should retuen Even", args{2}, "Even"},
		{"When input is 3 should retuen odd", args{3}, "Odd"},
		{"When input is 4 should retuen Even", args{4}, "Even"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, oddeven.Find(test.args.input))
		})
	}
}
