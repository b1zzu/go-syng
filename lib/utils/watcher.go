package utils

import (
    "github.com/fsnotify/fsnotify"
    "path/filepath"
    "os"
)

func RecursiveNewWatcher(directory string) (*fsnotify.Watcher, error) {
    
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        return nil, err;
    }

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

            }

            return nil
        })

    if err != nil {
        return nil, err;
    }

    return watcher, nil
}
