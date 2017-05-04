# Syng

## Todos

* Pass the config file as param or the directory where is place
* Give a name to the directive (Name should has multiple sections like monarch/api)
* Allow to set excluded file in directive
* Multiple directive sharing same start and end directory ( Maybe global variables )
* Do not exit the script if command fail but do not restart until same changes are doing to the script if it fail (only stop failed directive, not all)
* Crate plugin system to allow dynamically extend syng

## Bugs

* Shell script ( To test ) [Done]
* Watcher on single file does not work ( To test with new version ) [Done]
* Test cross platform (Shell will for sure not work, but we can implement powershell or cmd)

## Build for Linux
```
$Env:GOOS="linux";$Env:GOARCH="amd64";go build -o .\out\syng .\syng.go
```

# Build for Windows

### Build from windows for linux
```
$Env:GOOS="linux";$Env:GOARCH="amd64";go build -o .\out\syng .\syng.go
```