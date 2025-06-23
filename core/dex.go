package core

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

func ExtractDexStrings(folder string) []string {
	var results []string
	re := regexp.MustCompile(`[\x20-\x7E]{6,}`) 

	filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(path) != ".dex" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		matches := re.FindAll(data, -1)
		for _, match := range matches {
			results = append(results, string(match))
		}
		return nil
	})

	return results
}