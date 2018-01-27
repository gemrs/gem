package log

import (
	"fmt"
	"io"
	"regexp"
	"sync"
	"text/template"

	"github.com/fatih/color"
)

var standardFormat = "{{ print (fill 8 (print \"[\" .Level \"]\")) \" \" (print .Tag \":\" | fill 12) | color .Level}} {{.Context}} {{highlight .Message}}\n"

// TextTarget formats Records using a template, and writes to an io.Writer
type TextTarget struct {
	w   io.Writer
	fmt *template.Template
}

func NewTextTarget(w io.Writer) *TextTarget {
	target := &TextTarget{
		w: w,
	}
	target.SetFormat(standardFormat)
	return target
}

type colorFunc func(format string, a ...interface{}) string

var colorMap = map[string]colorFunc{
	"black":   color.BlackString,
	"blue":    color.BlueString,
	"cyan":    color.CyanString,
	"green":   color.GreenString,
	"magenta": color.MagentaString,
	"red":     color.RedString,
	"white":   color.WhiteString,
	"yellow":  color.YellowString,

	/* log levels */
	"INFO":   nil,
	"ERROR":  color.RedString,
	"DEBUG":  color.YellowString,
	"NOTICE": color.GreenString,

	"special": color.CyanString,
}

// doColor colors a string
func doColor(colorIf interface{}, str string) string {
	var color string
	if c, ok := colorIf.(string); ok {
		color = c
	} else if c, ok := colorIf.(Level); ok {
		color = string(c)
	}

	if fn, ok := colorMap[color]; ok && fn != nil {
		return fn(str)
	}
	return str
}

// doFill pads a string with spaces
func doFill(size int, str string) string {
	return fmt.Sprintf(fmt.Sprintf("%%-%vs", size), str)
}

var bracketedValue = regexp.MustCompile(`\[[^\]]*\]`)

// doHighlight colors special elements of a log message
func doHighlight(str string) string {
	return bracketedValue.ReplaceAllStringFunc(str, func(s string) string {
		fn := colorMap["special"]
		return "[" + fn(s[1:len(s)-1]) + "]"
	})
}

// SetFormat customizes the log template. The Record is passed as a parameter to the template.
func (t *TextTarget) SetFormat(fmt string) {
	funcMap := template.FuncMap{
		"color":     doColor,
		"fill":      doFill,
		"highlight": doHighlight,
	}
	t.fmt = template.Must(template.New("record").Funcs(funcMap).Parse(fmt))
}

var logLock sync.Mutex

// Handle formats a record and writes to the io.Writer
func (t *TextTarget) Handle(r Record) {
	logLock.Lock()
	defer logLock.Unlock()

	err := t.fmt.ExecuteTemplate(t.w, "record", r)
	if err != nil {
		panic(err)
	}
}
