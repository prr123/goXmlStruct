package xmlStruct

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type YamlInp struct {
	Config struct {
//    attrNameSuffix               = pflag.String("attr-name-suffix", xmlstruct.DefaultAttrNameSuffix, "attribute name suffix")
    AttrNameSuffix	string `yaml:"AttrNameSuffix"` // attribute name suffix
//    charDataFieldName            = pflag.String("char-data-field-name", xmlstruct.DefaultCharDataFieldName, "char data field name")
    CharDataFieldName string `yaml:"CharDataFieldName"` // char data field name
//    compactTypes                 = pflag.Bool("compact-types", xmlstruct.DefaultCompactTypes, "create compact types")
    CompactTypes  bool `yaml:"CompactTypes"`  // create compact types
//    elemNameSuffix               = pflag.String("elem-name-suffix", xmlstruct.DefaultElemNameSuffix, "element name suffix")
    ElemNameSuffix string  `yaml:"ElemNameSuffix"` // element name suffix
//    formatSource                 = pflag.Bool("format-source", xmlstruct.DefaultFormatSource, "format source")
    FormatSource bool `yaml:"FormatSource"` // format source
//    header                       = pflag.String("header", xmlstruct.DefaultHeader, "header")
    Header string `yaml:"Header"` // header
//    ignoreErrors                 = pflag.Bool("ignore-errors", false, "ignore errors")
    IgnoreErrors bool  `yaml:"IgnoreErrors"` // ignore errors
//    ignoreNamespaces             = pflag.Bool("ignore-namespaces", true, "ignore namespaces")
    IgnoreNamespaces bool `yaml:"IgnoreNamespaces"` // ignore namespaces
//    imports                      = pflag.Bool("imports", xmlstruct.DefaultImports, "generate import statements")
    Imports bool `yaml:"Imports"` // generate import statements
//    intType                      = pflag.String("int-type", xmlstruct.DefaultIntType, "int type")
    IntType string `yaml:"IntType"` //int type
//    namedRoot                    = pflag.Bool("named-root", xmlstruct.DefaultNamedRoot, "create an XMLName field for the root element")
    NamedRoot bool `yaml:"NamedRoot"` // create an XMLName field for the root element
//    namedTypes                   = pflag.Bool("named-types", xmlstruct.DefaultNamedTypes, "create named types for all elements")
    NamedTypes bool `yaml:"NamedTypes"` // create named types for all elements
//    noEmptyElements              = pflag.Bool("no-empty-elements", !xmlstruct.DefaultEmptyElements, "use type string instead of struct{} for empty ele>
    NoEmptyElements bool `yaml:"NoEmptyElements"` // use type string instead of struct{} for empty elements
//    noExport                     = pflag.Bool("no-export", false, "create unexported types")
    NoExport bool `yaml:"NoExport"` // create unexported types
//    output                       = pflag.String("output", "", "output filename")
    Output string `yaml:"Output"` // output filename
//    packageName                  = pflag.String("package-name", "main", "package name")
    PackageName string `yaml:"PackageName"` // package name
//    pattern                      = pflag.String("pattern", "", "filename pattern to observe")
    Pattern string `yaml:"Pattern"` // filename pattern to observe
//    preserveOrder                = pflag.Bool("preserve-order", xmlstruct.DefaultPreserveOrder, "preserve order of types and fields")
    PreserveOrder bool `yaml:"PreserveOrder"` // preserve order of types and fields
//    timeLayout                   = pflag.String("time-layout", "2006-01-02T15:04:05Z", "time layout")
    TimeLayout string `yaml:"TimeLayout"` // template for time layout
//    topLevelAttributes           = pflag.Bool("top-level-attributes", xmlstruct.DefaultTopLevelAttributes, "include top level attributes")
    TopLevelAttributes bool `yaml:"TopLevelAttributes"` // include top level attributes
//    typesOnly                    = pflag.Bool("types-only", false, "generate structs only, without header, package, or imports")
    TypesOnly bool `yaml:"TypesOnly"` // generate structs only, without header, package, or imports
//    usePointersForOptionalFields = pflag.Bool("use-pointers-for-optional-fields", xmlstruct.DefaultUsePointersForOptionalFields, "use pointers for opt>
    UsePointersForOptionalFields bool `yaml:"UsePointersForOptionalFields"` // use pointers for optional fields
//    useRawToken                  = pflag.Bool("use-raw-token", xmlstruct.DefaultUseRawToken, "use encoding/xml.Decoder.RawToken")
    UseRawToken  bool `yaml:"UseRawToken"`// use encoding/xml.Decoder.RawToken
} `yaml:"Config"`
}

func RdYaml(filnam string) (config YamlInp, err error) {

	yamlData, err := os.ReadFile(filnam)
	if err != nil {return config, fmt.Errorf("cannot read yaml file [%s]: %v", filnam, err)}

	err = yaml.Unmarshal([]byte(yamlData), &config)
	if err != nil {return config, fmt.Errorf("cannot unmarshal yamldata: %v", err)}

	return config, nil
}

func PrintConfig(cfg YamlInp) {

	config := cfg.Config
	fmt.Println("****** Config *******")
    fmt.Printf("  AttrNameSuffix:    %s\n", config.AttrNameSuffix)	// attribute name suffix
    fmt.Printf("  CharDataFieldName: %s\n", config.CharDataFieldName) // char data field name
    fmt.Printf("  CompactTypes:      %t\n", config.CompactTypes) // create compact types
    fmt.Printf("  ElemNameSuffix:    %s\n", config.ElemNameSuffix) // element name suffix
    fmt.Printf("  FormatSource:      %t\n", config.FormatSource) // format source
    fmt.Printf("  Header:            %s\n", config.Header) // header

    fmt.Printf("  IgnoreErrors:       %t\n", config.IgnoreErrors) // ignore errors
    fmt.Printf("  IgnoreNamespaces:   %t\n", config.IgnoreNamespaces) // ignore namespaces
    fmt.Printf("  Imports:            %t\n", config.Imports) // generate import statements

    fmt.Printf("  IntType:            %s\n", config.IntType) //int type
    fmt.Printf("  NamedRoot:          %t\n", config.NamedRoot) // create an XMLName field for the root element
    fmt.Printf("  NamedTypes:         %t\n", config.NamedTypes) // create named types for all elements

    fmt.Printf("  NoEmptyElements:    %t\n", config.NoEmptyElements) // use type string instead of struct{} for empty elements
    fmt.Printf("  NoExport:           %t\n", config.NoExport) // create unexported types
    fmt.Printf("  Output:             %s\n", config.Output) // output filename

    fmt.Printf("  PackageName:        %s\n", config.PackageName) // package name
    fmt.Printf("  Pattern:            %s\n", config.Pattern) // filename pattern to observe
    fmt.Printf("  PreserveOrder:      %t\n", config.PreserveOrder) // preserve order of types and fields

    fmt.Printf("  TimeLayout:         %s\n", config.TimeLayout) // template for time layout
    fmt.Printf("  TopLevelAttributes: %t\n", config.TopLevelAttributes) // include top level attributes
    fmt.Printf("  TypesOnly:          %t\n", config.TypesOnly) // generate structs only, without header, package, or imports

    fmt.Printf("  UsePointersForOptionalFields: %t\n", config.UsePointersForOptionalFields) // use pointers for optional fields
    fmt.Printf("  UseRawToken:        %t\n", config.UseRawToken)// use encoding/xml.Decoder.RawToken

	fmt.Println("**** End Config *****")

}
