package main


import (
	"fmt"
	"log"
	"os"
	"strings"

	"xmlstruct/xmlStruct"
//	"github.com/goccy/go-yaml"
    util "github.com/prr123/utility/utilLib"
)

var dbg bool

func main() {

	var config xmlStruct.YamlInp
	fmt.Printf("config: %v\n", config)

    numarg := len(os.Args)
    flags:=[]string{"dbg", "yaml", "inp", "out"}

    useStr := "/yaml=<yamlfile> /inp=<xml file> /out=<go file> [/dbg]"
    helpStr := "program to extract go struct from xml file"

    if numarg > len(flags) +1 {
        fmt.Println("too many arguments in cl!")
        fmt.Println("usage: %s %s\n", os.Args[0], useStr)
        os.Exit(-1)
    }

    if numarg == 1 || (numarg > 1 && os.Args[1] == "help") {
        fmt.Printf("help: %s\n", helpStr)
        fmt.Printf("usage is: %s %s\n", os.Args[0], useStr)
        os.Exit(1)
    }

    flagMap, err := util.ParseFlags(os.Args, flags)
    if err != nil {log.Fatalf("util.ParseFlags: %v\n", err)}

    dbg = false
    _, ok := flagMap["dbg"]
    if ok {dbg = true}

    yamlFilnam := ""
    yval, ok := flagMap["yaml"]
    if !ok {
        log.Fatalf("error -- no yaml flag provided!\n")
    } else {
        if yval.(string) == "none" {log.Fatalf("error -- no yaml file name provided!\n")}
        yamlFilnam = yval.(string)
        idx := strings.IndexByte(yamlFilnam, '.')
        if idx > -1 {log.Fatalf("error -- yaml file <%s> has an extension!\n", yamlFilnam)}
    }

    inpFilnam := ""
    ival, ok := flagMap["yaml"]
    if !ok {
        log.Fatalf("error -- no inp flag provided!\n")
    } else {
        if ival.(string) == "none" {log.Fatalf("error -- no inp file name provided!\n")}
        inpFilnam = ival.(string)
        idx := strings.IndexByte(inpFilnam, '.')
        if idx > -1 {log.Fatalf("error -- inp file <%s> has an extension!\n", inpFilnam)}
    }

    outFilnam := ""
    oval, ok := flagMap["out"]
    if !ok {
        log.Fatalf("error -- no out flag provided!\n")
    } else {
        if oval.(string) == "none" {log.Fatalf("error -- no out file name provided!\n")}
        outFilnam = oval.(string)
        idx := strings.IndexByte(outFilnam, '.')
        if idx > -1 {log.Fatalf("error -- out file <%s> has an extension!\n", outFilnam)}
    }

	// post processing
    yamlFilDir := "yaml/"
    yamlFilnam = yamlFilDir + yamlFilnam + ".yaml"
    yamlB, err := os.ReadFile(yamlFilnam)
    if err != nil {log.Fatalf("error -- cannot read yaml file: %v\n", err)}
	if dbg {
		fmt.Printf("*** yaml File ***\n")
		fmt.Printf("%s\n",yamlB)
		fmt.Println()
	}

    xmlFilDir := "inp/"
    xmlFilnam := xmlFilDir + inpFilnam + ".xml"
    xmlB, err := os.ReadFile(xmlFilnam)
    if err != nil {log.Fatalf("error -- cannot read xml file: %v\n", err)}
	if dbg {
		fmt.Printf("*** xml File ***\n")
		fmt.Printf("%s\n",xmlB)
		fmt.Println()
	}

    outFilDir := "out/"
    outFilnam = outFilDir + outFilnam + ".go"
    outFil, err := os.Create(outFilnam)
    if err != nil {log.Fatalf("error -- cannot create file: %v\n", err)}
	if dbg {
		fmt.Printf("*** out File created! ***\n")
		fmt.Println()
	}
	defer outFil.Close()
	outFil.WriteString("// go struct for " + xmlFilnam +"\n")


}
