package greetings

import "fmt"





func Hello(name string)string {
	message := fmt.Sprintf("He, v% Welcome!", name)
	return message
}