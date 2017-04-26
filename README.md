# Syng

## Todos

* Pass the config file as param or the directory where is place
* Give a name to the directive (Name should has multiple sections like monarch/api)
* Allow to set excluded file in directive
* Multiple directive sharing same start and end directory ( Maybe global variables )

## Bugs

* Shell script ( To test )
* Watcher on single file does not work ( To test with new version )

## Build

### Build from windows for linux
```
$Env:GOOS="linux";$Env:GOARCH="amd64";go build -o .\bin\nesync .\nesync.go
```