package check_template

import (
	"context"
	"encoding/json"
	"fmt"
)

// Schema is a custom defined struct that will hold the check configuration
type Schema struct {
	Target string `key:"target"` // Make sure to use the scorify tags to define the key in the config
	Port   int    `key:"port"`

	// Add any additional fields that you want to pass in as config
}

// Run is the function that will get called to run an instance of a check
func Run(ctx context.Context, config string) error {
	// Define a new Schema
	schema := Schema{}

	// Unmarshal the config to the Schema
	err := json.Unmarshal([]byte(config), &schema)
	if err != nil {
		return err
	}

	// Custom logic to run the check

	fmt.Println("Running check with target:", schema.Target, "and port:", schema.Port)

	return nil
}
