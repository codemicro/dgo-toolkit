package command

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"regexp"
	"strings"
)

// Kit is the core model for command parsing and routing
type Kit struct {
	Session *discordgo.Session
	ErrorHandler func(error)
	Prefixes []string
	IsCaseSensitive bool
	DebugMode bool

	commandSet []*Command
	reactionSet []*Reaction

}

// HandleError is the internal function used to handle an error that accounts for *kit.ErrorHandler being nil
func (b *Kit) handleError(e error) {
	if b.ErrorHandler == nil {
		_, _ = fmt.Fprintf(os.Stderr, e.Error())
	} else {
		b.ErrorHandler(e)
	}
}

// AddCommand adds commands to the command set for this instance of Kit
func (b *Kit) AddCommand(commands ...*Command) {

	for _, c := range commands {
		var rx []string
		for _, x := range c.CommandText {
			rx = append(rx, regexp.QuoteMeta(x))
		}

		var isc string
		if b.IsCaseSensitive {
			isc = `(?i)`
		}

		c.detectRegexp = regexp.MustCompile(isc + `^` + strings.Join(rx, ` +`))

		b.commandSet = append(b.commandSet, c)
	}

}

func (b *Kit) CreateHandlers() {
	if b.commandSet != nil && len(b.commandSet) > 0 {
		b.Session.AddHandler(b.onMessageCreate)
	}
	// TODO: on reaction add/remove
}

// caseCompare compares two strings either with or without case sensitivity depending on the value set in the parent Kit
func (b *Kit) caseCompare(x, y string) bool {
	if b.IsCaseSensitive {
		return x == y
	}
	return strings.EqualFold(x, y)
}

// hasPrefix is an implementation of strings.HasPrefix that uses caseCompare
func (b *Kit) hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && b.caseCompare(s[0:len(prefix)], prefix)
}

// trimPrefix is an implementation of strings.TrimPrefix that uses caseCompare
func (b *Kit) trimPrefix(s, prefix string) string {
	if b.hasPrefix(s, prefix) {
		return s[len(prefix):]
	}
	return s

}