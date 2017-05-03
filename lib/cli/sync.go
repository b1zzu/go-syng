package cli

import (
    "github.com/davbizz/go-syng/lib/config"
    "log"
    "fmt"
)

func RunSync(configFile string) {
    
    for {
        runSync(configFile)
    }
}

func runSync(configFile string) {

    conf, err := config.LoadConfiguration(configFile)
    if err != nil {
        log.Fatalf("[-] Error during loading of configuration: %s\n", err)
    }
    fmt.Println("[i] Yaml config loaded!")


    log.Println(conf)

    // Directives
    // ----------

    done := make(chan struct{})
    RunConfig(conf, done)

    // Config
    // ------

    // Run a special watcher on
    w := config.NewConfigWatcher(configFile)
    go w.Watch()

    for {
        select {
        case err := <-w.Errc:
            log.Fatalf("[-] Error on config file watching: %s\n", err)
        case <-w.Change:
            close(done) // Send close signal to all watcher
            return;      // Reload the config file
        }
    }
    
}

func RunConfig(c config.Config, done chan struct{}) {
    
    errc := make(chan error)
    donew := make(chan struct{})
    
    for n, d := range c.Directives {
        
        log.Println("Loop inside directive: ", n+1)
        d.N = n+1;

        err := d.Execute()
        if err != nil {
            log.Fatalf("[-] Error with direcive: %d. Error: %s\n", d.N, err)
            continue;
        }
        fmt.Printf("[i] Executed directive: %d\n", d.N)

        log.Println("Start watcher for directive: ", d.N)
        go d.Watch(donew, errc)
    }
    
    go func() {
        select {
        case err := <- errc:
            log.Fatalf("[-] Error in one directive: %s", err)
        case <-done:
            log.Println("Close config watcher")
            return
        }
    }()
}