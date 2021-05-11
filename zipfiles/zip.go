package zipfiles

import (
	"archive/zip"
	"io"
	"log"
	"os"
)

func ZipFile(filename string, files []string) error {

	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(newZipFile *os.File) {
		err := newZipFile.Close()
		if err != nil {

		}
	}(newZipFile)
	zipWriter := zip.NewWriter(newZipFile)
	defer func(zipWriter *zip.Writer) {
		err := zipWriter.Close()
		if err != nil {

		}
	}(zipWriter)

	for _,file := range files {
		if err = AddFileToZip(zipWriter, file); err != nil {
			return err
		}
	}

	return nil
}
func AddFileToZip(zipWriter *zip.Writer, filename string)error{
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func(fileToZip *os.File) {
		err := fileToZip.Close()
		if err != nil {
			log.Println("Could not Close the file")
		}
	}(fileToZip)

	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = filename
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err

}
