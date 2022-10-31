package main

import (
	"github.com/hitlyl/prest/cmd"
	"github.com/hitlyl/prest/config"
)

func main() {
	config.Load()
	cmd.Execute()
}
