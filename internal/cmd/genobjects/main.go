package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/lestrrat-go/slack/internal/codegen"
	"github.com/lestrrat-go/slack/internal/stringutil"
	"github.com/pkg/errors"
)

func main() {
	if err := _main(); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}

func _main() error {
	f, err := os.Open("objects.json")
	if err != nil {
		return err
	}
	defer f.Close()

	var list []codegen.Object
	if err := json.NewDecoder(f).Decode(&list); err != nil {
		return err
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Name < list[j].Name
	})

	if err := writeInterface(list); err != nil {
		return errors.Wrap(err, `failed to write interfaces file`)
	}
	if err := writeBlock(list); err != nil {
		return errors.Wrap(err, `failed to write block file`)
	}
	if err := writeObjects(list); err != nil {
		return errors.Wrap(err, `failed to write objects file`)
	}
	if err := writeResponses(list); err != nil {
		return errors.Wrap(err, `failed to write responses file`)
	}
	if err := writeLists(list); err != nil {
		return errors.Wrap(err, `failed to write lists file`)
	}

	return nil
}

func writeInterface(list []codegen.Object) error {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "\npackage objects")

	fmt.Fprintf(&buf, "\n\ntype TextType string")
	fmt.Fprintf(&buf, "\n\nconst (")
	fmt.Fprintf(&buf, "\nMarkdownTextType = %s", strconv.Quote("mrkdwn"))
	fmt.Fprintf(&buf, "\nPlainTextType = %s", strconv.Quote("plain_text"))
	fmt.Fprintf(&buf, "\n)")

	fmt.Fprintf(&buf, "\n\ntype BlockType string")
	fmt.Fprintf(&buf, "\n\nconst (")
	for _, def := range list {
		if !strings.HasSuffix(def.Name, "Block") {
			continue
		}
		fmt.Fprintf(&buf, "\n%sType BlockType = %s", def.Name, strconv.Quote(stringutil.Snake(def.Name[:len(def.Name)-5])))
	}
	fmt.Fprintf(&buf, "\n)")

	fmt.Fprintf(&buf, "\n\ntype Block interface {")
	fmt.Fprintf(&buf, "\nType() BlockType")
	fmt.Fprintf(&buf, "\n}")
	fmt.Fprintf(&buf, "\ntype BlockList []Block")

	fmt.Fprintf(&buf, "\n\ntype ElementType string")
	fmt.Fprintf(&buf, "\n\nconst (")
	for _, elementType := range []string{"Image", "Button", "SelectMenu", "MultiSelectMenu", "OverflowMenu", "DatePicker", "Input"} {
		fmt.Fprintf(&buf, "\n%sElementType ElementType = %s", elementType, strconv.Quote(stringutil.Snake(elementType)))
	}
	fmt.Fprintf(&buf, "\n)")

	fmt.Fprintf(&buf, "\n\ntype BlockElement interface {")
	fmt.Fprintf(&buf, "\nType() ElementType")
	fmt.Fprintf(&buf, "\n}")

	for _, def := range list {
		fmt.Fprintf(&buf, "\n\ntype %s struct {", def.Name)
		for _, field := range def.Fields {
			fmt.Fprintf(&buf, "\n\t%s %s `json:%s`", field.GoName(), field.Type, strconv.Quote(field.Tag()))
		}
		fmt.Fprintf(&buf, "\n}")
		fmt.Fprintf(&buf, "\n\ntype %sBuilder struct {", def.Name)
		for _, field := range def.Fields {
			fmt.Fprintf(&buf, "\n\t%s %s", field.GoName(), field.Type)
		}
		fmt.Fprintf(&buf, "\n}")

		if def.Name == "ReactionsGetResponse" || (def.Group == "Object" && !def.SkipList) {
			fmt.Fprintf(&buf, "\n\ntype %[1]sList []*%[1]s", def.Name)
		}
	}

	return codegen.WriteGoCodeToFile("objects/interface_gen.go", buf.Bytes())
	return nil
}

