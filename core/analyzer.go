package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func AnalyzeExtractedFiles(baseDir string) (map[string]int, map[string]float64, []string) {
	counts := map[string]int{"xml": 0, "dex": 0, "so": 0, "ttf": 0, "asset": 0, "other": 0}
	sizes := map[string]float64{"xml": 0, "dex": 0, "so": 0, "ttf": 0, "asset": 0, "other": 0}
	var suspiciousSO []string

	filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		sizeMB := float64(info.Size()) / (1024 * 1024)
		l := strings.ToLower(path)

		switch {
		case strings.HasSuffix(l, ".xml"):
			counts["xml"]++
			sizes["xml"] += sizeMB
		case strings.HasSuffix(l, ".dex"):
			counts["dex"]++
			sizes["dex"] += sizeMB
		case strings.HasSuffix(l, ".so"):
			counts["so"]++
			sizes["so"] += sizeMB
			if info.Size() > 10*1024*1024 {
				fmt.Printf("\x1b[31m[!] Suspicious SO: %s (%.2f MB)\x1b[0m\n", filepath.Base(path), sizeMB)
				suspiciousSO = append(suspiciousSO, filepath.Base(path))
			}
		case strings.HasSuffix(l, ".ttf"):
			counts["ttf"]++
			sizes["ttf"] += sizeMB
		case strings.Contains(l, "asset"):
			counts["asset"]++
			sizes["asset"] += sizeMB
		default:
			counts["other"]++
			sizes["other"] += sizeMB
		}
		return nil
	})

	fmt.Println("\n[✓] File summary:")
	for k := range counts {
		fmt.Printf("  ├─ %-6s: %d file(s) (%.2f MB)\n", strings.ToUpper(k), counts[k], sizes[k])
	}

	return counts, sizes, suspiciousSO
}