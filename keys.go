package simples

import (
	"bufio"
	"os"
	"strings"
)

// loadKeyValues ... Loads in key/value pairs (whitespace-separated) from a file.
func loadKeyValues(filename string) (Sections, error) {
	results := make(Sections)
	f, err := os.Open(filename)
	if err != nil {
		return results, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	idx := 0
	seq := 1
	section := "DEFAULT"
	results[section] = []Entry{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			if strings.HasPrefix(line, "#") {
				continue
			}
			kv := strings.SplitN(line, "=", 2)
			if len(kv) == 1 {
				if strings.HasPrefix(kv[0], "[") && strings.HasSuffix(kv[0], "]") {
					section = strings.ToUpper(kv[0])
					section = strings.TrimPrefix(section, "[")
					section = strings.TrimSuffix(section, "]")
					if _, ok := results[section]; !ok {
						results[section] = []Entry{}
						seq = 1
					} else {
						seq = len(results[section]) + 1
					}
				}
				idx++
			}
			if len(kv) == 2 {
				kv[0] = strings.TrimSpace(kv[0])
				kv[1] = strings.TrimSpace(kv[1])
				e := Entry{
					Section:  section,
					Key:      kv[0],
					KeyUpper: strings.ToUpper(kv[0]),
					Value:    kv[1],
					Sequence: seq,
				}
				results[section] = append(results[section], e)
				idx++
				seq++
			}
		}
	}
	return results, nil
}
