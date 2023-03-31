package greetings

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)

// Hello returns a greeting for the named person.
func Tasukete(name string) (string, error) {
    if name == "" { 
        return "", errors.New("empty name")
    }
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func Tasuketes(names []string) (map[string]string, error) {
    if len(names) < 1{ 
        return nil, errors.New("empty array")
    }

    messages := make(map[string]string)

    for _, name := range names {
        message, err := Tasukete(name)
        if err != nil {
            return nil, err
        }

        messages[name] = message

    }

    return messages, nil
}

func init(){
    rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
    formats := []string{
        "Tasukete %v",
        "Sayonara %v",
        "%v-chan wa Kyou mo Kawaii~~",
    }

    return formats[rand.Intn(len(formats))]
}