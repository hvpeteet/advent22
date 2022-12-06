package unique_substr

import "fmt"

func FirstUniqueSubstrIndex(s string, length int) (int, error) {
	runes := []rune(s)
	for i := range runes {
		if i < length-1 {
			continue
		}
		unique := map[rune]struct{}{}
		for _, r := range runes[i-length+1 : i+1] {
			unique[r] = struct{}{}
		}
		if len(unique) == length {
			return i + 1, nil
		}
	}
	return 0, fmt.Errorf("no unique substr of length %d found", length)
}
