package greetings

import (
	"fmt"

	"go.temporal.io/sdk/temporal"
)

// @@@SNIPSTART samples-go-dependency-sharing-activities
type Activities struct {
	Name     string
	Greeting string
}

// GetGreeting Activity.
func (a *Activities) GetGreeting() (string, error) {
	return a.Greeting, nil
}

// @@@SNIPEND

// GetName Activity.
func (a *Activities) GetName() (string, error) {
	return a.Name, nil
}

// SayGreeting Activity.
func (a *Activities) SayGreeting(greeting string, name string) (string, error) {
	result := fmt.Sprintf("Greeting: %s %s!\n", greeting, name)
	fmt.Println(result)
	return result, temporal.NewApplicationError("no status change", "NoStatusChange", nil)
}
