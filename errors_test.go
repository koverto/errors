package errors

import (
	"net/http"
	"testing"

	"github.com/micro/go-micro/v2/errors"
)

func TestNotImplemented(t *testing.T) {
	type args struct {
		id string
	}

	tests := []struct {
		name string
		args args
		want errors.Error
	}{
		{"NotImplemented error", args{"com.koverto.errors.test"}, errors.Error{
			Id:     "com.koverto.errors.test",
			Code:   http.StatusNotImplemented,
			Detail: "Not Implemented",
			Status: "Not Implemented",
		}},
	}

	for _, tt := range tests {
		test := tt
		t.Run(test.name, func(t *testing.T) {
			if got := NotImplemented(test.args.id); got.Error() != test.want.Error() {
				t.Errorf("NotImplemented() error = %v, want %v", *got, test.want)
			}
		})
	}
}
