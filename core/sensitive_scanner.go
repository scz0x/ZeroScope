package core

import (
	"regexp"
	"strings"
)

func ScanSensitiveStringsFromList(lines []string) []string {
	var results []string

	patterns := []*regexp.Regexp{
		regexp.MustCompile(`(?i)api[_-]?key[\s:=]{0,5}["']?[A-Za-z0-9\-_=]{12,}`),
		regexp.MustCompile(`(?i)authorization["']?\s*[:=]\s*["']?Bearer\s+[A-Za-z0-9\-._~+/]+=*`),
		regexp.MustCompile(`(?i)(access|secret)[-_ ]?(token|key)[\s:=]+[A-Za-z0-9\-_=]{12,}`),
		regexp.MustCompile(`(?i)eyJ[a-zA-Z0-9]{10,}`),
	}

	for _, line := range lines {
		if isMostlyBinary(line) {
			continue
		}
		for _, pat := range patterns {
			if pat.MatchString(line) {
				results = append(results, strings.TrimSpace(line))
				break
			}
		}
	}

	return unique(results)
}

func ScanForAPICallsGrouped(lines []string) map[string][]string {
	result := map[string][]string{
		"GET":    {},
		"POST":   {},
		"PUT":    {},
		"DELETE": {},
		"OTHER":  {},
	}

	blocklistPhrases := []string{
		"failed to get", "cannot get", "attempted to get", "could not be instantiated",
		"analytics", "fragmentmanager", "theme.appcompat", "motion", "onelink",
		"content-lenght", "app instance id", "legacy", "appcompat widget",
		"jce unlimited strength", "pinging server", "error pinging",
		"default proxies", "shapeappearance", "cached requests",
		"bucket region", "billingclient/api", "image url must be",
	}

	for _, line := range lines {
		text := strings.TrimSpace(line)
		if len(text) < 10 || len(text) > 200 || isMostlyBinary(text) {
			continue
		}

		matched, _ := regexp.MatchString(`^[A-Z]?Lcom/.+;$`, text)
		if matched {
			continue
		}

		lower := strings.ToLower(text)
		skip := false
		for _, phrase := range blocklistPhrases {
			if strings.Contains(lower, phrase) {
				skip = true
				break
			}
		}
		if skip {
			continue
		}

		switch {
		case strings.HasPrefix(lower, "get "):
			result["GET"] = append(result["GET"], text)
		case strings.HasPrefix(lower, "post "):
			result["POST"] = append(result["POST"], text)
		case strings.HasPrefix(lower, "put "):
			result["PUT"] = append(result["PUT"], text)
		case strings.HasPrefix(lower, "delete "):
			result["DELETE"] = append(result["DELETE"], text)
		case strings.Contains(lower, "http://") || strings.Contains(lower, "https://") || strings.Contains(lower, "/api/"):
			result["OTHER"] = append(result["OTHER"], text)
		}
	}

	for method := range result {
		result[method] = unique(result[method])
	}

	return result
}

func ScanForSensitivePaths(lines []string) []string {
	var results []string
	suspiciousPaths := []string{
		"/data/data/", "/sdcard/", "/mnt/", "storage/emulated/",
		"/proc/", "/dev/", "/system/", ".env", "secrets", "logcat",
	}

	for _, line := range lines {
		if isMostlyBinary(line) {
			continue
		}
		for _, kw := range suspiciousPaths {
			if strings.Contains(line, kw) {
				results = append(results, strings.TrimSpace(line))
				break
			}
		}
	}
	return unique(results)
}

func unique(input []string) []string {
	seen := make(map[string]bool)
	var output []string
	for _, v := range input {
		if !seen[v] {
			seen[v] = true
			output = append(output, v)
		}
	}
	return output
}

func isMostlyBinary(s string) bool {
	count := 0
	for _, r := range s {
		if r < 32 && r != 9 && r != 10 && r != 13 {
			count++
		}
	}
	return count > 3
}