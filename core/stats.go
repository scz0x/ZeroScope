package core

import (
	"os"
	"path/filepath"
)


func CountFileTypes(folder string) (map[string]int, map[string]float64) {
	counts := make(map[string]int)
	sizes := make(map[string]float64)

	filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		counts[ext]++
		sizes[ext] += float64(info.Size()) / (1024 * 1024) 
		return nil
	})

	return counts, sizes
}


func DetectSuspiciousFiles(folder string) []string {
	var result []string

	filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".so" && info.Size() > 5*1024*1024 {
			result = append(result, path)
		}
		return nil
	})

	return result
}