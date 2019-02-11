package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/oshankkumar/preetyjson/version"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const desc = `    
    preetyjson is a tool for processing JSON inputs, applying the
    given filter to its JSON text inputs and producing the
    filter's results as JSON on standard output.`

const example = `
    # format json input from a file source
    preetyjson -f inputFile.json 
 
    # format json input and use tabs for indentation
    preetyjson -f inputFile.json -i

    # format json input in bold format
    preetyjson -f inputFile.json -F bold
   
    # format json input in bold and italic format
    preetyjson -f inputFile.json -F bold,italic		

    # format based on the JSON passed into stdin.
    cat inputFile.json | preetyjson
    curl -XGET http://ip:port/url/path | preetyjson
`

func NewPreetyJsonCmd() *cobra.Command {
	opts := &PreetyJsonRunOptions{}
	cmd := &cobra.Command{
		Use:     "preetyjson",
		Short:   "",
		Long:    desc,
		RunE:    opts.RunPreety,
		Example: example,
		Version: version.Version(),
	}

	cmd.Flags().BoolVarP(&opts.compact, "compact", "c", false, "compact instead of pretty-printed output")
	cmd.Flags().StringVarP(&opts.fileName, "file", "f", "", "name of file use to process")
	cmd.Flags().BoolVar(&opts.colorize, "colorize", false, "colorize JSON")
	cmd.Flags().BoolVarP(&opts.indent, "indent", "i", false, "use tabs for indentation")
	cmd.Flags().StringSliceVarP(&opts.formats, "format", "F", nil, "output print formats [ bold|italic|faint|blink ]")

	return cmd
}

type PreetyJsonRunOptions struct {
	compact  bool
	fileName string
	colorize bool
	indent   bool
	bold     bool
	formats  []string
}

func (p *PreetyJsonRunOptions) GetAttr() []color.Attribute {
	var attrs []color.Attribute

	for _, format := range p.formats {
		switch format {
		case "bold":
			attrs = append(attrs, color.Bold)
		case "italic":
			attrs = append(attrs, color.Italic)
		case "faint":
			attrs = append(attrs, color.Faint)
		case "blink":
			attrs = append(attrs, color.BlinkSlow)
		}
	}

	if p.colorize {
		attrs = append(attrs, color.FgCyan)
	}

	return attrs
}

func (p *PreetyJsonRunOptions) InputBody(cmd *cobra.Command, args []string) (body []byte, err error) {
	reader := os.Stdin
	if p.fileName != "" {
		if reader, err = os.Open(p.fileName); err != nil {
			return
		}
	}

	return ioutil.ReadAll(reader)
}

func (p *PreetyJsonRunOptions) RunPreety(cmd *cobra.Command, args []string) error {

	body, err := p.InputBody(cmd, args)
	if err != nil {
		return err
	}

	attrs := p.GetAttr()
	writer := color.New(attrs...)

	var contents json.RawMessage

	if err := json.Unmarshal(body, &contents); err != nil {
		return err
	}

	if p.compact {
		return json.NewEncoder(WriterFunc(func(b []byte) (int, error) {
			return writer.Println(string(b))
		})).Encode(contents)
	}

	indent := "  "
	if p.indent {
		indent = "\t"
	}

	body, err = json.MarshalIndent(contents, "", indent)
	if err != nil {
		return err
	}

	_, err = writer.Println(string(body))

	return err
}
