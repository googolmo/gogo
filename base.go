package main

import (
    web "github.com/hoisie/web"
    "github.com/hoisie/mustache"
    "log"
    "os"
    "time"
)

func index() string {
    t := time.Now()
    return mustache.RenderFile("view/index.html", map[string]string{"msg":string(t.Format("2006-01-02 15:04:05"))}) 
}


func main() {
    f, err := os.Create("server.log")
    if err != nil {
        println(err.Error())
        return
    }
    logger := log.New(f, "info.", log.Ldate|log.Ltime)
    web.Get("/", index)
    web.SetLogger(logger)
    web.Config.StaticDir = "static"
    web.Run("0.0.0.0:9001")
}
