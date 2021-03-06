package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/chzchzchz/sitbot/bot"
	"github.com/chzchzchz/sitbot/msl"
)

func addEvent(p *bot.Profile, ev *msl.Event) {
	pat := bot.Pattern{
		Match:    ev.Match(),
		Template: "mybot " + ev.Name() + " " + ev.Flags(),
	}
	if ev.Type == "text" {
		p.Patterns = append(p.Patterns, pat)
	} else {
		p.PatternsRaw = append(p.PatternsRaw, pat)
	}
}

func main() {
	fname := os.Args[1]
	s, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	g := msl.Grammar{Buffer: string(s), Pretty: true}
	g.Init()
	if err := g.Parse(); err != nil {
		fmt.Println(err)
	}
	//g.PrintSyntaxTree()
	g.Execute()

	prof := bot.Profile{Id: fname}
	for _, ev := range g.Events {
		addEvent(&prof, &ev)
	}
	b, err := json.Marshal(prof)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	fmt.Println("main.go--------")
	fmt.Println(`package main

import (
	"github.com/chzchzchz/sitbot/runtime"
)

func main() {
	addEvents()
	runtime.Start()
}
`)

	// Hook events.
	fmt.Println("func addEvents() {")
	for _, ev := range g.Events {
		fmt.Printf("\truntime.AddEvent(%q, %s)\n", ev.Name(), ev.Name())
	}
	fmt.Println("}\n")

	// Emit event handlers.
	for _, ev := range g.Events {
		fmt.Printf("func %s() {\n", ev.Name())
		fmt.Printf("%s", ev.Command.Emit())
		fmt.Println("}\n")
	}

}
