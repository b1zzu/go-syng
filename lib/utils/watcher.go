package utils

import (
    "github.com/fsnotify/fsnotify"
    "path/filepath"
    "os"
    "log"
)

func RecursiveNewWatcher(directory string) (*fsnotify.Watcher, error) {
    
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        return nil, err;
    }

    // Firs add the main directory
    err = watcher.Add(directory)
    if err != nil {
        return nil, err;
    }
    log.Printf("Added directory %s to watcher", directory)
    
    // Walk trough all subdirectory and add them
    err = filepath.Walk(
        directory,
        func(fileName string, fileInfo os.FileInfo, err error) (error) {
            if err != nil {
                return err
            }

            if fileInfo.IsDir() {

                err := watcher.Add(fileName)
                if err != nil {
                    return err
                }
                log.Printf("Added directory %s to watcher", fileName)

            }

            return nil
        })

    if err != nil {
        return nil, err;
    }

    return watcher, nil
}
