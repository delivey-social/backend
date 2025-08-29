package restaurante_test

import (
	"testing"

	"comida.app/src/internal/restaurante"
)

var (
	validCNPJ = []string{
		"18781203/0001-28",
		"187812030001-28",
		"18781203/000128",
		"00.000.100/0002-07",
		"92.038.476/0001-30",
	}
	invalidCNPJ = []string{
		"",
		"11.444.777/0001-60",
		"18781203/000129",
		"12345678/0001-28",
		"12345678000128",
		"12.345.678/0001-28",
		"00000000/0000-00",
		"abcdefgh/ijkl-mn",
		"11111111/1111-11",
		"22222222/2222-22",
		"33333333/3333-33",
		"44444444/4444-44",
		"55555555/5555-55",
		"66666666/6666-66",
		"77777777/7777-77",
		"88888888/8888-88",
		"99999999/9999-99",
		"1234567/0001-28",
		"123456789/0001-28",
		"12345678/00012-8",
		"12345678/0001-2",
		"12345678/0001-288",
		"12345678/0001-2a",
		"abcdefghijklmno",
		"12/3456780001-28",
		"12-3456780001-28",
		"12.3456780001-28",
		"12.345.6780001-28",
		"12.345.678/000128",
		"12.345.678/0001-2",
		"12.345.678/0001-288",
		"12.345.678/0001-2a",
		"00000000/0001-00",
		"00000000/0001-01",
		"00000000/0001-02",
		"00000000/0001-03",
		"00000000/0001-04",
		"00000000/0001-05",
		"00000000/0001-06",
		"00000000/0001-07",
		"00000000/0001-08",
		"00000000/0001-09",
		"00000000/0001-10",
		"00000000/0001-11",
		"00000000/0001-12",
		"00000000/0001-13",
		"00000000/0001-14",
		"00000000/0001-15",
		"00000000/0001-16",
		"00000000/0001-17",
		"00000000/0001-18",
		"00000000/0001-19",
	}
)

type CNPJTest struct {
	Value     string
	String    string
	Formatted string
}

var valueTestCase = []CNPJTest{
	{
		Value:     "187812030001-28",
		String:    "18781203000128",
		Formatted: "18.781.203/0001-28",
	},
}

func TestCNPJ(t *testing.T) {
	t.Run("Valid CNPJs", func(t *testing.T) {
		for _, input := range validCNPJ {
			_, err := restaurante.NewCNPJ(input)
			if err != nil {
				t.Errorf("Unexpected error for %s: %v", input, err)
			}
		}
	})

	t.Run("Invalid CNPJs", func(t *testing.T) {
		for _, input := range invalidCNPJ {
			_, err := restaurante.NewCNPJ(input)
			if err == nil {
				t.Errorf("Expected error for %s, got nothing", input)
			}
		}
	})

	t.Run("String value", func(t *testing.T) {
		for _, testCase := range valueTestCase {
			cnpj, err := restaurante.NewCNPJ(testCase.Value)
			if err != nil {
				t.Fatal(err)
			}

			if cnpj.String() != testCase.String {
				t.Errorf("expected string %s got %s", testCase.String, cnpj.String())
			}
		}
	})

	t.Run("Formatted value", func(t *testing.T) {
		for _, testCase := range valueTestCase {
			cnpj, err := restaurante.NewCNPJ(testCase.Value)
			if err != nil {
				t.Fatal(err)
			}

			if cnpj.Formatted() != testCase.Formatted {
				t.Errorf("expected formatted value %s got %s", testCase.Formatted, cnpj.String())
			}
		}
	})
}
