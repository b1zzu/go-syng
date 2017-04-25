package cli

import (
    "github.com/davbizz/go-syng/lib/config"
    "log"
)

func RunSync(configFile string) {
    
    for {
        conf, err := config.LoadConfiguration(configFile)
        if err != nil {
            log.Fatal(err)
        }

        // Directives
        // ----------
        
        done := make(chan bool)
        errc := make(chan error)

        openWatcher := 0
        for n, directive := range conf.Directives {

            err := directive.Execute()
            if err != nil {
                log.Fatalf("[-] Error with %d direcive. Error: %s\n", n + 1, err)
                continue;
            }

            go directive.RunWatcher(done, errc)
            openWatcher++
        }

        // Config
        // ------
        
        // Run a special watcher on
        w := &config.ConfigWatcher{File:configFile}
        go w.Watch();

        for {
            select {
            case err := <-w.Errc:
                log.Fatalf("[-] Error on running directive: %s\n", err)
            case <-w.Change:
                done <- true; // Send close signal to all watcher
                break; // Reload the config file
            }
        }
    }
}