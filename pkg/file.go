package zeus

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Returns the size of PATH file
func Size(path string) (int64, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

// Open FILE path and returns it's sha256sum
func Sha256sum(file string) (hash string, err error) {
	h := sha256.New()
	fd, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer fd.Close()
	io.Copy(h, fd)
	return fmt.Sprintf("%X", h.Sum(nil)), err
}

// Returns TRUE if PATH exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// Returns index Directory of a Tarball
func TarIndex(path string) (string, error) {
	fi, err := os.Open(path)
	if err != nil {
		return "", err
	}
	gz, err := gzip.NewReader(fi)
	if err != nil {
		return "", err
	}
	tr := tar.NewReader(gz)
	for true {
		header, err := tr.Next()
		if err != nil {
			return "", err
		}
		if header.Typeflag == tar.TypeDir {
			return filepath.Clean(header.Name), err
		}
	}
	return "", fmt.Errorf("Tar index directory not found")
}
