package core

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
)

type usesPermission struct {
	Name string `xml:"name,attr"`
}

type rawManifest struct {
	XMLName        xml.Name         `xml:"manifest"`
	UsesPermission []usesPermission `xml:"uses-permission"`
}

var dangerousPermissions = []string{
	"android.permission.READ_SMS",
	"android.permission.SYSTEM_ALERT_WINDOW",
	"android.permission.REQUEST_INSTALL_PACKAGES",
	"android.permission.CALL_PHONE",
	"android.permission.RECORD_AUDIO",
	"android.permission.READ_CONTACTS",
	"android.permission.WRITE_EXTERNAL_STORAGE",
	"android.permission.READ_PHONE_STATE",
}

func AnalyzePermissions(manifestPath string) []string {
	data, err := ioutil.ReadFile(manifestPath)
	if err != nil {
		fmt.Println("[✗] Failed to read manifest:", err)
		return nil
	}

	var parsed rawManifest
	err = xml.Unmarshal(data, &parsed)
	if err != nil {
		fmt.Println("[✗] Failed to parse manifest:", err)
		return nil
	}

	var permissions []string
	fmt.Println("\n[✓] Requested Permissions:")
	for _, perm := range parsed.UsesPermission {
		p := perm.Name
		permissions = append(permissions, p)
		if isDangerous(p) {
			fmt.Printf("  \x1b[31m[!] %s\x1b[0m\n", p)
		} else {
			fmt.Printf("  %s\n", p)
		}
	}
	return permissions
}

func isDangerous(p string) bool {
	for _, d := range dangerousPermissions {
		if strings.EqualFold(p, d) {
			return true
		}
	}
	return false
}