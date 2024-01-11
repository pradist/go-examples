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
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, oddeven.Find(test.args.input))
		})
	}
}
