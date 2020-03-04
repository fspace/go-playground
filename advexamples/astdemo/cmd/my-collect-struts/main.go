package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"os"

	"unicode"
)

// structType contains a structType node and it's name. It's a convenient
// helper type, because *ast.StructType doesn't contain the name of the struct
type structType struct {
	name string
	node *ast.StructType
}

// config defines how tags should be modified
type config struct {
	file   string
	output string
	write  bool
	fset   *token.FileSet

	remove        []string
	removeOptions []string

	transform   string
	sort        bool
	clear       bool
	clearOption bool
}

func main() {
	if err := realMain(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func realMain() error {
	var (
		// file flags
		flagFile = flag.String("file", "", "Filename to be parsed")
		//flagWrite = flag.Bool("w", false,
		//	"Write result to (source) file instead of stdout")
		//flagOutput = flag.String("format", "source", "Output format."+
		//	"By default it's the whole file. Options: [source, json]")

		// processing modes
		// flagStruct = flag.String("struct", "", "Struct name to be processed")

		// tag flags
		//flagRemoveTags = flag.String("remove-tags", "",
		//	"Remove tags for the comma separated list of keys")

		// option flags
		//flagRemoveOptions = flag.String("remove-options", "",
		//	"Remove the comma separated list of options from the given keys, "+
		//		"i.e: json=omitempty,hcl=squash")
		//flagClearOptions = flag.Bool("clear-options", false,
		//	"Clear all tag options")
		//flagAddOptions = flag.String("add-options", "",
		//	"Add the options per given key. i.e: json=omitempty,hcl=squash")
	)

	// don't output full help information if something goes wrong
	flag.Usage = func() {}
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		return nil
	}

	cfg := &config{
		file: *flagFile,
	}
	//if *flagRemoveOptions != "" {
	//	cfg.removeOptions = strings.Split(*flagRemoveOptions, ",")
	//}

	err := cfg.validate()
	if err != nil {
		return err
	}

	node, err := cfg.parse()
	if err != nil {
		return err
	}
	structs := collectStructs(node)
	structNames := []string{}
	for _, s := range structs {
		structNames = append(structNames, s.name)
	}
	// pretty.Print(structNames)
	data, err := json.Marshal(structNames)
	fmt.Println(string(data))
	//out, err := cfg.format(rewrittenNode, errs)
	//if err != nil {
	//	return err
	//}
	//fmt.Println(out)
	return nil
}

// collectStructs collects and maps structType nodes to their positions
func collectStructs(node ast.Node) map[token.Pos]*structType {
	structs := make(map[token.Pos]*structType, 0)
	collectStructs := func(n ast.Node) bool {
		var t ast.Expr
		var structName string

		switch x := n.(type) {
		case *ast.TypeSpec:
			if x.Type == nil {
				return true

			}

			structName = x.Name.Name
			t = x.Type
		case *ast.CompositeLit:
			t = x.Type
		case *ast.ValueSpec:
			structName = x.Names[0].Name
			t = x.Type
		}

		x, ok := t.(*ast.StructType)
		if !ok {
			return true
		}

		structs[x.Pos()] = &structType{
			name: structName,
			node: x,
		}
		return true
	}
	ast.Inspect(node, collectStructs)
	return structs
}

func (c *config) parse() (ast.Node, error) {
	c.fset = token.NewFileSet()
	var contents interface{}
	//if c.modified != nil {
	//	//	archive, err := buildutil.ParseOverlayArchive(c.modified)
	//	//	if err != nil {
	//	//		return nil, fmt.Errorf("failed to parse -modified archive: %v", err)
	//	//	}
	//	//	fc, ok := archive[c.file]
	//	//	if !ok {
	//	//		return nil, fmt.Errorf("couldn't find %s in archive", c.file)
	//	//	}
	//	//	contents = fc
	//	//}

	return parser.ParseFile(c.fset, c.file, contents, parser.ParseComments)
}

//func (c *config) format(file ast.Node, rwErrs error) (string, error) {
//	switch c.output {
//	case "source":
//		var buf bytes.Buffer
//		err := format.Node(&buf, c.fset, file)
//		if err != nil {
//			return "", err
//		}
//
//		if c.write {
//			err = ioutil.WriteFile(c.file, buf.Bytes(), 0)
//			if err != nil {
//				return "", err
//			}
//		}
//
//		return buf.String(), nil
//	case "json":
//		// NOTE(arslan): print first the whole file and then cut out our
//		// selection. The reason we don't directly print the struct is that the
//		// printer is not capable of printing loosy comments, comments that are
//		// not part of any field inside a struct. Those are part of *ast.File
//		// and only printed inside a struct if we print the whole file. This
//		// approach is the sanest and simplest way to get a struct printed
//		// back. Second, our cursor might intersect two different structs with
//		// other declarations in between them. Printing the file and cutting
//		// the selection is the easier and simpler to do.
//		var buf bytes.Buffer
//
//		// this is the default config from `format.Node()`, but we add
//		// `printer.SourcePos` to get the original source position of the
//		// modified lines
//		cfg := printer.Config{Mode: printer.SourcePos | printer.UseSpaces | printer.TabIndent, Tabwidth: 8}
//		err := cfg.Fprint(&buf, c.fset, file)
//		if err != nil {
//			return "", err
//		}
//
//		lines, err := parseLines(&buf)
//		if err != nil {
//			return "", err
//		}
//
//		// prevent selection to be larger than the actual number of lines
//		if c.start > len(lines) || c.end > len(lines) {
//			return "", errors.New("line selection is invalid")
//		}
//
//		out := &output{
//			Start: c.start,
//			End:   c.end,
//			Lines: lines[c.start-1 : c.end],
//		}
//
//		if rwErrs != nil {
//			if r, ok := rwErrs.(*rewriteErrors); ok {
//				for _, err := range r.errs {
//					out.Errors = append(out.Errors, err.Error())
//				}
//			}
//		}
//
//		o, err := json.MarshalIndent(out, "", "  ")
//		if err != nil {
//			return "", err
//		}
//
//		return string(o), nil
//	default:
//		return "", fmt.Errorf("unknown output mode: %s", c.output)
//	}
//}

func isPublicName(name string) bool {
	for _, c := range name {
		return unicode.IsUpper(c)
	}
	return false
}

// validate validates whether the config is valid or not
func (c *config) validate() error {
	if c.file == "" {
		return errors.New("no file is passed")
	}
	//if (c.add == nil || len(c.add) == 0) &&
	//	(c.addOptions == nil || len(c.addOptions) == 0) &&
	//	!c.clear &&
	//	!c.clearOption &&
	//	(c.removeOptions == nil || len(c.removeOptions) == 0) &&
	//	(c.remove == nil || len(c.remove) == 0) {
	//	return errors.New("one of " +
	//		"[-add-tags, -add-options, -remove-tags, -remove-options, -clear-tags, -clear-options]" +
	//		" should be defined")
	//}

	return nil
}
