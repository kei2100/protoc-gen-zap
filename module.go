package main

import (
	"fmt"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	zappb "github.com/shankulkarni/protoc-gen-zap/zap"
	"strings"
)

type zapGen struct {
	pgs.ModuleBase
	pgsgo.Context
}

func (*zapGen) Name() string {
	return "zap"
}

func (m *zapGen) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.Context = pgsgo.InitContext(c.Parameters())
}

func (m *zapGen) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {

	for _, f := range targets {

		name := m.Context.OutputPath(f).SetExt(".zap.go").String()
		fm := fileModel{PackageName: m.Context.PackageName(f).String()}

		for _, msg := range f.AllMessages() {

			fields := msg.Fields()
			mp := messageModel{}

			// .<package>.<ParentMessage>.<ChildMessage>
			fqName := msg.FullyQualifiedName()
			messageName := strings.ReplaceAll(strings.TrimPrefix(fqName, fmt.Sprintf(".%s.", msg.Package().ProtoName().String())), ".", "_")
			mp.Name = messageName

			list := make([]zapField, len(f.AllMessages()))

			for _, v := range fields {

				redact := false
				_, err := v.Extension(zappb.E_Redact, &redact)
				if err != nil {
					m.Log(err)
				}

				var accessor string
				if v.InOneOf() {
					accessor = fmt.Sprintf("Get%s()", v.Name().UpperCamelCase().String())
				} else {
					accessor = v.Name().UpperCamelCase().String()
				}
				r := zapField{
					Redact:      redact,
					Name:        v.Name().UpperCamelCase().String(),
					Accessor:    accessor,
					Type:        v.Descriptor().Type.String(),
					Label:       v.Descriptor().GetLabel().String(),
					TypeName:    v.Descriptor().GetTypeName(),
					FType:       v.Type(),
					SamePackage: false,
				}

				if strings.HasPrefix(r.TypeName, "."+fm.PackageName+".") {
					r.SamePackage = true
				}

				list = append(list, r)

			}
			mp.Fields = list
			fm.Messages = append(fm.Messages, mp)
		}

		m.OverwriteGeneratorTemplateFile(
			name,
			T.Lookup("File"),
			&fm,
		)
	}

	return m.Artifacts()

}

type zapField struct {
	Name        string
	Accessor    string
	Type        string
	Redact      bool
	Label       string
	TypeName    string
	FType       pgs.FieldType
	SamePackage bool
}

type messageModel struct {
	Name   string
	Fields []zapField
}

type fileModel struct {
	PackageName string
	Messages    []messageModel
}
