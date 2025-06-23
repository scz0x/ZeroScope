package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ExtractArchives(root string) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		if strings.HasSuffix(strings.ToLower(info.Name()), ".zip") {
			target := path + "_extracted"
			err := os.MkdirAll(target, 0755)
			if err != nil {
				fmt.Println("✗ Failed to create dir:", target)
				return nil
			}

			r, err := zip.OpenReader(path)
			if err != nil {
				fmt.Println("✗ Failed to open zip:", path)
				return nil
			}
			defer r.Close()

			for _, f := range r.File {
				fpath := filepath.Join(target, f.Name)
				if f.FileInfo().IsDir() {
					os.MkdirAll(fpath, os.ModePerm)
					continue
				}

				os.MkdirAll(filepath.Dir(fpath), os.ModePerm)
				outFile, err := os.Create(fpath)
				if err != nil {
					continue
				}
				rc, err := f.Open()
				if err != nil {
					outFile.Close()
					continue
				}
				io.Copy(outFile, rc)
				outFile.Close()
				rc.Close()
			}

			fmt.Println("📦 Extracted:", path)
		}
		return nil
	})
}