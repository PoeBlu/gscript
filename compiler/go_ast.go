package compiler

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"reflect"
	"regexp"
	"sync"

	"github.com/gen0cide/gscript/compiler/computil"
	"github.com/gen0cide/gscript/compiler/translator"
)

var (
	builtInGoTypes = map[string]bool{
		"bool":       true,
		"byte":       true,
		"complex128": true,
		"complex64":  true,
		"error":      true,
		"float32":    true,
		"float64":    true,
		"int":        true,
		"int8":       true,
		"int16":      true,
		"int32":      true,
		"int64":      true,
		"rune":       true,
		"string":     true,
		"uint":       true,
		"uint8":      true,
		"uint16":     true,
		"uint32":     true,
		"uint64":     true,
		"uintptr":    true,
	}

	invalidGoTypes = map[string]bool{
		"complex128": true,
		"complex64":  true,
		// "float32":    true,
		// "float64":    true,
		// "uintptr": true,
	}

	binaryImports = map[string]string{
		"bytes":           "bytes",
		"compress/gzip":   "gzip",
		"crypto/aes":      "aes",
		"crypto/cipher":   "cipher",
		"encoding/base64": "base64",
		"fmt":             "fmt",
		"io":              "io",
		"github.com/gen0cide/gscript/engine":          "engine",
		"github.com/robertkrimen/otto":                "otto",
		"github.com/gen0cide/gscript/debugger":        "debugger",
		"github.com/gen0cide/gscript/logger/standard": "standard",
	}

	funcRegexp  = regexp.MustCompile(`^func\({1}(?P<args>.*?)?\){1}\s*\(?(?P<rets>.*?)\)??$`)
	multipleRet = regexp.MustCompile(`,`)
)

// MaskedImport is used to separate import namespaces within the intermediate representation
type MaskedImport struct {
	// ImportPath of the masked Import
	ImportPath string

	// OldAlias represents the alias in the target package source
	OldAlias string

	// NewAlias represents the aliased package name in the intermediate representation
	NewAlias string
}

// GoPackage holds all the information about a Golang package that is being resolved to a given script
type GoPackage struct {
	sync.RWMutex

	// Dir is the local path where this package is found
	Dir string

	// MaskedName is the masked import representation of this gopackage
	MaskedName string

	// ImportPath is the golang import path used for this package
	ImportPath string

	// Name defines the go package's name
	Name string

	// VM references the script that is importing this package
	VM *GenesisVM

	// Namespace is the namespace aliased in the parent script
	Namespace string

	// ImportKey is the import path defined in the parent script
	ImportKey string

	// ScriptCallers maps the go package function names to genesis AST function calls
	ScriptCallers map[string]*FunctionCall

	// ImportsByFile is the map of each file within the package and what dependencies it imports
	ImportsByFile map[string][]*ast.ImportSpec

	// ImportsByAlias is a map of each import's alias and its dependency information
	ImportsByAlias map[string]*ast.ImportSpec

	// FuncToFileMap maps each public function within the package to it's corrasponding source file
	FuncToFileMap map[string]string

	// FuncTable is the map of each function name to it's Golang AST declaration
	FuncTable map[string]*ast.FuncDecl

	// LinkedFuncs defines references to the dynamically linked functions for this go package
	LinkedFuncs []*LinkedFunction

	// Reference to know if this go package is part of the standard library
	IsStandardLib bool

	// FileSet is used by the parser to interpret the current golang's file tokens
	FileSet *token.FileSet
}

// GoParamDef defines a type to represent parameters found in a Golang function declaration (arguments or return types)
type GoParamDef struct {
	// SigBuffer is used to create a definition to the actual type declaration in the genesis compiler's linker
	SigBuffer bytes.Buffer

	// NameBuffer is used to create a label that the genesis linker can use to create it's translations
	NameBuffer bytes.Buffer

	// ImportRefs holds a mapping of any golang dependencies required for this parameter's type declaration
	ImportRefs map[string]*ast.ImportSpec

	// VarName holds a representation of the logical representation of a parameter's label with it's offset appended
	VarName string

	// ParamIdx will be the relative position within the parameter declaration.
	// This is inclusive to multiple declarations of the same type. Example:
	//
	// func Foo(a, b string, c int) {}
	// GoParamDef objects for "a" and "b" would hold the same ParamIdx, but different
	// ArgOffset values
	ParamIdx int

	// ArgOffset defines the absolute position within the parameter declaration. It will increment
	// Regardless of multiple labels defined in the same type.
	// This is used by the linker to correctly translate arguments for golang functions
	ArgOffset int

	// ExtSig will hold the final result of the SigBuffer rendering as a string
	ExtSig string

	// GoLabel is used to represent the label name within Golang
	GoLabel string

	// LinkedFUnction is used to reference the parent LinkedFunction object
	LinkedFunction *LinkedFunction
}

