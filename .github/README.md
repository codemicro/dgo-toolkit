# dgo-toolkit

*A DiscordGo command parser and toolkit*

![License](https://img.shields.io/github/license/codemicro/dgo-toolkit) [![Go Reference](https://pkg.go.dev/badge/github.com/codemicro/dgo-toolkit.svg)](https://pkg.go.dev/github.com/codemicro/dgo-toolkit)

![OSS Lifecycle](https://img.shields.io/osslifecycle/codemicro/dgo-toolkit) ![Lines of code](https://img.shields.io/tokei/lines/github/codemicro/dgo-toolkit) [![Go Report Card](https://goreportcard.com/badge/github.com/codemicro/dgo-toolkit)](https://goreportcard.com/report/github.com/codemicro/dgo-toolkit)

----

`toolkit` is a collection of modules to ease the process of creating a Discord bot using Golang and the DiscordGo package. It provides command parsing and routing capabilities, reaction handler groups and other conveniences.

## Example (routing)

```go
session, _ := discordgo.New("Bot " + "<your token>")

kit := route.NewKit(session, []string{"*"})

kit.AddCommand(&route.Command{
    Name:        "Hello",
    Help:        "Say hello to someone",
    CommandText: []string{"hello"},
    Arguments: []route.Argument{
        {
            Name: "name",
            Type: route.String,
            Default: func(_ *discordgo.Session, message *discordgo.MessageCreate) (interface{}, error) {
                return message.Author.Username, nil
            },
        },
    },
    Restrictions: nil,
    Run: func(ctx *route.MessageContext) error {
        name := ctx.Arguments["name"].(string)
        _, err := ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, "Hi there "+name)
        if err != nil {
            return err
        }
        return nil
    },
})

kit.AddReaction(&route.Reaction{
    Name:  "Add notifier",
    Run: func(ctx *route.ReactionContext) error {
        fmt.Printf("ADD: %+v\n", ctx.Reaction)
        return nil
    },
    Event: route.ReactionAdd,
})

kit.CreateHandlers()

_ = session.Open()
```

## TODO

* Confirmation popups
* Pagination helper
* Middleware interface
