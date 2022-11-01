package greetings

import "fmt"





func Hello(name string)string {
	message := fmt.Sprintf("He, s% Welcome!", name)
	return message
}