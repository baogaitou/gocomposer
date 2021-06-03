package main

import "time"

// PackageInfo
type PackageInfo struct {
	Author      string    `uri:"author" binding:"required" json:"author"`
	PackageName string    `uri:"package" binding:"required" json:"package_name"`
	FileName    string    `json:"file_name"`
	CacheTime   time.Time `json:"cache_time"`
}

// tarballInfo
type tarballInfo struct {
	Author      string `uri:"author" binding:"required"`
	PackageName string `uri:"package" binding:"required"`
	FileFormat  string `uri:"format"`
	FileHash    string `uri:"hash"`
}

// DefaultPackages packages.json
type DefaultPackages struct {
	Packages         []interface{}       `json:"packages"`
	NotifyBatch      string              `json:"notify-batch"`
	ProvidersURL     string              `json:"providers-url"`
	MetadataURL      string              `json:"metadata-url"`
	SearchURL        string              `json:"search"`
	PackageList      string              `json:"list"`
	ProvidersAPI     string              `json:"providers-api"`
	WarningMessage   string              `json:"warning"`
	WarningVersions  string              `json:"warning-versions"`
	ProviderIncludes map[string]fileHash `json:"provider-includes"`
}

// fileHash fileHash
type fileHash struct {
	Sha256 string `json:"sha256"`
}

type PackageProviders struct {
	Path string `uri:"provider" binding:"required" json:"provider"`
}

type pRequest struct {
	Package string `uri:"package" binding:"required" json:"package"`
	Path    string `uri:"path" json:"path"`
}