// NewGoPackage is a constructor for a gopackage that will be used in dynamically linking native code
func NewGoPackage(v *GenesisVM, ns, ikey string, stdlib bool) *GoPackage {
	return &GoPackage{
		VM:             v,
		Namespace:      ns,
		ImportKey:      ikey,
		ImportPath:     ikey,
		ScriptCallers:  map[string]*FunctionCall{},
		ImportsByFile:  map[string][]*ast.ImportSpec{},
		ImportsByAlias: map[string]*ast.ImportSpec{},
		FuncToFileMap:  map[string]string{},
		FuncTable:      map[string]*ast.FuncDecl{},
		LinkedFuncs:    []*LinkedFunction{},
		IsStandardLib:  stdlib,
		MaskedName:     computil.RandLowerAlphaString(6),
	}
}

// NewGoParamDef creates a new definition object for a go parameter (either return or argument) and returns a pointer to itself.
func NewGoParamDef(l *LinkedFunction, idx int) *GoParamDef {
	gpd := &GoParamDef{
		LinkedFunction: l,
		ImportRefs:     map[string]*ast.ImportSpec{},
		ParamIdx:       idx,
	}
	gpd.NameBuffer.WriteString("_")
	return gpd
}

// NewMaskedImport creates a new import mask based on an import path and old alias
func NewMaskedImport(ip, oa string) *MaskedImport {
	alias := computil.RandLowerAlphaString(6)
	if val, ok := binaryImports[ip]; ok {
		alias = val
	}
	return &MaskedImport{
		ImportPath: ip,
		OldAlias:   oa,
		NewAlias:   alias,
	}
}

// IsDefaultImport tests a golang import path to determine if it is already defined in the intermediate representation
func IsDefaultImport(ip string) bool {
	return binaryImports[ip] != ""
}

// GetDefaultImportNamespace returns the corrasponding namespace to the import path provided that is used in the intermediate representation
func GetDefaultImportNamespace(ip string) string {
	return binaryImports[ip]
}

// import (
//   rekt "net/url"
// )
// EXAMPLE: func Foo(a0 map[*rekt.URL][]*ast.Field)
// 0 ""
// 1 "map["
// 2 "map[*"
// 3 "map[*url"
// 3.5 "map[*url."
// 3.7 "map[*url.URL"
// 4 "map[*url.URL]"
// 5 "map[*url.URL][]"
// 6 "map[*url.URL][]*"
// 7 "map[*url.URL][]*ast"
// 7.5 "map[*url.URL][]*ast."
// 7.7 "map[*url.URL][]*ast.Field"

// Interpret is a recursive walk function that is used to dispatch the next walker
// depending on the type of the provided interface (i). This is used to build up
// buffers of both names and golang type declarations to be used during linking.
func (p *GoParamDef) Interpret(i interface{}) error {
	switch t := i.(type) {
	case *ast.StarExpr:
		return p.ParseStarExpr(t)
	case *ast.SelectorExpr:
		return p.ParseSelectorExpr(t)
	case *ast.Ident:
		return p.ParseIdent(t)
	case *ast.ArrayType:
		return p.ParseArrayType(t)
	case *ast.MapType:
		return p.ParseMapType(t)
	case *ast.ChanType:
		return fmt.Errorf("function %s includes an unsupported parameter type: %s", p.LinkedFunction.GoDecl.Name.Name, "chan")
	case *ast.FuncType:
		return fmt.Errorf("function %s includes an unsupported parameter type: %s", p.LinkedFunction.GoDecl.Name.Name, "func")
	case *ast.InterfaceType:
		return fmt.Errorf("function %s includes an unsupported parameter type: %s", p.LinkedFunction.GoDecl.Name.Name, "interface{}")
	case *ast.StructType:
		return fmt.Errorf("function %s includes an unsupported parameter type: %s", p.LinkedFunction.GoDecl.Name.Name, "struct")
	default:
		valType := reflect.ValueOf(t)
		return fmt.Errorf("could not determine the golang ast type of %s in func %s", valType.Type().String(), p.LinkedFunction.Function)
	}
}

