package utils

import "os"

func Remove(src string) (error) {
    
    isDir, err := IsDir(src)
    if err != nil {
        return err
    }

    if isDir {
        return os.RemoveAll(src)

    } else {
        return os.Remove(src)
    }
}