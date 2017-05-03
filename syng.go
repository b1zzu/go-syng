package main

import (
    "github.com/davbizz/go-syng/lib/cli"
)

const ConfigFile string = "./syng.yaml"

func main() {
    cli.RunSync(ConfigFile)
}