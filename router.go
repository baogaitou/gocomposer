package main

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Static file for download
	r.Static("/cache", "./cache")

	// Packages.json
	r.GET("/repo/private/packages.json", HandlerPrivateFunc)
	r.GET("/packages.json", HandlerPublicFunc)

	// Package detail
	r.GET("/p2/:author/:package", HandlerPackageRequest)

	// Packeg V1
	// http://package.nxdev.cn:2048/p/provider-2013%2469d51c2ae2a799177ccd09575538c32e04853b07b38b9bf5432bf73d8d17495d.json
	r.GET("/p/:provider", HandlerPackageRequestV1)

	// Repo detail
	r.GET("/repos/:author/:package/:format/:hash", HandlerRepoDownload)

	// Simple package list
	r.GET("/dashboard", HandlerDashboard)

	return r
}
