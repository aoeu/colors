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

func newHTMLTemplate() (*template.Template, error) {
	funcMap := template.FuncMap{
		"safe": func(s string) template.CSS {
			return template.CSS(s)
		},
	}
	t := template.New("colors").Funcs(funcMap)
	return t.Parse(templText)
}

func genHSLColors(numColors int, saturation, lightness float64) []string {
	s := make([]string, numColors)
	for i, h := range colors.NewHSLSet(numColors, saturation, lightness) {
		s[i] = fmt.Sprint(h.ToRGB())
	}
	return s
}

func genKellySafe() []string {
	s := make([]string, len(colors.KellySafe))
	i := 0
	for _, c := range colors.KellySafe {
		s[i] = c.String()
		i++
	}
	return s
}

func main() {
	// TODO(aoeu): Add flag values as an anonymous struct
	serve := flag.Bool("preview", false, "Preview colors as a web page")
	port := flag.String("port", ":8081", "The port to serve on")
	numColors := flag.Int("num", 14, "The number of colors")
	saturation := flag.Float64("saturation", 0.9, "Amount of saturation")
	js := flag.Bool("js", false, "Dump colors as a ECMAScript array")
	dump := flag.Bool("dump", false, "Dump preview results to standard output.")
	lightness := flag.Float64("lightness", 0.5, "Amount of lightness")
	flag.Parse()
	s := genHSLColors(*numColors, *saturation, *lightness)
	s = genKellySafe()
	if *js {
		fmt.Println(toECMAScriptArray(s))
	}
	t := template.Must(newHTMLTemplate())
	if *dump {
		t.Execute(os.Stdout, &s)
	}
	if !*serve {
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
