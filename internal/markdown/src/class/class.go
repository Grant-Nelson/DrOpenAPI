package class

import (
	"fmt"

	"github.com/grant-nelson/DrOpenAPI/internal/markdown"
)

type classImp struct {
	name   string
	body   markdown.StringBuffer
	hasCon map[string]bool
	cons   markdown.StringBuffer
}

func New(factory markdown.Factory, name string) markdown.Class {
	return &classImp{
		name:   name,
		body:   factory.StringBuffer(),
		hasCon: map[string]bool{},
		cons:   factory.StringBuffer(),
	}
}

func (imp *classImp) AddEntry(entry string) {
	imp.body.Write("  %s\n", entry)
}

func (imp *classImp) AddMember(name, typeName string) {
	imp.body.Write("  %s %s\n", typeName, name)
}

func (imp *classImp) ConnectTo(name string) {
	if !imp.hasCon[name] {
		imp.hasCon[name] = true
		imp.cons.Write("%s --> %s\n", imp.name, name)
	}
}

func (imp *classImp) String() string {
	return fmt.Sprintf("class %s {\n%s}\n%s",
		imp.name, imp.body.String(), imp.cons.String())
}