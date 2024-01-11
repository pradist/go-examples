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
		{"When input is 5 should retuen odd", args{5}, "Odd"},
		{"When input is 6 should retuen Even", args{6}, "Even"},
		{"When input is 59 should retuen odd", args{59}, "Odd"},
		{"When input is 60 should retuen Even", args{60}, "Even"},
		{"When input is 99 should retuen odd", args{99}, "Odd"},
		{"When input is 100 should retuen Even", args{100}, "Even"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, oddeven.Find(test.args.input))
		})
	}
}
