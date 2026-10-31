package xmlStruct


type YamlInp struct {
//    attrNameSuffix               = pflag.String("attr-name-suffix", xmlstruct.DefaultAttrNameSuffix, "attribute name suffix")
    AttrNameSuffix	string // attribute name suffix
//    charDataFieldName            = pflag.String("char-data-field-name", xmlstruct.DefaultCharDataFieldName, "char data field name")
    CharDataFieldName string // char data field name
//    compactTypes                 = pflag.Bool("compact-types", xmlstruct.DefaultCompactTypes, "create compact types")
    CompactTypes  bool       // create compact types
//    elemNameSuffix               = pflag.String("elem-name-suffix", xmlstruct.DefaultElemNameSuffix, "element name suffix")
    ElemNameSuffix string  // element name suffix
//    formatSource                 = pflag.Bool("format-source", xmlstruct.DefaultFormatSource, "format source")
    FormatSource bool  // format source
//    header                       = pflag.String("header", xmlstruct.DefaultHeader, "header")
    Header string  // header
//    ignoreErrors                 = pflag.Bool("ignore-errors", false, "ignore errors")
    IgnoreErrors bool  // ignore errors
//    ignoreNamespaces             = pflag.Bool("ignore-namespaces", true, "ignore namespaces")
    IgnoreNamespaces bool  // ignore namespaces
//    imports                      = pflag.Bool("imports", xmlstruct.DefaultImports, "generate import statements")
    Imports bool  // generate import statements
//    intType                      = pflag.String("int-type", xmlstruct.DefaultIntType, "int type")
    IntType string  //int type
//    namedRoot                    = pflag.Bool("named-root", xmlstruct.DefaultNamedRoot, "create an XMLName field for the root element")
    NamedRoot bool // create an XMLName field for the root element
//    namedTypes                   = pflag.Bool("named-types", xmlstruct.DefaultNamedTypes, "create named types for all elements")
    NamedTypes bool // create named types for all elements
//    noEmptyElements              = pflag.Bool("no-empty-elements", !xmlstruct.DefaultEmptyElements, "use type string instead of struct{} for empty ele>
    NoEmptyElements bool // use type string instead of struct{} for empty elements
//    noExport                     = pflag.Bool("no-export", false, "create unexported types")
    NoExport bool //                    = pflag.Bool("no-export", false, "create unexported types")
//    output                       = pflag.String("output", "", "output filename")
    Output string // output filename
//    packageName                  = pflag.String("package-name", "main", "package name")
    PackageName string // package name
//    pattern                      = pflag.String("pattern", "", "filename pattern to observe")
    Pattern string // filename pattern to observe
//    preserveOrder                = pflag.Bool("preserve-order", xmlstruct.DefaultPreserveOrder, "preserve order of types and fields")
    PreserveOrder bool // preserve order of types and fields

//    timeLayout                   = pflag.String("time-layout", "2006-01-02T15:04:05Z", "time layout")
    TimeLayout string // template for time layout
//    topLevelAttributes           = pflag.Bool("top-level-attributes", xmlstruct.DefaultTopLevelAttributes, "include top level attributes")
    TopLevelAttributes bool // include top level attributes
//    typesOnly                    = pflag.Bool("types-only", false, "generate structs only, without header, package, or imports")
    TypesOnly bool // generate structs only, without header, package, or imports
//    usePointersForOptionalFields = pflag.Bool("use-pointers-for-optional-fields", xmlstruct.DefaultUsePointersForOptionalFields, "use pointers for opt>
    UsePointersForOptionalFields bool // use pointers for optional fields
//    useRawToken                  = pflag.Bool("use-raw-token", xmlstruct.DefaultUseRawToken, "use encoding/xml.Decoder.RawToken")
    UseRawToken  bool // use encoding/xml.Decoder.RawToken
}
