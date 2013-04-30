package main

import (
    "flag"
    web "github.com/hoisie/web"
    "github.com/hoisie/mustache"
    "log"
    "os"
    "time"
    "fmt"
    "encoding/json"
)

const (
    LOG = "log"
    PORT = "port"
    HOST = "host"
)

type Config struct {
    Host string
    Port string
    Log string
}

func index() string {
    t := time.Now()
    return mustache.RenderFile("view/index.html", map[string]string{"msg":string(t.Format("2006-01-02 15:04:05"))}) 
}


func readConfig(path string)  (config Config, err error) {
    file, err := os.Open(path)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer file.Close()
    dec := json.NewDecoder(file)

    err = dec.Decode(&config)
    if err != nil {
        return 
    }
    return
}

func main() {

    logPath := "server.log"
    port := "9001"
    host := "127.0.0.1"

    config, err := readConfig("config.json")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Printf("%+v\n", config) 
        logPath = config.Log
        port = config.Port
        host = config.Host
    }
    println(logPath) 

    //lp = flag.String(LOG, *logPath, "log file location")
    //pt = flag.String(PORT, *port, "listen port")
    //ht = flag.String(HOST, *host, "listen host")
    flag.Parse()
    fmt.Println(logPath)
    fmt.Println(port)
    fmt.Println(host)
    f, err := os.Create(logPath)
    if err != nil {
        println(err.Error())
        return
    }
    logger := log.New(f, "", log.Ldate|log.Ltime)
    

    web.Get("/", index)

    web.SetLogger(logger)
    web.Config.StaticDir = "static"
    web.Run(host + ":" + port)
}
