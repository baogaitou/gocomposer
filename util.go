package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/levigross/grequests"
)

func downloadJSON(url string) ([]byte, error) {
	log.Println("Request:", url)
	ro := &grequests.RequestOptions{
		Headers: map[string]string{
			"User-Agent": "Composer/2.0.14 (Darwin; 20.4.0; PHP 7.3.24; cURL 7.64.1)",
		},
	}
	r, err := grequests.Get(url, ro)
	return r.Bytes(), err
}

func writeFile(fileName string, content []byte) (bool, error) {
	err2 := ioutil.WriteFile(fileName, content, 0666)
	if err2 != nil {
		log.Println("Write cache file err:", err2)
	}
	log.Println("Write cache OK:", fileName)
	return true, nil
}

// IsDirExist
func dirExists(path string) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		errMk := os.MkdirAll(path, 0775)
		if errMk != nil {
			log.Println("os.MkdirAll err:", path)
			return false, errMk
		}
		errCm := os.Chmod(path, 0775)
		if errCm != nil {
			log.Println("os.Chmod err:", path)
			return false, errCm
		}
	}
	return true, nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func replaceTarballURL(content []byte) []byte {
	rt := strings.Replace(string(content), "https://api.github.com", os.Getenv("domain"), -1)
	return []byte(rt)
}

func downloadFile(filepath string, url string) error {
	log.Println("[DL] Start download file:", url)
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	// log.Println("[DL] Create file:", filepath)
	if err != nil {
		log.Println("[DL] Create file err:", err)
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func listPackages(packageDir string) ([]PackageInfo, error) {
	var pkg []PackageInfo
	err := filepath.Walk(packageDir, func(path string, info os.FileInfo, err error) error {

		// Without dir
		if !info.IsDir() {
			r := strings.Split(path, "/")

			p := PackageInfo{}
			p.Author = r[2]
			p.PackageName = strings.Replace(r[3], ".json", "", -1)
			p.FileName = path
			p.CacheTime = info.ModTime()

			pkg = append(pkg, p)
		}
		return nil
	})

	if err != nil {
		log.Println("filepath.Walk err:", err)
	}

	return pkg, nil
}
