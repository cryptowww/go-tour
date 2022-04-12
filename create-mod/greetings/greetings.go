// this package demo how to create a mod.
package greetings

import (
    "fmt"
    "errors"
    "math/rand"
    "time"
)
// Fun information as follows
//
// this func is used to say hello to one person in a rand way.
//  message := fmt.Sprintf(randomFormat(), name)
// as like up.
func Hello(name string) (string, error) {
    if name == ""{
        return "", errors.New("empty name ")
    }
    message := fmt.Sprintf(randomFormat(), name)
    //message := fmt.Sprintf(randomFormat())
    return message, nil
}

func init(){
    rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v!Well met!",
    }

    return formats[rand.Intn(len(formats))]
}


func Hellos(names []string) (map[string]string, error){
    messages := make(map[string]string)

    for _,name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil,err
        }
        messages[name] = message
    }

    return messages, nil
}
