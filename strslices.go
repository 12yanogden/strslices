package strslices

import (
	"regexp"
	"slices"
	"strings"
)

func Explode(text string, delimiter string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	} else {
		return strings.Split(text, delimiter)
	}
}

func ToCSV(slice []string) string {
	return strings.Join(slice, ",")
}

func CSVToSlice(csv string) []string {
	delimiter := ","

	if regexp.MustCompile(", ").Match([]byte(csv)) {
		delimiter = ", "
	}

	return Explode(csv, delimiter)
}

func Pretty(a []string) string {
	out := "[\n    "

	out += strings.Join(a, "\n    ")

	out += "\n]\n"

	return out
}

func PrettyToSlice(s string) []string {
	// Remove brackets
	s = strings.TrimLeft(s, "[\n    ")
	s = strings.TrimRight(s, "\n]")

	// Convert to array and return
	return strings.Split(s, "\n    ")
}

func Equals(a1, a2 []string) bool {
	return slices.Equal(a1, a2)
}