package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var OriginSite = os.Getenv("mirror")
var Domain = os.Getenv("domain")

// var RunMode = os.Getenv("runmode")

func HandlerRepoDownload(c *gin.Context) {
	// log.Println("Request", c.Request.URL)
	var f = tarballInfo{}
	if err := c.ShouldBindUri(&f); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	// 如果文件存在,则redirect到静态目录
	path := fmt.Sprintf("cache/repos/%s/%s/%s", f.Author, f.PackageName, f.FileFormat)
	file := fmt.Sprintf("%s/%s.zip", path, f.FileHash)

	fe := fileExists(file)
	if !fe {
		log.Println("Cache file not found")
		// Download file
		ok, err := dirExists(path)
		if !ok {
			log.Println("Create dir err:", err)
		}

		remoteURL := fmt.Sprintf("https://api.github.com%s", c.Request.URL)
		// Download file
		dlerr := downloadFile(file, remoteURL)
		if dlerr != nil {
			log.Println("Download fail:", dlerr)
			c.String(404, "Proxy download fail.")
		}
	}

	// Redirect to Static file in cache dir
	c.Redirect(307, "/"+file)
}

func HandlerPackageRequest(c *gin.Context) {
	// log.Println("Request", c.Request.URL)
	var pkg PackageInfo
	if err := c.ShouldBindUri(&pkg); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	path := fmt.Sprintf("cache/packages/%s", pkg.Author)
	ok, err := dirExists(path)
	if !ok {
		log.Println("Create dir err:", err)
	}

	cacheFile := fmt.Sprintf("cache/packages/%s/%s", pkg.Author, pkg.PackageName)
	content, err := ioutil.ReadFile(cacheFile)
	if err != nil {
		log.Println("Cache not found, fetching...")

		// Fetch
		url := fmt.Sprintf("%s/%s", OriginSite, c.Request.URL.String())
		rt, err := downloadJSON(url)
		if err != nil {
			log.Println("Fetch resource err:", err)
		}

		// Replace download url in json
		rt = replaceTarballURL(rt)

		// Cache file
		ok, err := writeFile(cacheFile, rt)
		if !ok {
			log.Println("Write cache file err:", err)
		}
		content = rt
	}

	c.Header("Content-Type", "application/json")
	c.String(200, string(content))
}

func HandlerPublicFunc(c *gin.Context) {
	cacheFile := "cache/packages.json"
	content, err := ioutil.ReadFile(cacheFile)
	if err != nil {
		log.Println("Cache not found, fetching...")

		// Fetch
		url := fmt.Sprintf("%s/packages.json", OriginSite)
		rt, err := downloadJSON(url)
		if err != nil {
			log.Println("Fetch resource err:", err)
		}

		// Cache file
		ok, err := writeFile(cacheFile, rt)
		if !ok {
			log.Println("Write cache file err:", err)
		}
		content = rt
	}

	c.Header("Content-Type", "application/json")
	c.String(200, string(content))
}

func HandlerPrivateFunc(c *gin.Context) {
	cacheFile := "cache/packages.json"
	content, err := ioutil.ReadFile(cacheFile)
	if err != nil {
		// log.Println("Cache not found")
		url := fmt.Sprintf("%s/packages.json", OriginSite)
		rt, err := downloadJSON(url)
		if err != nil {
			log.Println("Fetch resource err:", err)
		}

		// Cache file
		ok, err := writeFile(cacheFile, rt)
		if !ok {
			log.Println("Write cache file err:", err)
		}
		content = rt
	}

	c.Header("Content-Type", "application/json")
	c.String(200, string(content))
}

// HandlerDashboard 显示管理面板
func HandlerDashboard(c *gin.Context) {
	pkgs, _ := listPackages("./cache/packages")
	htmlString := "<pre>"
	htmlString += fmt.Sprintf(`<h2>Package cached: %d</h2><p><a href="https://github.com/baogaitou/gocomposer" target="_blank">gocomposer@Github</a></p><hr>`, len(pkgs))

	for _, p := range pkgs {
		htmlString += fmt.Sprintf(`<li>[<a href="%s" target="_blank">Info</a>] %s/%s  <small style="color:#CCCCCC">%s</small>`, p.FileName, p.Author, p.PackageName, p.CacheTime.Format("2006-01-02 15:04:05"))
	}

	c.Header("Content-Type", "text/html")
	c.String(200, htmlString)
}
