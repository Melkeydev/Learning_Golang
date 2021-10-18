package validator

import (
  "regexp"
)

// This just confirms that a valid email was passed into a respective field
var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type Validator struct {
  Errors map[string]string
}

// Helper which creates a new validator instance 
func New() *Validator {
  return &Validator{Errors: make(map[string]string)}
} 

// Helper which returns valid if the error map is nil
func (v *Validator) Valid() bool{
  return len(v.Errors) == 0
}

// Helper that adds an error message to the map
func (v *Validator) AddError(key, message string) {
  if _, exists := v.Errors[key]; !exists {
    v.Errors[key] = message
  }
}

// Helper that checks if there is a an error message in the map
func (v *Validator) Check(ok bool, key, message string) {
  if !ok {
    v.AddError(key, message)
  }
}

// Helper that returns a specific value if it is in a list of strings
func In(value string, list ...string) bool {
  for i := range list {
    if value == list[i] {
      return true
    }
  }
  return false
}

// Matches return true if a string values matches a specific regexp pattern
func Matches(value string, rx *regexp.Regexp) bool {
  return rx.MatchString(value)
}

// Unique returns true if all string values in slice are unique 
func Unique(values []string) bool {
  uniqueValues := make(map[string]bool)
  for _, value := range values {
    uniqueValues[value] = true
  }

  return len(values) == len(uniqueValues)
}
