package debugger

import (
	"fmt"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/color"
	"github.com/gen0cide/gscript/engine"
	"github.com/gen0cide/gscript/logger"
	"github.com/gen0cide/gscript/logger/standard"
	"github.com/robertkrimen/otto"
	readline "gopkg.in/readline.v1"
)

// Debugger is a wrapper type for handling interactive debug consoles in the genesis engine
type Debugger struct {
	VM        *engine.Engine
	Logger    logger.Logger
	OldLogger logger.Logger
	// Not Today
	// Asyncs    []*AsyncVM
}

// Not Today
// // ASyncVM is a container for a managed runtime
// type AsyncVM struct {
// 	VM        *otto.Otto
// 	Debugger  *Debugger
// 	Interrupt chan bool
// }

// // NewAsyncVM returns a new async container environment for a javascript runtime
// func NewAsyncVM(o *otto.Otto, d *Debugger) *AsyncVM {
// 	newInterrupt := make(chan bool, 1)
// 	return &AsyncVM{
// 		VM:        o,
// 		Debugger:  d,
// 		Interrupt: newInterrupt,
// 	}
// }

// func (a *AsyncVM) ExecuteWithInterrupt(eval string) {
// 	var wg sync.WaitGroup
// 	finChan := make(chan bool, 1)
// 	wg.Add(1)
// 	go func() {
// 		defer func() {
// 			if caught := recover(); caught != nil {
// 				a.Debugger.Logger.Infof("Halted AsyncVM with error: %v", caught)
// 				return
// 			}
// 			a.Debugger.Logger.Infof("AsyncVM finished execution")
// 			return
// 		}()
// 		a.VM.Eval(eval)
// 		finChan <- true
// 		wg.Done()
// 	}()

// 	go func() {
// 		wg.Wait()
// 		close(finChan)
// 	}()
// 	select {
// 	case <-a.Interrupt:
// 		a.VM.Interrupt <- func() {
// 			panic("gtfo")
// 		}
// 		wg.Done()
// 		return
// 	case <-finChan:
// 		wg.Done()
// 		return
// 	}
// }

// func (a *AsyncVM) Halt() {
// 	a.Interrupt <- true
// 	return
// }

// New returns a new debugger object wrapping the provided engine
func New(e *engine.Engine) *Debugger {
	dbgLogger := standard.NewStandardLogger(nil, "debugger", false, true)
	dbg := &Debugger{
		VM:        e,
		Logger:    dbgLogger,
		OldLogger: e.Logger,
		// Not Today
		// Asyncs:    []*AsyncVM{},
	}
	return dbg
}

// InjectDebugConsole injects the DebugConsole command into the runtime
func (d *Debugger) InjectDebugConsole() error {
	d.VM.VM.Set("_DEBUGGER", d)
	err := d.VM.VM.Set("DebugConsole", d.vmDebugConsole)
	if err != nil {
		return err
	}
	err = d.VM.VM.Set("SymbolTable", d.vmSymbolTable)
	if err != nil {
		return err
	}
	// Not Today
	// err = d.VM.VM.Set("async", d.vmAsync)
	// if err != nil {
	// 	return err
	// }
	return d.VM.VM.Set("TypeOf", d.vmTypeChecker)
}

// Not Today
// func (d *Debugger) vmAsync(call otto.FunctionCall) otto.Value {
// 	arg, _ := call.Argument(0).ToString()
// 	newVM := NewAsyncVM(call.Otto.Copy(), d)
// 	d.Asyncs = append(d.Asyncs, newVM)
// 	go newVM.ExecuteWithInterrupt(arg)
// 	retval, _ := call.Otto.ToValue(newVM)
// 	return retval
// }

func (d *Debugger) vmDebugConsole(call otto.FunctionCall) otto.Value {
	d.VM.SetLogger(d.Logger)
	d.runDebugger()
	d.VM.SetLogger(d.OldLogger)
	return otto.UndefinedValue()
}

func (d *Debugger) vmSymbolTable(call otto.FunctionCall) otto.Value {
	sym := d.AvailableFuncs()
	for ns, funcs := range sym {
		d.Logger.Infof(">>> %s Package\n\t%s\n", ns, strings.Join(funcs, "\n\t"))
	}
	return otto.UndefinedValue()
}

func (d *Debugger) vmTypeChecker(call otto.FunctionCall) otto.Value {
	if len(call.ArgumentList) == 0 {
		return d.VM.Raise("arg", "no argument provided")
	} else if len(call.ArgumentList) == 1 {
		val, err := call.Argument(0).Export()
		if err != nil {
			return d.VM.Raise("jsexport", "could not convert argument number 0")
		}
		retVal, _ := call.Otto.ToValue(spew.Sdump(val))
		return retVal
	} else {
		return d.VM.Raise("arg", "too many arguments provided")
	}
}

func (d *Debugger) runDebugger() error {
	prompt := fmt.Sprintf("%s%s", color.HiRedString("gscript"), color.HiWhiteString("> "))
	c := &readline.Config{
		Prompt: prompt,
	}
	cu, err := user.Current()
	if err == nil {
		c.HistoryFile = filepath.Join(cu.HomeDir, ".gscript_history")
	}

	rl, err := readline.NewEx(c)
	if err != nil {
		return err
	}
	standard.PrintLogo()
	title := fmt.Sprintf(
		"%s %s %s %s",
		color.HiWhiteString("***"),
		color.HiRedString("GSCRIPT"),
		color.YellowString("INTERACTIVE SHELL"),
		color.HiWhiteString("***"),
	)
	fmt.Fprintf(color.Output, "%s\n", title)
	rl.Refresh()

	for {
		l, err := rl.Readline()
		if err != nil {
			if err == readline.ErrInterrupt {
				if d != nil {
					d = nil
					rl.SetPrompt(prompt)
					rl.Refresh()
					continue
				}
				break
			}
			return err
		}
		if l == "" {
			continue
		}
		if l == "exit" {
			break
		}
		s, err := d.VM.VM.Compile("debugger", l)
		if err != nil {
			d.Logger.Errorf("%v", err)
			rl.SetPrompt(prompt)
			rl.Refresh()
			continue
		}
		v, err := d.VM.VM.Eval(s)
		if err != nil {
			if oerr, ok := err.(*otto.Error); ok {
				d.Logger.Error(oerr.Error())
			} else {
				d.Logger.Error(err.Error())
			}
		} else {
			rl.Write([]byte(fmt.Sprintf(">>> %s\n", v.String())))
		}
		rl.Refresh()
	}

	return rl.Close()
}

// AvailableFuncs returns the current debugger's symbol table
func (d *Debugger) AvailableFuncs() map[string][]string {
	ret := map[string][]string{}
	for name, p := range d.VM.Packages {
		ret[name] = []string{}
		idx := 0
		for _, f := range p.SymbolTable {
			ret[name] = append(ret[name], fmt.Sprintf("%d) %s", idx, f.Signature))
			idx++
		}
	}
	return ret
}
