package main

import (
    "text/template"
    "log"
    "bytes"
    "fmt"
)

func main() {
    data := make(map[string]string)
    //data["AppVersion"] = "Octane_3.0"

    text := "APP_VERSION={{.AppVersion}}"
    tmpl, err := template.New("").Option("missingkey=default").Parse(text)
    if err != nil {
        log.Fatal(err)
    }

    var b bytes.Buffer
    err = tmpl.Execute(&b, data)
    if err != nil {
        fmt.Println("template.Execute failed", err)
    }

    fmt.Println("Template text:", text)
    fmt.Println("Expanded:", b.String())
}
