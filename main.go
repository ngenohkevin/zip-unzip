package main

import (
	"fmt"
	"github.com/ngenohkevin/zip-unzip_files/unzipfiles"
	"github.com/ngenohkevin/zip-unzip_files/zipfiles"
	"log"
	"strings"
)

func main() {
	file, err := unzipfiles.Unzip("zip.zip", "codes")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Unzipping\n" + strings.Join(file, "\n"))

	files := []string{"code"}
	output := "file1.zip"

	if err := zipfiles.ZipFile(output, files); err != nil {
		panic(err)
	}
	fmt.Println("Zipped File: ", output)
}
