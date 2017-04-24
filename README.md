# Nesyc

## Todos

* Allow to set excluded file in directive
* Multiple directive sharing same start and end directory ( Maybe global variables )

## Bugs

* If same file in multiple directive it will be updated only one and not both
* Watcher on single file doesn't work

## Build

### Build from windows for linux
```
$Env:GOOS="linux";$Env:GOARCH="amd64";go build -o .\bin\nesync .\nesync.go
```