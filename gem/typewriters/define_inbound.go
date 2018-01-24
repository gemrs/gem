package typewriters

import (
	"fmt"
	"html/template"
	"io"

	"github.com/clipperhouse/typewriter"
)

func init() {
	err := typewriter.Register(&DefineInbound{})
	if err != nil {
		panic(err)
	}
}

type InboundTemplate struct {
	Name      string
	ProtoName string
	Number    string
	Size      string
}

type DefineInbound struct {
	packets []InboundTemplate
}

func (p DefineInbound) Name() string {
	return "define_inbound"
}

func (p DefineInbound) Imports(t typewriter.Type) []typewriter.ImportSpec {
	return nil
}

var inboundDefineTmpl = template.Must(template.New("").Parse(`
var {{.Name}}Definition = InboundPacketDefinition{
	Number: {{.Number}},
	Size:   encoding.{{.Size}},
}
`))

func (p DefineInbound) Write(w io.Writer, t typewriter.Type) error {
	tags, ok := t.FindTag(p)
	if !ok {
		return nil
	}

	values := tags.Values

	if len(values) < 2 {
		return fmt.Errorf("format: Pkt123,Size")
	}

	inboundDefineTmpl.Execute(w, struct {
		Name   string
		Number string
		Size   string
	}{
		Name:   t.Name,
		Number: values[0].Name[3:],
		Size:   values[1].Name,
	})
	return nil
}

func (p *DefineInbound) Visit(t typewriter.Type) error {
	tags, ok := t.FindTag(p)
	if !ok {
		return nil
	}

	values := tags.Values

	if len(values) < 2 {
		return fmt.Errorf("format: Pkt123,Size,[protoname]")
	}

	protoName := t.Name
	if len(values) == 3 {
		protoName = values[2].Name
	}

	p.packets = append(p.packets, InboundTemplate{
		Name:      t.Name,
		ProtoName: protoName,
		Number:    values[0].Name[3:],
		Size:      values[1].Name,
	})
	return nil
}

var inboundBuildTmpl = template.Must(template.New("").Parse(`
var inboundPacketBuilders = map[int]func() encoding.Decodable{
{{range .}}
	{{.Number}}: func() encoding.Decodable {
		return &encoding.PacketHeader{
			Number: {{.Name}}Definition.Number,
			Size: {{.Name}}Definition.Size,
			Object: new({{.Name}}),
		}
	},
{{end}}
}

func (p protocolImpl) Decode(message encoding.Decodable) server.Message {
	switch message := message.(type) {
{{range .}}
	case *{{.Name}}:
		return (*protocol.{{.ProtoName}})(message)
{{end}}

	case *UnknownPacket:
		return (*protocol.UnknownPacket)(message)

	case *encoding.PacketHeader:
		return p.Decode(message.Object.(encoding.Decodable))
	}
	panic(fmt.Sprintf("cannot decode %T", message))
}
`))

func (p DefineInbound) Collect(w io.Writer) error {
	if len(p.packets) > 0 {
		return inboundBuildTmpl.Execute(w, p.packets)
	}
	return nil
}
