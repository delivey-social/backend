package shared_test

import (
	"testing"

	"comida.app/src/shared"
)

var validPhones = []string{
	"+ 55 (41) 99999-0000",
	"+55 (41) 99999-0000",
	"+55 41 99999-9999",
	"+55 41 9999-9999",
	"55 41 9999-9999",
	"5541999990000",
	"41 9999-9999",
	"419999-9999",
	"4199990000",
	"4199999999",
}
var invalidPhones = []string{
	"aaa",
	"xx 9999-0000",
	"9999-0000a",
	"99999-999",
	"9999-9999",
	"55419999999999",
	"+54 (41) 9999-9999",
}

type TelefoneTest struct {
	Input     string
	String    string
	Formatted string
}

var telefoneValues = []TelefoneTest{
	{
		Input:     "+55 (41) 99999-0000",
		String:    "5541999990000",
		Formatted: "+55 (41) 99999-0000",
	},
	{
		Input:     "+55 (41) 9999-0000",
		String:    "554199990000",
		Formatted: "+55 (41) 9999-0000",
	},
}

func TestTelefone(t *testing.T) {
	t.Run("Valid phones", func(t *testing.T) {
		for _, phone := range validPhones {
			_, err := shared.NewTelefone(phone)
			if err != nil {
				t.Errorf("Unexpected error for value %s: %v", phone, err)
			}
		}
	})

	t.Run("Invalid phones", func(t *testing.T) {
		for _, phone := range invalidPhones {
			phone, err := shared.NewTelefone(phone)
			if err == nil {
				t.Errorf("Expected error, got %s", phone)
			}
		}
	})

	t.Run("Phone string", func(t *testing.T) {
		for _, phone := range telefoneValues {
			final, err := shared.NewTelefone(phone.Input)
			if err != nil {
				t.Fatal(err)
			}
			if final.String() != phone.String {
				t.Errorf("Expected %s got %s", phone.String, final.String())
			}

		}
	})

	t.Run("Phone formatted", func(t *testing.T) {
		for _, phone := range telefoneValues {
			final, err := shared.NewTelefone(phone.Input)
			if err != nil {
				t.Fatal(err)
			}
			if final.Formatted() != phone.Formatted {
				t.Errorf("Expected %s got %s", phone.Formatted, final.Formatted())
			}

		}
	})

}
