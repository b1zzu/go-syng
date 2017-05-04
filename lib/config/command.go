package config

import (
    "path"
    "github.com/davbizz/go-syng/lib/utils"
    "strings"
    "os/exec"
    "github.com/radovskyb/watcher"
    "log"
)


// Sync
// ====

func (directive *Directive)executeSync() (error) {

    src := directive.Src
    dest := directive.Dest
    log.Println("Sync files from: ", src, " to: ", dest, " for directive: ", directive.N)
    
    return utils.Sync(src, dest)
}

func (d *Directive)eventSync(e watcher.Event) (error) {


    log.Println(e.Path, d.Src)
    file := strings.Split(e.Path, d.Src)[1]
    
    log.Printf("File ready to sync %s\n", file)

    src := e.Path
    dest := path.Join(d.Dest, file)
    
    
    if e.Op == watcher.Remove {
        log.Printf("Remove vent file: %s\n", dest)
        return utils.Remove(dest)
    } else {
        log.Printf("Sync event src: %s, dest: %s\n", src, dest)
        return utils.Sync(src, dest)
    }
}


// Sh
// ==

func (d *Directive)executeSh() (error) {

    log.Println("Run shell for directive: ", d.N)
    if d.Sh == "" {
        return nil
    }

    // Execute every command in path
    cmd := exec.Command("/bin/sh", "-c", d.Sh)
    cmd.Dir = d.Dest

    log.Println("Executing Line: ", d.Sh, " in: ", d.Dest )
    out, err := cmd.Output()
    if err != nil {
        if ee, ok := err.(*exec.ExitError); ok {
            log.Println("Stderr: ", string(ee.Stderr))
        }
        return err
    }

    log.Println("Shell: ", string(out))

    return nil
}

func (d *Directive)eventSh(e watcher.Event) (error) {
    return d.executeSh()
}
