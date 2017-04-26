package utils

func Sync(source, destination string) (error) {

    isDir, err := IsDir(source)
    if err != nil {
        return err
    }

    if isDir {

        err := CopyDir(source, destination)
        if err != nil {
            return err
        }

    } else {

        err := CopyFile(source, destination)
        if err != nil {
            return err
        }

    }

    return nil;
}