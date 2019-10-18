package codegen

import (
	"strings"

	"github.com/lestrrat-go/slack/internal/stringutil"
)

type Object struct {
	Group    string        `json:"group"`
	SkipList bool          `json:"skip_list"`
	Name     string        `json:"name"`
	Fields   []ObjectField `json:"fields"`
	Typ      string        `json:"type"`
	Validate bool          `json:"validate"`
}

type ObjectField struct {
	AccessorName string `json:"accessor_name,omitempty"`
	Name         string `json:"name"`
	Required     bool   `json:"required"`
	Type         string `json:"type"`
}

func (f *ObjectField) IsSliceType() bool {
	return strings.HasPrefix(f.Type, "[]") ||
		strings.HasSuffix(f.Type, "List")
}

func (f *ObjectField) SliceElementType() string {
	if strings.HasPrefix(f.Type, "[]") {
		return f.Type[2:]
	}

	if f.Type == "BlockList" {
		return "Block"
	}
	return "*" + f.Type[:len(f.Type)-4]
}

func (f *ObjectField) Tag() string {
	tag := stringutil.Snake(f.Name)
	if !f.Required {
		tag += ",omitempty"
	}
	return tag
}

func (f *ObjectField) GoName() string {
	return stringutil.LowerCamel(f.Name)
}

func (f *ObjectField) GoAccessorName() string {
	if f.AccessorName != "" {
		return f.AccessorName
	}
	name := stringutil.Camel(f.Name)
	switch name {
	case "Id":
		name = "ID"
	case "Url":
		name = "URL"
	} // TODO
	return name
}

func (d *Object) GoName() string {
	return stringutil.Camel(d.Name)
}

func (d *Object) Type() string {
	if d.Typ == "" {
		return stringutil.Snake(d.Name[:len(d.Name)-5])
	}
	return d.Typ
}
