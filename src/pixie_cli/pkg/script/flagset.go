package script

import (
	"flag"
	"fmt"
	"io"
	"strings"
)

// FlagSet is a wrapper around flag.FlagSet, because the latter
// does not support required args without a default value.
type FlagSet struct {
	baseFlagSet flag.FlagSet
	// Keeps track of which args have values (whether it is a default value or a passed in value)
	// Used to differentiate between an unset arg and an arg that has an empty default value.
	argHasValue map[string]bool
}

// NewFlagSet creates a new FlagSet.
func NewFlagSet() FlagSet {
	return FlagSet{
		argHasValue: make(map[string]bool),
	}
}

// String is a wrapper around flag.FlagSet's String function.
// It declares the presence of an argument.
// It differs from FlagSet's string in that defaultValue is allowed to be nil.
func (f *FlagSet) String(name string, defaultValue *string, usage string) {
	f.argHasValue[name] = defaultValue != nil
	if defaultValue != nil {
		f.baseFlagSet.String(name, *defaultValue, usage)
	} else {
		f.baseFlagSet.String(name, "", usage)
	}
}

// Parse wraps flag.FlagSet's Parse function to parse args.
func (f *FlagSet) Parse(arguments []string) error {
	// Get the flag values defined, so we can mark which ones are actually set.
	for _, arg := range arguments {
		// Not a flag
		if arg[0] != '-' {
			continue
		}
		splits := strings.SplitN(strings.Trim(arg, "-"), "=", 2)
		if len(splits) < 1 {
			return fmt.Errorf("Error parsing argument string: %s", arg)
		}
		f.argHasValue[splits[0]] = true
	}
	return f.baseFlagSet.Parse(arguments)
}

// Set wraps flag.FlagSet's Set function.
func (f *FlagSet) Set(name, value string) error {
	f.argHasValue[name] = true
	return f.baseFlagSet.Set(name, value)
}

// Lookup wraps flag.FlagSet's Lookup function, returning an error if we
// look up a required arg (without a default value) that hasn't been set.
func (f *FlagSet) Lookup(name string) (string, error) {
	if f.argHasValue[name] {
		return f.baseFlagSet.Lookup(name).Value.String(), nil
	}
	return "", fmt.Errorf("No value provided for argument '%s'", name)
}

// SetOutput wraps flag.FlagSet's SetOutput function.
func (f *FlagSet) SetOutput(output io.Writer) {
	f.baseFlagSet.SetOutput(output)
}

// Usage wraps flag.FlagSet's Usage function.
func (f *FlagSet) Usage() {
	f.baseFlagSet.Usage()
}
