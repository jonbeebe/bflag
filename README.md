# bflag [![GoDoc](https://godoc.org/github.com/jonbeebe/bflag?status.png)](https://godoc.org/github.com/jonbeebe/bflag)

**bflag** is a simple [Go](https://golang.org) package for building command-line interfaces. It primarily does two things:

1. Allows you to define valid command-line options for your program.
2. Parses user-provided options and gives you all the other arguments (minus the options) so your program can react accordingly.

## What about flag?

There is already a similar official package called [flag](https://golang.org/pkg/flag/) that has many more features than bflag, but it ultimately didn't fit my needs because it only supports single-hyphen (`-flag`) options.

## Example usage

    package main

    import (
        "fmt"
        "github.com/jonbeebe/bflag"
    )

    func main() {
        // Define options for our program
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