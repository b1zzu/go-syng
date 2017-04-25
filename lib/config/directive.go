package config

import (
    "github.com/davbizz/go-syng/lib/utils"
    "log"
    "errors"
    "strings"
    "os/exec"
    "path"
    "path/filepath"
    "os"
    "github.com/fsnotify/fsnotify"
)

type Directive struct {
    Src, Dest, Sh string
}

func (directive *Directive)Execute() (error) {

    err := directive.SyncFiles("")
    if err != nil {
        return err;
    }

    if directive.Sh != "" {

        err := directive.RunShell()
        if err != nil {
            log.Fatal(err);
        }
    }

    return nil
}

func (directive *Directive)RunShell() (error) {

    if directive.Sh == "" {
        return errors.New("Impossible to run empty shell command!")
    }

    // Iterate over every new line
    for _, line := range strings.Split(directive.Sh, "\n") {

        // Skip if empty line
        if line == "" {
            continue;
        }

        err := runShellLine(directive.Dest, line)
        if err != nil {
            return err;
        }
    }

    return nil
}

func runShellLine(directory, line string) (error) {

    if line == "" {
        return errors.New("Impossible to run empty shell command!")
    }

    // Get the command to execute dived by args
    args := strings.Split(line, " ")

    // Execute every command in path
    cmd := exec.Command(args[0], args[1:]...)
    cmd.Dir = directory

    err := cmd.Run()
    if err != nil {
        return err
    }

    return nil
}

func (directive *Directive)SyncFiles(file string) (error) {

    source := directive.Src;
    destination := directive.Dest;

    if file != "" {
        source = path.Join(source, file)
        destination = path.Join(source, destination)
    }

    isDir, err := utils.IsDir(source);
    if err != nil {
        return err
    }
    
    if isDir {

        err := utils.CopyDir(source, destination)
        if err != nil {
            return err
        }

    } else {

        err := utils.CopyFile(source, destination)
        if err != nil {
            return err
        }

    }

    return nil
}

func (directive *Directive)Validate() (error) {

    // Convert source path
    err := directive.convertSrcToAbsolute()
    if err != nil {
        return err;
    }

    return nil;
}

func (directive *Directive)convertSrcToAbsolute() (error) {

    absolute, err := filepath.Abs(directive.Src)
    if err != nil {
        return err
    }

    directive.Src = absolute;

    return nil;
}

func (directive *Directive)RunWatcher(closeWatcher chan bool, cErr chan error) {
    err := directive.runWatcher(closeWatcher)
    if err != nil {
        cErr <- err;
    }
}

// Execute in async
func (directive *Directive)runWatcher(closeWatcher chan bool) (error) {

    watcher, err := utils.RecursiveNewWatcher(directive.Src)
    if err != nil {
        return err
    }
    defer watcher.Close()

    for {
        select {
        case event := <-watcher.Events:
            if event.Op == fsnotify.Write || event.Op == fsnotify.Create {
                
                // Execute the directive for the triggered event
                err := directive.ExecuteEvent(event)
                if err != nil {
                    return err;
                }
                
                // If has been created a directory then add it to the watcher
                sourceInfo, _ := os.Stat(event.Name)
                if event.Op == fsnotify.Create && sourceInfo.IsDir() {
                    watcher.Add(event.Name)
                    if err != nil {
                        return err
                    }
                }
                
            }
        case err := <-watcher.Errors:
            return err;
        case <-closeWatcher:
            return nil
        }
    }

}

func (directive *Directive)ExecuteEvent(event fsnotify.Event)(error) {

    file := strings.TrimLeft(event.Name, directive.Src);

    err := directive.SyncFiles(file)
    if err != nil {
        return err;
    }

    if directive.Sh != "" {

        err := directive.RunShell()
        if err != nil {
            return err;
        }
    }

    return nil
}