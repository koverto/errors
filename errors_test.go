package errors

import (
	"testing"
)

func TestNotImplemented(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want Error
	}{
		{"NotImplemented error", args{"com.koverto.errors.test"}, Error{
			Id:     "com.koverto.errors.test",
			Code:   501,
			Detail: "Not Implemented",
			Status: "Not Implemented",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotImplemented(tt.args.id); got.Error() != tt.want.Error() {
				t.Errorf("NotImplemented() error = %v, want %v", *got, tt.want)
			}
		})
	}
}
