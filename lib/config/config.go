package config

import (
    "io/ioutil"
    "gopkg.in/yaml.v2"
    "github.com/fsnotify/fsnotify"
    "log"
)

type Config struct {
    Directives    []Directive
    directivesLen int
}

type ConfigWatcher struct {
    File   string
    Change chan bool
    Errc   chan error
    Done   chan bool
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

func (w *ConfigWatcher)Watch() {

    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        w.Errc <- err
    }
    defer watcher.Close()

    err = watcher.Add(w.File)
    if err != nil {
        w.Errc <- err
    }

    for {
        select {
        case event := <-watcher.Events:
            if event.Op == fsnotify.Write || event.Op == fsnotify.Create {
                w.Change <- true
                return;
            }
        case err := <-watcher.Errors:
            w.Errc <- err;
            return;
        case <-w.Done:
            return;
        }
    }
}

func (c *Config)GetDirectivesLen()(int){
    if c.directivesLen == 0 {
        c.directivesLen = len(c.Directives)
    }
    return c.directivesLen
}