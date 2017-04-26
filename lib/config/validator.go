package config

import (
    "path/filepath"
    "log"
)

func (directive *Directive)validateSrc() (error) {

    absolute, err := filepath.Abs(directive.Src)
    if err != nil {
        return err
    }

    log.Printf("Convert src %s of directive %d to absolute %s\n", directive.Src, directive.N, absolute)
    directive.Src = absolute;
    log.Printf("New directive src %s\n", directive.Src)

    return nil;
}