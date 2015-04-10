package main

import (
    "flag"
    "fmt"
    "code.google.com/p/gcfg"
)

type Config struct {
    Guests struct {
        Guest []string
    }
}

func main() {

    var cfg Config

    progInput := flag.String("in", "", "Input file or directory")
    progOutput := flag.String("out", "", "Output location")
    cfgFile := flag.String("cfg", "cfg/govo2.gcfg", "Optional config file")
    flag.Parse()

    gcfg.ReadFileInto(&cfg, *cfgFile)
    for guest := range cfg.Guests.Guest {
        fmt.Println(cfg.Guests.Guest[guest])
    }
    

}
