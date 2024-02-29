package functions

import "testing"

func TestNormalizePath(t *testing.T) {
	normalPaths := map[string]string{
		"test":                 "test",
		"test test":            "test-test",
		"test test test":       "test-test-test",
		"Fen İşleri Müdürlüğü": "fen-isleri-mudurlugu",
		"İstanbul":             "istanbul",
		"":                     "birim",
		"   ":                  "birim",
		"aa":                   "aa-birim",
		"aa-":                  "aa-birim",
		" aa  ":                "aa-birim",
		"AliVeli -Delikanlı":   "aliveli-delikanli",
		"çok çok uzun birim adı sadece 63 karakter olarak sınırlandırılmalıdır. ayrıca türkçe karakter içermemelidir.": "cok-cok-uzun-birim-adi-sadece-63-karakter-olarak-sinirlandirilm",
	}

	for input, expected := range normalPaths {
		if NormalizePath(input) != expected {
			t.Errorf("Expected %s but got %s", expected, NormalizePath(input))
		}
	}
}

func TestValidatePath(t *testing.T) {
	invalidPaths := []string{
		"test*",
		"test test",
		"test test test",
		"Fen İşleri Müdürlüğü",
		"İstanbul",
		"   ",
		"aa",
		"aa-",
		" aa  ",
		"AliVeli -Delikanlı",
	}

	for _, input := range invalidPaths {
		if _, err := ValidatePath(input); err == nil {
			t.Errorf("Expected error but got nil")
		}
	}

	validPaths := []string{
		"test",
		"test-test",
		"test-test-test",
		"fen-isleri-mudurlugu",
		"istanbul",
		"birim/alt-birim",
		"birim/alt-birim/alt-alt-birim",

		"birim",
	}

	for _, input := range validPaths {
		if _, err := ValidatePath(input); err != nil {
			t.Errorf("Expected nil but got %s", err)
		}
	}
}
