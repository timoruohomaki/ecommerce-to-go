package register

import (
	"fmt"
	"html/template"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	pages := []string{
		"internal/page/auth/register.htm",
	}

	tpl := template.Must(template.New("register.htm").ParseFiles(pages...))

	if err := tpl.Execute(w, nil); err != nil {
		fmt.Println(err)
	}
}
