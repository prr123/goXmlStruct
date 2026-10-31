# goXMLStruct

This program is a fork from https://github.com/twpayne/go-xmlstruct

The main changes are:
(1) reorganisation of the code to allow local dev
(2) change the cli form spflag to prr123/util and create a yaml config file to specify the opetions

Why?
In my view these changes provide more clarity and allow easier execution

## code organisation

main => xmlstruct
|
element.go
generator.go
value.go
xmlstruct.go

