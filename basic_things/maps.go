package main

import (
	"fmt"
)

// Maps syntax
// map[<key type>]<value type>{...entries}

func main () {
	var m map[string]string

	// Map types are reference types, like pointers or slices, and so
	// the value of m above is nil; it doesn’t point to an initialized map.
	// A nil map behaves like an empty map when reading, but attempts to
	// write to a nil map will cause a runtime panic; don’t do that. To
	// initialize a map, use the built in make function

	m = make(map[string]string)
	// Alternative: m = map[string]string{}

	// Add new keys
	m["dog"] = "perro"
	m["cat"] = "gato"
	m["bird"] = "pajaro"
	m["turtle"] = "tortuga"

	// Retrieves a value given a key.
	// If the key does not exist, we get the type's zero value. In this case
	// the type's zero value for string is empty string.
	value := m["dog"]

	// Deletes an entry from the map
	delete(m, "dog")

	// Check for existing entry
	// In this statement, the first value (val) contains the value stored under
	// the key "dog". If the key does not exist, contains the zero value, empty string
	// in this case. The other variable, ok, is a boolean that is true if the key
	// exists in the map and false otherwise.
	val, ok := m["dog"]

	// if you don't want the value:
	_, exist := m["dog"]

	// Iterate over a map
	for key, val := range m {
		fmt.Println("key: %s, value: %s", key, val)
	}
}
