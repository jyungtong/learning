package greetings

import (
    "fmt"
    "errors"
    "math/rand"
)

func Hello(name string) (string, error) {
    if name == "" {
        return name, errors.New("empty name")
    }

    message := fmt.Sprintf(randomGreeting(), name)
    return message, nil
}

func Hellos(names []string) (map[string]string, error) {
    messages := make(map[string]string)

    for _, name := range names {
        message, err := Hello(name)

        if err != nil {
            return nil, err
        }

        messages[name] = message
    }

    return messages, nil
}

func randomGreeting() string {
    templates := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    return templates[rand.Intn(len(templates))]
}
