package route

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
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

// Duration will parse a single duration in the form 1d2h3m4s
var Duration = durationType{}

type durationType struct{}

func (durationType) Parse(content *string) (interface{}, error) {

	a, b := takeFirstPart(*content)

	// remove all whitespace and make lowercase
	a = strings.ReplaceAll(a, " ", "")
	a = strings.ToLower(a)

	// the string has to start with a digit
	if !unicode.IsDigit(rune(a[0])) {
		return 0, errors.New("ParseDuration: duration string must start with a digit")
	}

	var dur time.Duration
	var currentDigitBuffer string

	for _, char := range a {

		if unicode.IsDigit(char) {
			currentDigitBuffer += string(char)
		} else {



			var mod time.Duration

			switch char {
			case 'd':
				mod = time.Hour * 24
			case 'h':
				mod = time.Hour
			case 'm':
				mod = time.Minute
			case 's':
				mod = time.Second
			default:
				return 0, fmt.Errorf("ParseDuration: unknown unit suffix \"%s\"", string(char))
			}

			num, err := strconv.Atoi(currentDigitBuffer)
			currentDigitBuffer = ""
			if err != nil {
				return 0, err
			}

			dur += time.Duration(num) * mod

		}
	}

	if currentDigitBuffer != "" {
		return 0, fmt.Errorf("ParseDuration: value %s without suffix not allowed", currentDigitBuffer)
	}

	*content = b
	return dur, nil

}
func (durationType) Name() string { return "duration" }
func (durationType) Help(_ string) string {
	return "A duration, for example `7d1h2m3s`. Valid time units are `s`, `m`, `h` and `d`."
}