func writeBuilder(dst io.Writer, def codegen.Object) error {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "\n\nfunc Build%s(", stringutil.Camel(def.Name))
	var requireds []codegen.ObjectField
	var optionals []codegen.ObjectField
	for _, field := range def.Fields {
		if !field.Required {
			optionals = append(optionals, field)
			continue
		}
		requireds = append(requireds, field)
		if field.IsSliceType() {
			fmt.Fprintf(&buf, "%s ...%s, ", field.GoName(), field.SliceElementType())
		} else {
			fmt.Fprintf(&buf, "%s %s, ", field.GoName(), field.Type)
		}
	}
	if len(requireds) > 0 {
		buf.Truncate(buf.Len() - 2)
	}
	fmt.Fprintf(&buf, ") *%sBuilder {", stringutil.Camel(def.Name))
	fmt.Fprintf(&buf, "\nvar b %sBuilder", stringutil.Camel(def.Name))
	for _, field := range requireds {
		fmt.Fprintf(&buf, "\nb.%s = %s", field.GoName(), field.GoName())
	}
	fmt.Fprintf(&buf, "\nreturn &b")
	fmt.Fprintf(&buf, "\n}")

	for _, field := range optionals {
		fmt.Fprintf(&buf, "\n\nfunc (b *%sBuilder) %s(v ", def.GoName(), field.GoAccessorName())
		if field.IsSliceType() {
			fmt.Fprintf(&buf, "...%s", field.SliceElementType())
		} else {
			fmt.Fprintf(&buf, "%s", field.Type)
		}
		fmt.Fprintf(&buf, ") *%sBuilder {", def.GoName())
		//			if field.IsSliceType() {
		//				fmt.Fprintf(&buf, "\nb.%s = %s(v)", field.GoName(), field.Type)
		//			} else {
		fmt.Fprintf(&buf, "\nb.%s = v", field.GoName())
		//			}
		fmt.Fprintf(&buf, "\nreturn b")
		fmt.Fprintf(&buf, "\n}")
	}

	fmt.Fprintf(&buf, "\n\nfunc (b *%[1]sBuilder) Build() (*%[1]s, error) {", def.GoName())
	if def.Validate {
		fmt.Fprintf(&buf, "\nif err := b.Validate(); err != nil {")
		fmt.Fprintf(&buf, "\nreturn nil, errors.Wrap(err, `validation for %s failed`)", def.GoName())
		fmt.Fprintf(&buf, "\n}")
	}

	fmt.Fprintf(&buf, "\nvar v %s", def.GoName())
	for _, field := range def.Fields {
		fmt.Fprintf(&buf, "\nv.%s = b.%s", field.GoName(), field.GoName())
	}
	fmt.Fprintf(&buf, "\nreturn &v, nil")
	fmt.Fprintf(&buf, "\n}")

	fmt.Fprintf(&buf, "\n\nfunc (b *%[1]sBuilder) MustBuild() *%[1]s {", def.GoName())
	fmt.Fprintf(&buf, "\nv, err := b.Build()")
	fmt.Fprintf(&buf, "\nif err != nil {")
	fmt.Fprintf(&buf, "\npanic("+`"error during %s.MustBuild: " + err.Error())`, def.GoName())
	fmt.Fprintf(&buf, "\n}")
	fmt.Fprintf(&buf, "\nreturn v")
	fmt.Fprintf(&buf, "\n}")

	buf.WriteTo(dst)
	return nil
}

func writeBlock(list []codegen.Object) error {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "\npackage objects")
	for _, def := range list {
		if !strings.HasSuffix(def.Name, "Block") {
			continue
		}

		writeBuilder(&buf, def)

		for _, field := range def.Fields {
			writeAccessor(&buf, def, field)
		}

		fmt.Fprintf(&buf, "\n\nfunc (b %s) Type() BlockType {", def.GoName())
		fmt.Fprintf(&buf, "\n\treturn %sType", def.GoName())
		fmt.Fprintf(&buf, "\n}")
	}

	return codegen.WriteGoCodeToFile("objects/blocks_gen.go", buf.Bytes())
}

