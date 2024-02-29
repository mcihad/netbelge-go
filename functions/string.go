package functions

import (
	"fmt"
	"regexp"
	"strings"
)

func NormalizePath(name string) string {
	trMap := strings.NewReplacer("ğ", "g", "Ğ", "g", "ı", "i", "İ", "i", "ö", "o", "Ö", "o", "ü", "u", "Ü", "u", "ş", "s", "Ş", "s", "ç", "c", "Ç", "c")
	letters := "abcdefghijklmnopqrstuvwxyz0123456789-"
	path := trMap.Replace(name)
	path = strings.ToLower(path)
	path = strings.ReplaceAll(path, " ", "-")
	parts := strings.Split(path, "-")
	var newPath []string
	for _, part := range parts {
		if part != "" {
			newPath = append(newPath, strings.ReplaceAll(part, "-", ""))
		}
	}

	// remove non-letter characters
	for i, part := range newPath {
		var newPart string
		for _, letter := range part {
			if strings.Contains(letters, string(letter)) {
				newPart += string(letter)
			}
		}
		newPath[i] = newPart
	}
	path = strings.Join(newPath, "-")
	if strings.HasPrefix(path, "/") || strings.HasPrefix(path, "-") {
		path = path[1:]
	}
	if strings.HasSuffix(path, "/") || strings.HasSuffix(path, "-") {
		path = path[:len(path)-1]
	}
	if len(path) > 63 {
		path = path[:63]
	}
	if len(path) < 3 {
		if len(path) == 0 {
			path = "birim"
		} else {
			path = path + "-birim"
		}
	}
	return path
}

func ValidatePath(value string) (string, error) {
	// Rules for path:
	// - Must be unique
	// - Must be at most 63 characters
	// - Must be at least 3 characters
	// - Must not contain special characters only letters, numbers and hyphen and /
	// - Must not start with hyphen or /
	// - Must not end with hyphen or /
	// - Must have placeholder value {yil} {ay} {gun} {saat} {dakika} {saniye} {belge_turu} {belge_no}
	// - Must not contain Turkish characters
	allowedChars := "abcdefghijklmnopqrstuvwxyz0123456789/-"
	// replace {yil} {ay} {gun} {saat} {dakika} {saniye} {belge_turu} {belge_no}
	placeholders := []string{
		"{yil}",
		"{ay}",
		"{gun}",
		"{saat}",
		"{dakika}",
		"{saniye}",
		"{belge_turu}",
		"{belge_no}",
	}

	re := regexp.MustCompile(`{\w+}`)
	foundPlaceholders := re.FindAllString(value, -1)
	// check if all placeholders are valid
	for _, placeholder := range foundPlaceholders {
		if !contains(placeholders, placeholder) {
			return "", fmt.Errorf("geçersiz değişken: %s", placeholder)
		}
	}

	value = strings.ReplaceAll(value, "{yil}", "yil")
	value = strings.ReplaceAll(value, "{ay}", "ay")
	value = strings.ReplaceAll(value, "{gun}", "gun")
	value = strings.ReplaceAll(value, "{saat}", "saat")
	value = strings.ReplaceAll(value, "{dakika}", "dakika")
	value = strings.ReplaceAll(value, "{saniye}", "saniye")
	value = strings.ReplaceAll(value, "{belge_turu}", "belge-turu")
	value = strings.ReplaceAll(value, "{belge_no}", "belge-no")

	if len(value) > 63 {
		return "", fmt.Errorf("dosya yolu en fazla 63 karakter olmalıdır")
	}
	if len(value) < 3 {
		return "", fmt.Errorf("dosya yolu en az 3 karakter olmalıdır")
	}

	if value[0] == '-' || value[0] == '/' {
		return "", fmt.Errorf("dosya yolu - veya / ile başlayamaz")
	}
	if value[len(value)-1] == '-' || value[len(value)-1] == '/' {
		return "", fmt.Errorf("dosya yolu - veya / ile bitemez")
	}
	for _, char := range value {
		if !strings.Contains(allowedChars, string(char)) {
			return "", fmt.Errorf("dosya yolu sadece harf, rakam, - ve / içerebilir")
		}
	}
	return value, nil
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
