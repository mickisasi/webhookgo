## Getting Started

### Installing

```sh
go get github.com/mickisasi/webhookgo
```

### Usage

Import the package into your project.

```go
import "github.com/mickisasi/webhookgo"
```

### Simple Message 
```go
package main

import (
	"github.com/mickisasi/webhookgo"
)

func main() {
	hook := webhookgo.NewWebhook("https://discord.com/api/webhooks/...").
		SetUsername("This is a username").
		SetAvatar("https://cdn.discordapp.com/embed/avatars/0.png")
	
	resp, err := hook.Send("this is a message")
	if err != nil {
		panic(err)
	}
}
```

### Custom Embed
```go
package main

import (
	"time"
	"github.com/mickisasi/webhookgo"
)

func main() {
	hook := webhookgo.NewWebhook("https://discord.com/api/webhooks/...").
		SetUsername("This is a username").
		SetAvatar("https://cdn.discordapp.com/embed/avatars/0.png")

	e := webhookgo.NewEmbed().
		SetTitle("My title here").
		SetAuthor("author", "https://cdn.discordapp.com/embed/avatars/0.png", "https://google.com").
		SetURL("https://google.com").
		AddField("test", "test value", false).
		AddField("test", "this is inline", true).
		SetColor(0x00b0f4).
		SetDescription("This is a description").
		SetImage("https://cdn.discordapp.com/embed/avatars/0.png").
		SetFooter("This is a footer", "https://cdn.discordapp.com/embed/avatars/0.png").
		SetTimestamp(time.Now())

	resp, err := hook.SendEmbed(e)
	if err != nil {
		panic(err)
	}
}
```

## This is a personal project, features only added when needed.
