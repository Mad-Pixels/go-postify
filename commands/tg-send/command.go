package tgsend

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/Mad-Pixels/postify"
	"github.com/urfave/cli/v2"
)

var (
	flagFromPath = "from"
	flagTgChat   = "chat-id"
	flagTgToken  = "bot-token"
	flagBlocks   = "with-blocks"

	defaultFlagBlocks = "main.md"

	usage = "prepare and send post to Telegram"
	name  = "tg-send"
)

type tmplUsage struct {
	FlagFromPath string
	FlagTgToken  string
	FlagTgChat   string
	FlagBlocks   string
}

func Command() *cli.Command {
	tmpl, err := template.New("usage").Funcs(template.FuncMap{
		"Join": strings.Join,
	}).Parse(usageTemplate)
	if err != nil {
		postify.Logger.Fatal(err)
	}

	var usageText bytes.Buffer
	err = tmpl.Execute(&usageText, tmplUsage{
		FlagFromPath: flagFromPath,
		FlagTgToken:  flagTgToken,
		FlagTgChat:   flagTgChat,
		FlagBlocks:   flagBlocks,
	})
	if err != nil {
		postify.Logger.Fatal(err)
	}
	return &cli.Command{
		Name:      name,
		Usage:     usage,
		UsageText: usageText.String(),
		Flags:     flags(),
		Action:    action,
	}
}
