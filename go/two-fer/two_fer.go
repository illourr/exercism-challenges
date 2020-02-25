/*
Package twofer provides a single function ShareWith.
This function will take a name and return a string with the message:
One for X, one for me.
*/
package twofer

import (
	"fmt"
)

// ShareWith takes a name and returns a formatted string.
// Ex: name="Dillon", output="One for Dillon, one for me."
func ShareWith(name string) string {
	// If the name is not passed it will be an empty string by default
	// If that is the case lets set name to "you"
	if name == "" {
		name = "you"
	}

	// Return the formatted sentence by interpolating the name.
	return fmt.Sprintf("One for %s, one for me.", name)
}
