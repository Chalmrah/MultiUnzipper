package main

import (
	"flag"
	"fmt"
	"github.com/bodgit/sevenzip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	folderPath      string
	extract         bool
	destinationPath string
	version         bool
)

func main() {

	flag.StringVar(&folderPath, "folder", "", "Folder path")
	flag.BoolVar(&extract, "extract", false, "Extract archive")
	flag.StringVar(&destinationPath, "destination", "", "Destination path")
	flag.BoolVar(&version, "version", false, "Show version")
	flag.Parse()

	if version {
		fmt.Print("MultiUnzipper version 1.0\n")
		fmt.Print("Go version 1.23\n")
		os.Exit(0)
	}

	dir, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatalf("Error reading folder: %v", err)
	}

	// List only
	if !extract {
		listFileSizes(dir)
	}

	if extract && destinationPath != "" {
		extractAllFiles(dir, destinationPath)
	}

	if extract {
		fmt.Printf("Destination not set!")
	}

}

func extractAllFiles(dir []os.DirEntry, extractedFolder string) {

	fmt.Printf("Extracting to %v\n", extractedFolder)

	for _, entry := range dir {
		fmt.Printf("Extracting %s\n", entry.Name())
		expandCompressedArchive(filepath.Join(folderPath, entry.Name()))
	}
}

func listFileSizes(dir []os.DirEntry) {
	var compressedSize int64
	var uncompressedSize int64

	for _, entry := range dir {
		if entry.IsDir() {
			continue
		}

		if strings.HasSuffix(entry.Name(), ".7z") {

			fileSize := getUncompressedZipSize(filepath.Join(folderPath, entry.Name()))

			compressedSize += fileSize[0]
			uncompressedSize += fileSize[1]
		}
	}

	fmt.Printf("Compressed size:\t %d MB\t %d GB\n", compressedSize/(1024*1024), compressedSize/(1024*1024*1024))
	fmt.Printf("Uncompressed size:\t %d MB\t %d GB\n", uncompressedSize/(1024*1024), uncompressedSize/(1024*1024*1024))
}

func expandCompressedArchive(filePath string) {

	archive, err := sevenzip.OpenReader(filePath)
	if err != nil {
		log.Fatalf("Error reading archive: %v", err)
	}
	defer func(archive *sevenzip.ReadCloser) {
		err := archive.Close()
		if err != nil {
			log.Printf("Error closing archive: %v", err)
		}
	}(archive)

	for _, file := range archive.File {
		destFilePath := filepath.Join(destinationPath, file.Name)
		err = os.MkdirAll(filepath.Dir(destFilePath), os.ModePerm)
		if err != nil {
			log.Printf("Error creating directory %v: %v", filepath.Dir(destFilePath), err)
			continue
		}

		rc, err := file.Open()
		if err != nil {
			log.Printf("Error opening file %v: %v", file.Name, err)
		}
		defer func(rc io.ReadCloser) {
			err := rc.Close()
			if err != nil {
				log.Printf("Error closing file %v: %v", file.Name, err)
			}
		}(rc)

		destFile, err := os.Create(destFilePath)
		if err != nil {
			log.Printf("Error creating file %v: %v", file.Name, err)
		}
		defer func(destFile *os.File) {
			err := destFile.Close()
			if err != nil {
				log.Printf("Error closing file %v: %v", file.Name, err)
			}
		}(destFile)

		_, err = io.Copy(destFile, rc)
		if err != nil {
			log.Printf("Error extracting file %v: %v", file.Name, err)
		}

	}

}

func getUncompressedZipSize(filePath string) [2]int64 {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open archive: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Error closing file %v: %v", file.Name, err)
		}
	}(file)

	reader, err := sevenzip.NewReader(file, fileInfoSize(file))
	if err != nil {
		log.Fatalf("Failed to open archive: %v", err)
	}

	var totalSize int64

	for _, entry := range reader.File {
		totalSize += int64(entry.UncompressedSize)
	}

	// get actual file size
	stat, err := file.Stat()
	if err != nil {
		log.Fatalf("Failed to stat archive: %v", err)
	}

	return [2]int64{stat.Size(), totalSize}
}

func fileInfoSize(file *os.File) int64 {
	info, err := file.Stat()
	if err != nil {
		log.Fatalf("Failed to get file size: %v", err)
	}
	return info.Size()
}