// ParseMapType interprets a golang map type into the appropriate GoParamDef structure
func (p *GoParamDef) ParseMapType(a *ast.MapType) error {
	p.SigBuffer.WriteString("map[")
	p.NameBuffer.WriteString("MapOf")
	err := p.Interpret(a.Key)
	if err != nil {
		return err
	}
	p.SigBuffer.WriteString("]")
	p.NameBuffer.WriteString("WithValType")
	err = p.Interpret(a.Value)
	return err
}

// ParseArrayType interprets a golang array/slice type into the appropriate GoParamDef structure
func (p *GoParamDef) ParseArrayType(a *ast.ArrayType) error {
	p.SigBuffer.WriteString("[]")
	p.NameBuffer.WriteString("ArrayOf")
	return p.Interpret(a.Elt)
}

// ParseSelectorExpr interprets a golang namespace external to the function declarations package
// and maps it into the appropriate GoParamDef structure
func (p *GoParamDef) ParseSelectorExpr(a *ast.SelectorExpr) error {
	x, ok := a.X.(*ast.Ident)
	if !ok {
		return fmt.Errorf("could not parse selector namespace in func %s", p.LinkedFunction.Function)
	}
	resolved, err := p.LinkedFunction.CanResolveImportDep(x.Name)
	if err != nil {
		return err
	}
	p.SigBuffer.WriteString(resolved)
	p.NameBuffer.WriteString(resolved)
	p.SigBuffer.WriteString(".")
	p.NameBuffer.WriteString("_")
	p.SigBuffer.WriteString(a.Sel.Name)
	p.NameBuffer.WriteString(a.Sel.Name)
	return nil
}

// ParseStarExpr interprets a golang pointer into the appropriate GoParamDef structure
func (p *GoParamDef) ParseStarExpr(a *ast.StarExpr) error {
	p.SigBuffer.WriteString("*")
	p.NameBuffer.WriteString("PointerTo")
	return p.Interpret(a.X)
}

// ParseIdent interprets a golang identifier into the appropriate GoParamDef structure
func (p *GoParamDef) ParseIdent(a *ast.Ident) error {
	if ok := builtInGoTypes[a.Name]; !ok {
		if IsDefaultImport(p.LinkedFunction.GoPackage.ImportPath) {
			p.SigBuffer.WriteString(GetDefaultImportNamespace(p.LinkedFunction.GoPackage.ImportPath))
			p.SigBuffer.WriteString(".")
			p.NameBuffer.WriteString(GetDefaultImportNamespace(p.LinkedFunction.GoPackage.ImportPath))
			p.NameBuffer.WriteString("_")
		} else {
			p.SigBuffer.WriteString(p.LinkedFunction.GoPackage.MaskedName)
			p.SigBuffer.WriteString(".")
			p.NameBuffer.WriteString(p.LinkedFunction.GoPackage.MaskedName)
			p.NameBuffer.WriteString("_")
		}
	}
	p.SigBuffer.WriteString(a.Name)
	p.NameBuffer.WriteString(a.Name)
	return nil
}

// IsBuiltInGoType takes a string argument and determines if is a valid built-in type
// in golang
func IsBuiltInGoType(s string) bool {
	return builtInGoTypes[s]
}

