package main

import (
    web "github.com/hoisie/web"
    "github.com/hoisie/mustache"
    "log"
    "os"
)

func index() string {
    return mustache.RenderFile("view/index.html", map[string]string{"msg":"123"}) 
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
