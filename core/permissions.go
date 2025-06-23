package core

import (
	"encoding/xml"
	"os"
	"strings"
)


func ExtractPermissions(manifestPath string) []string {
	type Manifest struct {
		UsesPermissions []struct {
			Name string `xml:"name,attr"`
		} `xml:"uses-permission"`
	}

	data, err := os.ReadFile(manifestPath)
	if err != nil {
		return []string{}
	}

	var m Manifest
	if err := xml.Unmarshal(data, &m); err != nil {
		return []string{}
	}

	var result []string
	for _, perm := range m.UsesPermissions {
		if perm.Name != "" {
			parts := strings.Split(perm.Name, ".")
			result = append(result, parts[len(parts)-1]) 
		}
	}
	return result
}


func FilterDangerousPermissions(perms []string) []string {
	dangerous := map[string]bool{
		"READ_SMS":        true,
		"RECEIVE_SMS":     true,
		"SEND_SMS":        true,
		"READ_CONTACTS":   true,
		"WRITE_CONTACTS":  true,
		"RECORD_AUDIO":    true,
		"READ_PHONE_STATE": true,
		"CAMERA":          true,
		"ACCESS_FINE_LOCATION": true,
	}

	var filtered []string
	for _, p := range perms {
		if dangerous[strings.ToUpper(p)] {
			filtered = append(filtered, p)
		}
	}
	return filtered
}