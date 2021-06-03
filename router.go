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
	r.GET("/p/:author/:package", HandlerPackageRequest)

	// Repo detail
	r.GET("/repos/:author/:package/:format/:hash", HandlerRepoDownload)

	// Simple package list
	r.GET("/dashboard", HandlerDashboard)

	return r
}
