package config

import (
    "io/ioutil"
    "gopkg.in/yaml.v2"
    "log"
    "fmt"
    "time"
    "github.com/radovskyb/watcher"
)

type Config struct {
    Directives    []Directive
    directivesLen int
}

func LoadConfiguration(file string) (Config, error) {
    
    // Read Config File
    data, err := readFile(file)
    if err != nil {
        return Config{}, err
    }

    // Convert in Config Struct
    conf, err := convertToConfig(data)
    if err != nil {
        return Config{}, err
    }

    // Validate Config
    err = conf.Validate();
    if err != nil {
        return Config{}, err
    }
    log.Println(conf)
    
    return conf, nil
}

func readFile(fileName string) ([]byte, error) {
    
    data, err := ioutil.ReadFile(fileName)
    if err != nil {
        return []byte{}, err
    }
    
    return data, nil
}

func convertToConfig(data []byte) (Config, error) {
    var conf Config
    
    err := yaml.Unmarshal(data, &conf)
    if err != nil {
        return Config{}, err
    }
    
    return conf, nil
}

func (conf *Config)Validate() (error) {
    
    for i := range conf.Directives {
        
        err := conf.Directives[i].Validate();
        if err != nil {
            return err;
        }
        
    }
    log.Print(conf)
    
    return nil;
}

type ConfigWatcher struct {
    File   string
    Change chan bool
    Errc   chan error
    Done   chan bool
}

func NewConfigWatcher(file string)(*ConfigWatcher) {
    return &ConfigWatcher{
        File: file,
        Change: make(chan bool),
        Errc: make(chan error),
        Done: make(chan bool),
    }
}

func (cw *ConfigWatcher)Watch() {

    w := watcher.New();

    err := w.Add(cw.File)
    if err != nil {
        cw.Errc <- err;
    }

    fmt.Println("[i] Watcher created for config file")

    go func() {
        for {

            select {
            case e := <-w.Event:
                if e.Op == watcher.Write || e.Op == watcher.Create {
                    log.Println("Cange to config file detected!")
                    cw.Change <- true
                    return;
                }
            case err := <-w.Error:
                cw.Errc <- err;
                return;
            case <-cw.Done:
                return;
            }
        }
    }()

    err = w.Start(time.Millisecond * 1000)
    if err != nil {
        cw.Errc <- err;
    }
}

func (c *Config)GetDirectivesLen()(int){
    if c.directivesLen == 0 {
        c.directivesLen = len(c.Directives)
    }
    return c.directivesLen
}