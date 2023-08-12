package main_test

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "runtime"
    "strings"
    "testing"
)


var (
    binName   = "todo"
    fileName  = ".todo.json"
)


// function use to prepare for tests run, in 
// this case we need to build the binary
func TestMain(m *testing.M) {
    fmt.Println("Building tool...")

    if runtime.GOOS == "windows" {
        binName += ".exe"
    }

    build := exec.Command("go", "build", "-o", binName)

    if err := build.Run(); err != nil {
        fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
        os.Exit(1)
    }

    fmt.Println("Running Tests...")
    result := m.Run()

    fmt.Println("Cleaning Up...")
    os.Remove(binName)
    os.Remove(fileName)

    os.Exit(result)
}


func TestTodoCli(t *testing.T) {
    task := "test task number 1"

    dir, err := os.Getwd()
    if err != nil {
        t.Fatal(err)
    }

    cmdPath := filepath.Join(dir, binName)

    // create test for adding a task
    t.Run("TestAddNewTask", func(t *testing.T) {
        cmd := exec.Command(cmdPath, strings.Split(task, " ")...)
        if err := cmd.Run(); err != nil {
            t.Fatal(err)
        }
    })

    // create test for listing tasks
    t.Run("TestListTasks", func(t *testing.T) {
        cmd := exec.Command(cmdPath)
        out, err := cmd.CombinedOutput()
        if err != nil {
            t.Fatal(err)
        }

        expected := task + "\n"
        if expected != string(out) {
            t.Errorf("Expected %q, got %q instead\n", expected, string(out))
        }
    })
}

