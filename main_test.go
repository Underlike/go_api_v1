package main

import (
	"testing"

	"./config"
)

func TestMain(t *testing.T) {
	config.InitializeDatabase()
}