func writeObjects(list []codegen.Object) error {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "\npackage objects")

	fmt.Fprintf(&buf, "\n\nimport (")
	fmt.Fprintf(&buf, "\n\n%s", strconv.Quote("github.com/pkg/errors"))
	fmt.Fprintf(&buf, "\n)")

	for _, def := range list {
		if def.Group != "Object" {
			continue
		}

		writeBuilder(&buf, def)
		for _, field := range def.Fields {
			writeAccessor(&buf, def, field)
		}

	}

	return codegen.WriteGoCodeToFile("objects/objects_gen.go", buf.Bytes())
}

func writeAccessor(dst io.Writer, def codegen.Object, field codegen.ObjectField) error {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "\n\nfunc(b *%s) %s() %s {", def.GoName(), field.GoAccessorName(), field.Type)
	fmt.Fprintf(&buf, "\nreturn b.%s", field.GoName())
	fmt.Fprintf(&buf, "\n}")

	buf.WriteTo(dst)
	return nil
}

func writeLists(list []codegen.Object) error {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "\npackage objects")

	writeAppend := func(dst io.Writer, sliceTyp, typ string) {
		fmt.Fprintf(&buf, "\n\nfunc (l %[1]s) Append(v %[2]s) %[1]s {", sliceTyp, typ)
		fmt.Fprintf(&buf, "\n*l = append(*l, v)")
		fmt.Fprintf(&buf, "\nreturn l")
		fmt.Fprintf(&buf, "\n}")
	}

	writeAppend(&buf, "*BlockList", "Block")
	for _, def := range list {
		if def.Name != "ReactionsGetResponse" && (def.Group != "Object" || def.SkipList) {
			continue
		}

		writeAppend(&buf, "*"+def.Name+"List", "*"+def.Name)
	}

	return codegen.WriteGoCodeToFile("objects/lists_gen.go", buf.Bytes())
}

func writeResponses(list []codegen.Object) error {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "\npackage objects")

	fmt.Fprintf(&buf, "\n\nimport (")
	for _, pkg := range []string{"encoding/json", "github.com/pkg/errors"} {
		fmt.Fprintf(&buf, "\n%s", strconv.Quote(pkg))
	}
	fmt.Fprintf(&buf, "\n)")

	for _, def := range list {
		if !strings.HasSuffix(def.Name, "Response") {
			continue
		}

		writeBuilder(&buf, def)
		for _, field := range def.Fields {
			writeAccessor(&buf, def, field)
		}

		fmt.Fprintf(&buf, "\n\nfunc (v *%s) UnmarshalJSON(data []byte) error {", def.Name)
		fmt.Fprintf(&buf, "\nvar proxy struct {")
		for _, field := range def.Fields {
			fmt.Fprintf(&buf, "\n%s %s `json:%s`", field.GoAccessorName(), field.Type, strconv.Quote(field.Name))
		}
		fmt.Fprintf(&buf, "\n}")
		fmt.Fprintf(&buf, "\nif err := json.Unmarshal(data, &proxy); err != nil {")
		fmt.Fprintf(&buf, "\nreturn errors.Wrap(err, `failed to unmarshal JSON`)")
		fmt.Fprintf(&buf, "\n}")
		fmt.Fprintf(&buf, "\n\nx, err := Build%s().", def.Name)
		for _, field := range def.Fields {
			var flatten string
			if field.IsSliceType() {
				flatten = "..."
			}
			fmt.Fprintf(&buf, "\n%[1]s(proxy.%[1]s%[2]s).", field.GoAccessorName(), flatten)
		}
		fmt.Fprintf(&buf, "\nBuild()")
		fmt.Fprintf(&buf, "\nif err != nil {")
		fmt.Fprintf(&buf, "\nreturn errors.Wrap(err, `failed to build object from JSON`)")
		fmt.Fprintf(&buf, "\n}")
		fmt.Fprintf(&buf, "\n*v = *x")
		fmt.Fprintf(&buf, "\nreturn nil")
		fmt.Fprintf(&buf, "\n}")
	}

	return codegen.WriteGoCodeToFile("objects/responses_gen.go", buf.Bytes())
}