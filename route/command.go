package route

import "regexp"

// Command represents a command that can be run by a user
type Command struct {
	Name         string
	Help         string
	CommandText  []string
	Arguments    []Argument
	Restrictions []CommandRestriction
	Run          MessageRunFunc

	detectRegexp *regexp.Regexp
}
