package core

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ScanStrings(apkPath string) []string {
	file, err := os.Open(apkPath)
	if err != nil {
		Log.Println("String scan failed:", err)
		fmt.Println("[✗] Failed to open APK:", err)
		return nil
	}
	defer file.Close()

	var results []string
	buffer := make([]byte, 4096)

	for {
		n, err := file.Read(buffer)
		if n > 0 {
			text := string(buffer[:n])
			words := strings.FieldsFunc(text, func(r rune) bool {
				return r < 32 || r > 126
			})
			for _, w := range words {
				if len(w) > 8 {
					results = append(results, w)
				}
			}
		}
		if err == io.EOF {
			break
		}
	}

	fmt.Printf("[✓] Found %d potential strings\n", len(results))
	Log.Printf("String scan completed: %d found\n", len(results))
	return results
}