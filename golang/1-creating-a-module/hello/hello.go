package main

import (
    "fmt"
    "log"
    "test/greetings"
)

func main() {
    log.SetPrefix("greetings:")
    log.SetFlags(0)

    names := []string{"gladys", "john doe", "yung tog"}

    messages, err := greetings.Hellos(names)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(messages)
}
