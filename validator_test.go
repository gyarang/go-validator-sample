package main

import (
	"testing"
)

func TestValidator(t *testing.T) {
	tests := []struct {
		Input  *SignUp
		Expect bool
	}{
		{
			// success
			&SignUp{
				ID:       "foobar",
				Password: "Abcdefghijklmnop1!",
				Gender:   "male",
				Age:      10,
				Email:    "foo@bar.com",
			}, true,
		}, {
			// ID error - not english
			&SignUp{
				ID:       "아이디아이디아이디",
				Password: "Abcdefghijklmnop1!",
				Gender:   "male",
				Age:      10,
				Email:    "foo@bar.com",
			}, false,
		}, {
			// ID error - not enough length
			&SignUp{
				ID:       "id",
				Password: "Abcdefghijklmnop1!",
				Gender:   "male",
				Age:      10,
				Email:    "foo@bar.com",
			}, false,
		}, {
			// Password error - short
			&SignUp{
				ID:       "foobar",
				Password: "abc1!",
				Gender:   "male",
				Age:      10,
				Email:    "foo@bar.com",
			}, false,
		}, {
			// Password error - no uppercase letter
			&SignUp{
				ID:       "foobar",
				Password: "abcdefghijklmnop1!",
				Gender:   "male",
				Age:      10,
				Email:    "foo@bar.com",
			}, false,
		}, {
			// Password error - no lowercase letter
			&SignUp{
				ID:       "foobar",
				Password: "ABCDEFGHIJKLMNOP1!",
				Gender:   "male",
				Age:      10,
				Email:    "foo@bar.com",
			}, false,
		}, {
			// Password error - no number
			&SignUp{
				ID:       "foobar",
				Password: "Abcdefghijklmnop!",
				Gender:   "male",
				Age:      10,
				Email:    "foo@bar.com",
			}, false,
		}, {
			// Password error - no special
			&SignUp{
				ID:       "foobar",
				Password: "Abcdefghijklmnop1",
				Gender:   "male",
				Age:      10,
				Email:    "foo@bar.com",
			}, false,
		}, {
			// gender error
			&SignUp{
				ID:       "foobar",
				Password: "Abcdefghijklmnop1!",
				Gender:   "Attack Helicopter",
				Age:      10,
				Email:    "foo@bar.com",
			}, false,
		},
	}

	for _, tt := range tests {
		v := NewSignUpValidator()
		err := v.Struct(tt.Input)

		if err != nil == tt.Expect {
			t.Fatalf("SignUp validete test failed result expect=%t, got=%t", tt.Expect, err == nil)
		}
	}
}
