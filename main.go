package main

import "github.com/alexflint/go-arg"
import "strings"

var (
	name    = "Word Templater"
	version = "1.0.0"
)

var args struct {
	Input        string   `arg:"positional,required"`
	Output       string   `arg:"positional,required"`
	ReplacePairs []string `arg:"positional"`
}

func main() {
	setupLogging()
	setupDebugLogging()
	Msg.Printf("%s %s\n\n", name, version)
	arg.MustParse(&args)
	Msg.Printf("Creating %s from %s with %s", args.Output, args.Input, args.ReplacePairs)

	docx, err := LoadDocx(args.Input)
	if err != nil {
		Error.Println(err)
	}

	replacer := strings.NewReplacer(args.ReplacePairs...)
	docx.WordXML = replacer.Replace(docx.WordXML)

	docx.WriteToFile(args.Output)
}
