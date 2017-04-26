package config

import (
    "log"
    "fmt"
    "github.com/radovskyb/watcher"
    "time"
)

type Directive struct {
    Src, Dest, Sh string
    N             int
}

type DirectiveEvent struct {
    Event watcher.Event
}

// Validate and improve directive structure
//
func (directive *Directive)Validate() (error) {

    // Convert source path
    err := directive.validateSrc()
    if err != nil {
        return err;
    }

    return nil;
}

// First execution
//
func (directive *Directive)Execute() (error) {

    err := directive.executeSync()
    if err != nil {
        return err;
    }     
    
    err = directive.executeSh()
    if err != nil {
        return err;
    }

    return nil
}

// Watch for file changes
//
func (d *Directive)Watch(done chan struct{}, errc chan error) {

    w := watcher.New();

    err := w.AddRecursive(d.Src)
    if err != nil {
        errc <- err;
    }

    fmt.Printf("[i] Watcher created for direcrive %d\n", d.N)

    go func() {
        for {
            log.Printf("Loop in event watch directive %d", d.N)
            select {
            case event := <-w.Event:
                log.Printf("Triggered event %s of directive %d\n", event.Op, d.N)

                err := d.Event(event)
                if err != nil {
                    errc <- err
                }
            case err := <-w.Error:
                log.Printf("Error in directive %d", d.N)
                errc <- err
            case <-w.Closed:
                fmt.Printf("[i] Close direcrive %d\n", d.N)
                return
            case <-done:
                w.Close();
                return
            }
        }
    }()

    err = w.Start(time.Millisecond * 100)
    if err != nil {
        errc <- err;
    }
}

// Execute on every change
//
func (directive *Directive)Event(event watcher.Event) (error) {
    
    log.Println("Running events")
    
    err := directive.eventSync(event)
    if err != nil {
        return err;
    }

    err = directive.eventSh(event)
    if err != nil {
        return err;
    }

    return nil
}

