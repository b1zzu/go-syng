package cli

import (
    "log"
    "../config"
)

func RunSync(configFile string) {
    
    for {
        conf, err := config.LoadConfiguration(configFile)
        if err != nil {
            log.Fatal(err)
        }

        done := make(chan bool)
        cErr := make(chan error)
        closeWatcher := make(chan bool)

        openWatcher := 0
        for n, directive := range conf.Directives {

            err := directive.Execute()
            if err != nil {
                log.Fatalf("[-] Error with %d direcive. Error: %s\n", n + 1, err)
                continue;
            }

            go directive.RunWatcher(closeWatcher, cErr)
            openWatcher++
        }

        // Run a special watcher on
        go watchConfigFile()

        for {
            select {
            case err := <-cErr:
                log.Fatalf("[-] Error on running directive: %s\n", err)
            case <-done:
                return
            }
        }
    }
}