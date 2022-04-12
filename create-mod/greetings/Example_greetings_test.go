package greetings_test

import (
    "fmt"
)

func ExampleHello(){
    message, err := greetings.hellos("jack ma")

    fmt.Println(message)

    // Output:
    // hello,jack ma
}


func ExampleHellos(){
    names := []string{"Gladys", "Samantha", "Darrin"}

    messages, err := greetings.Hellos(names)
    //message, err := greetings.Hello("")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(messages)
    
    // Output:
    // map[Darrin:Hi, Darrin. Welcome! Gladys:Hail, Gladys!Well met! Samantha:Hail, Samantha!Well met!]
}
