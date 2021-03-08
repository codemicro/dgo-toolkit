package route

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"regexp"
	"strconv"
	"strings"
)

// Argument represents an argument in a command
type Argument struct {
	Name    string
	Type    ArgumentType
	Default func(session *discordgo.Session, message *discordgo.MessageCreate) (interface{}, error)
}

// ArgumentType defines a possible argument, how to parse it and generates relevant help text
type ArgumentType interface {
	// Parse should consume the argument is has parsed
	Parse(content *string) (interface{}, error)
	// Name should return the readable name of the type, eg "integer" or "string"
	Name() string
	Help(name string) string
}

// ----- Pre-made ArgumentTypes -----

// parseQuote will parse and consume a quote surrounded string, eg "hello world" or 'hi there'
func parseQuote(content *string) (interface{}, error) {
	// TODO: quote escaping
	end := strings.Index((*content)[1:], string((*content)[0]))
	if end == -1 {
		return nil, errors.New("got an opening quotation mark but no closing quotation mark")
	}
	n := (*content)[1 : end+1]
	*content = (*content)[end+1:]
	return n, nil
}

var spaceSplitRegex = regexp.MustCompile(` +`)

// takeFirstPart will return the first section of a string when split by spaces. For example "hello  world hi" will
// return ("hello", "world hi")
func takeFirstPart(in string) (string, string) {
	xspl := spaceSplitRegex.Split(in, 2)
	var v string
	if len(xspl) > 1 {
		v = xspl[1]
	}
	return xspl[0], v
}

// String will parse a single (quote enclosed) string
var String = stringType{}

type stringType struct{}

func (stringType) Parse(content *string) (interface{}, error) {

	a, b := takeFirstPart(*content)

	// like anyone is ever going to use the quotes but ok
	if (*content)[0] == '"' || (*content)[0] == '\'' {
		return parseQuote(content)
	}

	*content = b
	return a, nil

}

func (stringType) Name() string         { return "string" }
func (stringType) Help(_ string) string { return "A string, for example `hello` or `\"hi there\"`" }

// RemainingString will parse a the remainder of the message as a string
var RemainingString = remainingStringType{}

type remainingStringType struct{}

func (remainingStringType) Parse(content *string) (interface{}, error) {

	if (*content)[0] == '"' || (*content)[0] == '\'' {
		return parseQuote(content)
	}

	n := *content
	*content = ""

	return n, nil

}
func (remainingStringType) Name() string         { return "string" }
func (remainingStringType) Help(n string) string { return String.Help(n) }

// Integer will parse a single integer
var Integer = integerType{}

type integerType struct{}

func (integerType) Parse(content *string) (interface{}, error) {

	a, b := takeFirstPart(*content)

	xi, err := strconv.Atoi(a)
	if err != nil {
		return nil, err
	}

	*content = b
	return xi, nil

}
func (integerType) Name() string         { return "integer" }
func (integerType) Help(_ string) string { return "A integer, for example `123`" }

// URL will parse a single URL
var URL = urlType{}
var urlRegex = regexp.MustCompile(`^((http[s]?|ftp):\/)?\/?([^:\/\s]+)((\/\w+)*\/)([\w\-\.]+[^#?\s]+)?(.*)?(#[\w\-]+)?$`)
type urlType struct{}

func (urlType) Parse(content *string) (interface{}, error) {

	a, b := takeFirstPart(*content)

	if !strings.HasSuffix(a, "/") {
		a += "/" // the regex only matches URLs with a trailing /
	}

	if urlRegex.MatchString(a) {
		*content = b
		return a, nil
	}
	return nil, errors.New("invalid URL")

}
func (urlType) Name() string         { return "url" }
func (urlType) Help(_ string) string { return "A URL, for example `https://www.example.com`" }

