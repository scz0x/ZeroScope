package core

import (
	"regexp"
	"strings"
)

func ScanForAPICalls(content string) []string {
	apiRegex := regexp.MustCompile(`Landroid\/[^;]+;->\w+\(`)
	return apiRegex.FindAllString(content, -1)
}

func ScanForSensitiveFilePaths(content string) []string {
	pathRegex := regexp.MustCompile(`/(data|sdcard|system)/[^\s'"]+`)
	return pathRegex.FindAllString(content, -1)
}

func AnalyzeSensitiveElements(content string, r *Report) {
	apiCalls := ScanForAPICalls(content)
	sensitivePaths := ScanForSensitiveFilePaths(content)
	r.APICalls = removeDuplicates(apiCalls)
	r.SensitivePaths = removeDuplicates(sensitivePaths)
}

func removeDuplicates(elements []string) []string {
	seen := map[string]bool{}
	result := []string{}
	for _, item := range elements {
		item = strings.TrimSpace(item)
		if item != "" && !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}