package unzipfiles

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Unzip(src string, dest string) ([]string, error) {

	var zipNames []string

	read, err := zip.OpenReader(src)
	if err != nil {
		return zipNames, err
	}
	defer func(read *zip.ReadCloser) {
		err := read.Close()
		if err != nil {

		}
	}(read)
	for _, file := range read.File {
		fpath := filepath.Join(dest, file.Name)

		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return zipNames, fmt.Errorf("%s wrong file path", zipNames)
		}
		zipNames = append(zipNames, fpath)

		if file.FileInfo().IsDir() {
			err := os.MkdirAll(fpath, os.ModePerm)
			if err != nil {
				return nil, err
			}
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return zipNames, err
		}
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return zipNames, err
		}
		zipRead, err := file.Open()
		if err != nil {
			return zipNames, err
		}
		_, err = io.Copy(outFile, zipRead)

		_ = outFile.Close()
		if err != nil {
			return zipNames, err
		}
		err = zipRead.Close()
		if err != nil {
			return zipNames, err
		}

	}

	return zipNames, err
}

