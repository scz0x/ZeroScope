package core

type Report struct {
	Version         string              `json:"version"`
	Permissions     []string            `json:"permissions"`
	Dangerous       []string            `json:"dangerous"`
	DangerousSet    map[string]bool     `json:"-"`
	SecretsFound    []string            `json:"secrets_found"`
	APICalls        map[string][]string `json:"api_calls"`
	SensitivePaths  []string            `json:"sensitive_paths"`
	Suspicious      []string            `json:"suspicious"`
	FileCounts      map[string]int      `json:"file_counts"`
	SizesMB         map[string]float64  `json:"sizes_mb"`
	Unclassified    []string            `json:"unclassified"`
}