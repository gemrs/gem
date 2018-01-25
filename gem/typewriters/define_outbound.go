package typewriters

import (
	"fmt"
	"html/template"
	"io"

	"github.com/clipperhouse/typewriter"
)

func init() {
	err := typewriter.Register(&DefineOutbound{})
	if err != nil {
		panic(err)
	}
}

type DefineOutbound struct {
	packetTypes        []string
	unimplementedTypes []string
	bareTypes          []string
}

func (p DefineOutbound) Name() string {
	return "define_outbound"
}

func (p DefineOutbound) Imports(t typewriter.Type) []typewriter.ImportSpec {
	return nil
}

var outboundDefineTmpl = template.Must(template.New("").Parse(`
var {{.Name}}Definition = OutboundPacketDefinition{
	Number: {{.Number}},
	Size:   {{.Size}},
}
`))

func (p DefineOutbound) Write(w io.Writer, t typewriter.Type) error {
	tags, ok := t.FindTag(p)
	if !ok {
		return nil
	}

	values := tags.Values

	if len(values) != 2 {
		return fmt.Errorf("format: Pkt123,Size")
	}

	outboundDefineTmpl.Execute(w, struct {
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

func (p *DefineOutbound) Visit(t typewriter.Type) error {
	tags, ok := t.FindTag(p)
	if !ok {
		return nil
	}

	if len(tags.Values) > 0 {
		if tags.Values[0].Name == "Unimplemented" {
			p.unimplementedTypes = append(p.unimplementedTypes, t.Name)
		} else {
			p.packetTypes = append(p.packetTypes, t.Name)
		}
	} else {
		p.bareTypes = append(p.bareTypes, t.Name)
	}
	return nil
}

var encodeTmpl = template.Must(template.New("").Parse(`
func (protocolImpl) Encode(message server.Message) encoding.Encodable {
	switch message := message.(type) {
{{range .PacketTypes}}
	case protocol.{{.}}:
		return {{.}}Definition.Pack({{.}}(message))
{{end}}
{{range .BareTypes}}
	case protocol.{{.}}:
		return {{.}}(message)
{{end}}
{{range .UnimplementedTypes}}
	case protocol.{{.}}:
        fmt.Println("{{.}} not implemented")
		return nil
{{end}}
	}
	panic(fmt.Sprintf("cannot encode %T", message))
}
`))

func (p DefineOutbound) Collect(w io.Writer) error {
	if len(p.packetTypes)+len(p.bareTypes) > 0 {
		return encodeTmpl.Execute(w, struct {
			PacketTypes        []string
			BareTypes          []string
			UnimplementedTypes []string
		}{
			PacketTypes:        p.packetTypes,
			BareTypes:          p.bareTypes,
			UnimplementedTypes: p.unimplementedTypes,
		})
	}
	return nil
}
