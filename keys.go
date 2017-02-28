package simples

import (
	"bufio"
	"os"
	"strings"
)

// loadKeyValues ... Loads in key/value pairs (whitespace-separated) from a file.
func loadKeyValues(filename string) (map[string]string, error) {
	results := make(map[string]string)
	f, err := os.Open(filename)
	if err != nil {
		return results, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	idx := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			if strings.HasPrefix(line, "#") {
				continue
			}
			kv := strings.SplitN(line, "=", 2)
			if len(kv) == 2 {
				kv[0] = strings.ToUpper(strings.TrimSpace(kv[0]))
				kv[1] = strings.TrimSpace(kv[1])
				results[kv[0]] = kv[1]
				idx++
			}
		}
	}
	return results, nil
}
