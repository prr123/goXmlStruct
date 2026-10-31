package xmlStruct


type YamlInp struct {
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
    UsePointersForOptionalFields bool `yaml:"UsePointersForOptionalFields"`// use pointers for optional fields
//    useRawToken                  = pflag.Bool("use-raw-token", xmlstruct.DefaultUseRawToken, "use encoding/xml.Decoder.RawToken")
    UseRawToken  bool `yaml:"UseRawToken"`// use encoding/xml.Decoder.RawToken
}
