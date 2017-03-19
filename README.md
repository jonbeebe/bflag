# bflag 

**bflag** is a simple Go package for building command-line interfaces.

## Example usage

    package main

    import (
        "fmt"
        "github.com/jonbeebe/bflag"
    )

    func main() {
        defer func() {
            if msg := recover(); msg != nil {
                fmt.Println(msg)
            }
        }()

        boolOpt := bflag.DefineBool("mybool", 'b', false)
        strOpt := bflag.DefineString("mystr", 's', "Hello world")
        intOpt := bflag.DefineInt("myint", 'i', 32)
        floatOpt := bflag.DefineFloat("myfloat", 'f', 16.5)
        bflag.Parse()

        // Output invalid arguments (if any were provided)
        args := bflag.Args()
        command := args[0]
        for _, v := range bflag.Invalid() {
            fmt.Println("bflagtest " + command + ": " + v)
        }

        // Print options
        fmt.Print("Options: ")
        fmt.Printf("--mybool: %t --mystr: %s --myint: %d --myfloat: %.2f\n", boolOpt.Value, strOpt.Value, intOpt.Value, floatOpt.Value)

        // Print arguments (bflag doesn't differentiate between commands and arguments)
        fmt.Print("Arguments: ")
        for _, v := range args {
            fmt.Printf("%s ", v)
        }
        fmt.Print("\n")
    }

## TODO

* Handling of `--help` and `--version` options
* Documentation
* Unit tests