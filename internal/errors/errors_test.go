package errors_test

import (
	"errors"
	"testing"

	apperrors "homestack/internal/errors"
)

func TestErrorFormatting(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "not found includes resource and id",
			err:  &apperrors.NotFoundError{Resource: "user", ID: "42"},
			want: "user with id 42 not found",
		},
		{
			name: "conflict surfaces raw message",
			err:  &apperrors.ConflictError{Message: "subdomain already taken"},
			want: "subdomain already taken",
		},
		{
			name: "validation includes field and message",
			err:  &apperrors.ValidationError{Field: "email", Message: "must be a valid address"},
			want: "validation error: email - must be a valid address",
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := tc.err.Error(); got != tc.want {
				t.Errorf("Error() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestTypedErrorsArePointerComparable(t *testing.T) {
	t.Parallel()
	// Handlers branch on *NotFoundError / *ConflictError / *ValidationError via
	// errors.As; verify each type round-trips so the contract isn't silently
	// broken by an Error() method moving onto a value receiver.
	var nf *apperrors.NotFoundError
	if !errors.As(&apperrors.NotFoundError{Resource: "vm", ID: "1"}, &nf) {
		t.Fatalf("errors.As did not match *NotFoundError")
	}
	var cf *apperrors.ConflictError
	if !errors.As(&apperrors.ConflictError{Message: "x"}, &cf) {
		t.Fatalf("errors.As did not match *ConflictError")
	}
	var ve *apperrors.ValidationError
	if !errors.As(&apperrors.ValidationError{Field: "f", Message: "m"}, &ve) {
		t.Fatalf("errors.As did not match *ValidationError")
	}
}
