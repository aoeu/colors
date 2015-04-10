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

func preview(c []colors.HSL) {
	// TODO(aoeu): Implement browser preview here.
}


/*

TODO(aoeu): Something like this is what gets jammed into d3.
Dump it on a command line? 
Sprintf out the whole line of js? (No.)
var color = d3.scale.ordinal()
    .range(["#98abc5", "#8a89a6", "#7b6888", "#6b486b", 
            "#a05d56", "#d0743c", "#ff8c00"]);

*/

func main() {
	port := *flag.String("port", ":8081", "The port to serve on")
	numColors := *flag.Int("num", 20, "The number of colors")
	saturation := *flag.Float64("saturation", 1.0, "Amount of saturation")
	lightness := *flag.Float64("lightness", 0.5, "Amount of lightness")
	flag.Parse()
	s := make([]string, numColors)
	for i, h := range colors.NewHSLSet(numColors, saturation, lightness) {
		s[i] = fmt.Sprint(h)
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

