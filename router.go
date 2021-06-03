package main

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// r.UseRawPath = true
	// r.UnescapePathValues = false

	// Static file for download
	r.Static("/cache", "./cache")

	// Packages.json
	r.GET("/repo/private/packages.json", HandlerPrivateFunc)
	r.GET("/packages.json", HandlerPublicFunc)

	// Package detail
	r.GET("/p2/:author/:package", HandlerPackageRequest)

	// Repo detail
	r.GET("/repos/:author/:package/:format/:hash", HandlerRepoDownload)

	// Simple package list
	r.GET("/dashboard", HandlerDashboard)

	// Support composer v1.x (WIP)
	// r.Static("/p", "cache/p")
	// r.GET("/p/:author/:package_with_hash", HandlerPackageRequestV1)
	r.GET("/p/:package", HandlerComposerV1Request)
	r.GET("/p/:package/:path", HandlerComposerV1Request)

	return r
}
