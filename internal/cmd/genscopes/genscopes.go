package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"strconv"
)

func main() {
	if err := _main(); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}

func _main() error {
	f, err := os.Open("oauth-scopes.json")
	if err != nil {
		return err
	}
	defer f.Close()

	var list []struct {
		Name  string `json:"name"`
		Scope string `json:"scope"`
	}
	if err := json.NewDecoder(f).Decode(&list); err != nil {
		return err
	}

	var buf bytes.Buffer
	buf.WriteString("// This file is auto-generated. DO NOT EDIT")
	buf.WriteString("\n\npackage slack")
	buf.WriteString("\n\n// These constants match the scopes provided by Slack API")
	buf.WriteString("\nconst (")
	for _, data := range list {
		fmt.Fprintf(&buf, "%s = %s\n", data.Name, strconv.Quote(data.Scope))
	}
	buf.WriteString("\n)")

	src, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	dst, err := os.Create("scopes_gen.go")
	if err != nil {
		return err
	}
	defer dst.Close()
	dst.Write(src)
	return nil
}
