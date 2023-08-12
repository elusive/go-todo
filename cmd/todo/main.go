package main

import (
    "fmt"
    "os"
    "strings"

    "github.com/elusive/todo"
)


const todoFileName = ".todo.json"

func main() {
    l := &todo.List{}

    // attempt to read any existing items from todoFileName
    if err := l.Get(todoFileName); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    // decide what to do based on number of args
    switch {
        case len(os.Args) == 1:
            // list todo items
            for _, item := range *l {
                fmt.Println(item.Task)
            }
        
        default:
            // concat all args with a space
            item := strings.Join(os.Args[1:], " ")
            l.Add(item)
            if err := l.Save(todoFileName); err != nil {
                fmt.Fprintln(os.Stderr, err)
                os.Exit(1)
            }
    }

}
