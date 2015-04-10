package main

import (
	"colors"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templText = `
<html>
	<body>
	{{range .}}
	<div style='width: 50px; height: 50px; display: inline-block; background-color: {{safe .}};'>blah</div>
	{{end}}
	</body>
</html>
`

func main() {
	port := *flag.String("port", ":8081", "The port to serve on")
	numColors := *flag.Int("num", 20, "The number of colors")
	saturation := *flag.Float64("saturation", 1.0, "Amount of saturation")
	lightness := *flag.Float64("lightness", 0.5, "Amount of lightness")
	flag.Parse()
	s := make([]string, numColors)
	for i, h := range colors.NewHSLSet(numColors, saturation, lightness) {
		s[i] = fmt.Sprint(h)
		fmt.Println(h, h.ToRGB(), s[i])
	}
	funcMap := template.FuncMap{
		"safe": func(s string) template.CSS {
			return template.CSS(s)
		},
	}
	t := template.New("colors").Funcs(funcMap)
	t, err := t.Parse(templText)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := t.Execute(w, &s)
		if err != nil {
			log.Fatal(err)
		}
	})
	log.Fatal(http.ListenAndServe(port, nil))
}
