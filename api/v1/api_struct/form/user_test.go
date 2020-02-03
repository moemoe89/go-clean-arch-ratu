//
//  Simple REST API
//
//  Copyright © 2020. All rights reserved.
//

package form_test

import (
	"simple-go-clean-arch/api/v1/api_struct/form"

	"testing"
)

func TestUserNameEmpty(t *testing.T) {
	user := &form.UserForm{}

	expected := "Name can't be empty"
	errs := user.Validate()
	if errs[0] != expected {
		t.Errorf("Should return %s, got %s", expected, errs[0])
	}
}

func TestUserInvalidEmail(t *testing.T) {
	user := &form.UserForm{
		Name:  "Momo",
		Email: "Baruno",
	}

	expected := "Invalid email address"
	errs := user.Validate()
	if errs[0] != expected {
		t.Errorf("Should return %s, got %s", expected, errs[0])
	}
}
