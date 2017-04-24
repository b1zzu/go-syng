package utils

import "os"

func IsDir(path string)(bool, error) {

    fileStat, err := os.Stat(path)
    if err != nil {
        return false, err
    }

    return fileStat.IsDir(), nil;
}