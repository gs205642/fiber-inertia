package inertia

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func version() string {
	if v, err := md5File("public/mix-manifest.json"); err == nil {
		return v
	}

	if v, err := md5File("public/build/manifest.json"); err == nil {
		return v
	}

	return ""
}

func md5File(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Fatal(err)
		}
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
