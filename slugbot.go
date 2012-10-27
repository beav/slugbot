package main

import "fmt"
import "flag"
import irc "github.com/fluffle/goirc/client"

func main() {
    flag.Parse() // parses the logging flags.
    c := irc.SimpleClient("slugbot", "slugbot", "slugbot :: GoIRC guts")
    // Optionally, enable SSL
    c.SSL = true

    // Add handlers to do things here!
    // e.g. join a channel on connect.
    c.AddHandler("connected",
        func(conn *irc.Conn, line *irc.Line) { conn.Join("#candlepin") })
    // And a signal on disconnect
    quit := make(chan bool)
    c.AddHandler("disconnected",
        func(conn *irc.Conn, line *irc.Line) { quit <- true })

    // Tell client to connect
    if err := c.Connect("irc.freenode.net"); err != nil {
        fmt.Printf("Connection error: %s\n", err)
    }

    // Wait for disconnect
    <-quit
}
