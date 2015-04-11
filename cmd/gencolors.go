package main

import (
	"colors"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var templText = `
<html>
	<body>
	{{range .}}
	<div style='text-align: center; width: 64px; height: 64px; display: inline-block; background-color: {{safe .}};'>{{.}}</div>
	{{end}}
	</body>
</html>
`

func preview(c []colors.HSL) {
	// TODO(aoeu): Implement browser preview here.
}

func toECMAScriptArray(c []string) string {
	s := fmt.Sprint("[")
	for i, cc := range c {
		if i > 0 {
			s += fmt.Sprint(",")
		}
		s += fmt.Sprintf(`"%v"`, cc)
	}
	s += fmt.Sprintln("]")
	return s
}

func main() {
	// TODO(aoeu): Add flag values as an anonymous struct
	serve := flag.Bool("preview", false, "Preview this as a web page")
	port := flag.String("port", ":8081", "The port to serve on")
	numColors := flag.Int("num", 14, "The number of colors")
	saturation := flag.Float64("saturation", 0.9, "Amount of saturation")
	dump := flag.Bool("dump", true, "Dump preview results to standard output.")
	lightness := flag.Float64("lightness", 0.5, "Amount of lightness")
	flag.Parse()

	s := make([]string, *numColors)
	for i, h := range colors.NewHSLSet(*numColors, *saturation, *lightness) {
		s[i] = fmt.Sprint(h.ToRGB())
	}
	funcMap := template.FuncMap{
		"safe": func(s string) template.CSS {
			return template.CSS(s)
		},
	}

	fmt.Println(toECMAScriptArray(s))

	if !*serve {
		return
	}
	t := template.New("colors").Funcs(funcMap)
	t, err := t.Parse(templText)
	if err != nil {
		log.Fatal(err)
	}
	if *dump {
		t.Execute(os.Stdout, &s)
		return
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := t.Execute(w, &s)
		if err != nil {
			log.Fatal(err)
		}
	})
	log.Fatal(http.ListenAndServe(*port, nil))
}
