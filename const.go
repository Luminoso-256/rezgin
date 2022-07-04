package main

import "embed"

const (
	NAME     = "rezgin_test"
	VERSION  = "0.0"
	DEV_MODE = true
)

//go:embed data/*
var embedFS embed.FS