// WalkGoFileAST walks the AST of a golang file and determines if it should be included as a linked
// function based on one of the following statements being true:
// Parent GoPackage is a member of the standard library
// OR
// Compiler option ImportAllNativeFunc is set to true
// OR
// VM Script calls this function explicitly
func (gop *GoPackage) WalkGoFileAST(goast *ast.File, wg *sync.WaitGroup, errChan chan error) {
	ast.Inspect(goast, func(n ast.Node) bool {
		// TODO: Started work on trying to grab CONST declarations, but fuck no. That gets wacky fast.
		// Revisit when more time have I will.
		// decl, ok := n.(*ast.GenDecl)
		// if ok {
		// 	if decl.Tok == token.CONST {
		// 		for _, cdecl := range decl.Specs {
		// 			if valdecl, ok := cdecl.(*ast.ValueSpec); ok {
		// 				if len(valdecl.Names) > 0 && valdecl.Names[0].IsExported() {
		// 					gop.VM.Logger.Infof("Discovered CONST: %s", spew.Sdump(valdecl))
		// 				}
		// 				// for
		// 				// if valdecl.Name[0].IsExported() {
		// 				// 	gop.VM.Logger.Infof("Discovered CONST: %s.%s = %s", gop.Name, valdecl.Name.Name, valdecl.)
		// 				// }
		// 			}
		// 		}
		// 	}
		// }
		fn, ok := n.(*ast.FuncDecl)
		if ok {
			funcName := fn.Name.Name
			if fn.Name.IsExported() && fn.Recv == nil {
				gop.Lock()
				caller := gop.ScriptCallers[funcName]
				if caller == nil && gop.IsStandardLib == false && gop.VM.Options.ImportAllNativeFuncs == false {
					gop.Unlock()
					return true
				}
				sig := new(bytes.Buffer)
				printer.Fprint(sig, gop.FileSet, fn.Type)
				lf, err := gop.VM.Linker.NewLinkedFunction(
					funcName,
					caller,
					goast,
					fn,
					goast.Imports,
					gop,
				)
				if err != nil {
					gop.Unlock()
					errChan <- err
					wg.Done()
					return false
				}
				match := funcRegexp.FindStringSubmatch(sig.String())
				result := make(map[string]string)
				for i, name := range funcRegexp.SubexpNames() {
					if i != 0 && name != "" {
						result[name] = match[i]
					}
				}

				newSigBuf := new(bytes.Buffer)
				if result["rets"] != "" {
					if multipleRet.MatchString(result["rets"]) {
						newSigBuf.WriteString("[")
					}
					newSigBuf.WriteString(result["rets"])
					if multipleRet.MatchString(result["rets"]) {
						newSigBuf.WriteString("]")
					}
					newSigBuf.WriteString(" = ")
				}
				if gop.IsStandardLib {
					newSigBuf.WriteString("G.")
					newSigBuf.WriteString(gop.Name)
				} else {
					newSigBuf.WriteString(gop.Namespace)
				}
				newSigBuf.WriteString(".")
				newSigBuf.WriteString(funcName)
				newSigBuf.WriteString("(")
				newSigBuf.WriteString(result["args"])
				newSigBuf.WriteString(")")
				lf.Signature = newSigBuf.String()
				if len(gop.ImportsByFile[goast.Name.Name]) == 0 {
					gop.ImportsByFile[goast.Name.Name] = goast.Imports
				}
				gop.FuncTable[funcName] = fn
				gop.FuncToFileMap[funcName] = goast.Name.Name
				gop.LinkedFuncs = append(gop.LinkedFuncs, lf)
				gop.Unlock()
			}
		}
		return true
	})
	wg.Done()
	return
}

// SanityCheckScriptCallers enumerates all of the parent gopackage's script callers looking for any
// callers who do not have an entry in this go package's symbol table
func (gop *GoPackage) SanityCheckScriptCallers() error {
	for fnName := range gop.ScriptCallers {
		if gop.FuncTable[fnName] == nil {
			return fmt.Errorf("function %s is not a valid function in package %s", fnName, gop.Name)
		}
	}
	return nil
}

// SuccessfullyLinkedFuncs is used during rendering to make sure that only linked functions that successfully
// swizzled are going to be built into the source
func (gop *GoPackage) SuccessfullyLinkedFuncs() []*LinkedFunction {
	lf := []*LinkedFunction{}
	for _, l := range gop.LinkedFuncs {
		if l.SwizzleSuccessful {
			lf = append(lf, l)
		}
	}

	return lf
}

// BuiltInTranslationRequired is a compiler helper to determine if the param definition requires a built in translation
func (p *GoParamDef) BuiltInTranslationRequired() bool {
	return translator.BuiltInMap[p.ExtSig] != ""
}

// BuiltInJSType returns the type that JS will return that we need to convert the ParamDef to golang
func (p *GoParamDef) BuiltInJSType() string {
	return translator.BuiltInMap[p.ExtSig]
}
