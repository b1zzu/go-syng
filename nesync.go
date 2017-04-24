package main

import (
    "./cli"
)

const ConfigFile string = "./nesync.yaml"

func main() {

    cli.RunSync(ConfigFile)
   
}