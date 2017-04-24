package config

import (
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

type Config struct {
    Directives []Directive
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
    
    for _, directive := range conf.Directives {
        
        err := directive.Validate();
        if err != nil {
            return err;
        }
        
    }
    
    return nil;
}
