package route

import "regexp"

// Command represents a command
type Command struct {
	Name         string
	Help         string
	CommandText  []string
	Arguments    []Argument
	Restrictions []CommandRestriction
	Run          MessageRunFunc

	detectRegexp *regexp.Regexp
}
