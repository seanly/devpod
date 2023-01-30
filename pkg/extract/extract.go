package extract

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"github.com/pkg/errors"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

func Extract(origReader io.Reader, destFolder string) error {
	bufioReader := bufio.NewReader(origReader)
	testBytes, err := bufioReader.Peek(2) //read 2 bytes
	if err != nil {
		return err
	}

	// is gzipped?
	var reader io.Reader
	if testBytes[0] == 31 && testBytes[1] == 139 {
		gzipReader, err := gzip.NewReader(bufioReader)
		if err != nil {
			return errors.Errorf("error decompressing: %v", err)
		}
		defer gzipReader.Close()

		reader = gzipReader
	} else {
		reader = bufioReader
	}

	tarReader := tar.NewReader(reader)
	for {
		shouldContinue, err := extractNext(tarReader, destFolder)
		if err != nil {
			return errors.Wrap(err, "decompress")
		} else if !shouldContinue {
			return nil
		}
	}
}

func extractNext(tarReader *tar.Reader, destFolder string) (bool, error) {
	header, err := tarReader.Next()
	if err != nil {
		if err != io.EOF {
			return false, errors.Wrap(err, "tar reader next")
		}

		return false, nil
	}

	relativePath := getRelativeFromFullPath("/"+header.Name, "")
	outFileName := path.Join(destFolder, relativePath)
	baseName := path.Dir(outFileName)

	// Check if newer file is there and then don't override?
	if err := os.MkdirAll(baseName, 0755); err != nil {
		return false, err
	}

	if header.FileInfo().IsDir() {
		if err := os.MkdirAll(outFileName, 0755); err != nil {
			return false, err
		}

		return true, nil
	}

	// Create / Override file
	outFile, err := os.Create(outFileName)
	if err != nil {
		// Try again after 5 seconds
		time.Sleep(time.Second * 5)
		outFile, err = os.Create(outFileName)
		if err != nil {
			return false, errors.Wrapf(err, "create %s", outFileName)
		}
	}

	defer outFile.Close()

	if _, err := io.Copy(outFile, tarReader); err != nil {
		return false, errors.Wrapf(err, "io copy tar reader %s", outFileName)
	}
	if err := outFile.Close(); err != nil {
		return false, errors.Wrapf(err, "out file close %s", outFileName)
	}

	// Set permissions
	_ = os.Chmod(outFileName, header.FileInfo().Mode())

	// Set mod time from tar header
	_ = os.Chtimes(outFileName, time.Now(), header.FileInfo().ModTime())

	return true, nil
}

func getRelativeFromFullPath(fullpath string, prefix string) string {
	return strings.TrimPrefix(strings.ReplaceAll(strings.ReplaceAll(fullpath[len(prefix):], "\\", "/"), "//", "/"), ".")
}
