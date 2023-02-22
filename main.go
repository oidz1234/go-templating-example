package main

import (
    "fmt"
    "os"
    "text/template"
)

type Config struct {
    Id string
    Name string
    Int string
}

func main () {

    if len(os.Args) != 4 {
        fmt.Println("You need more arguments, you need 3! 🗿")
        fmt.Println("Call the program like so:")
        fmt.Println("./gen_vlan <Vlan_Id> <Vlan_Name> <Vlan_Interface>")
        os.Exit(0)
    }

    gen := Config{os.Args[1], os.Args[2], os.Args[3]}

    tmpl, err := template.ParseFiles("vlan_template.txt")
    if err != nil {
        panic(err)
    }

    err = tmpl.Execute(os.Stdout, gen)
    if err != nil {
        panic(err)
    }







}


