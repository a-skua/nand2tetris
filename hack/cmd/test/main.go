// Command test runs a Hack machine code program (.hack) on [hack.Computer]
// and prints the registers and memory after the run.
//
// Usage:
//
//	test [flags] program.hack
//
// The run stops when the program reaches a halt loop such as
//
//	(END)
//	@END
//	0;JMP
//
// or when -cycles instructions have been executed.
//
// Flags:
//
//	-cycles n    maximum number of instructions to execute (default 10000)
//	-ram list    RAM addresses to print, e.g. "0-2,16" (default "0-15")
//	-set list    initial memory values, e.g. "0=3,1=5"
//	-trace       print the registers after each instruction
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/a-skua/nand2tetris/hack"
	"github.com/a-skua/nand2tetris/hack/builtin"
)

// memorySize is the number of addressable words of [hack.Memory] (up to Keyboard).
const memorySize = 0x6001

// Command-line flags.
var (
	cycles  int
	ramList string
	setList string
	trace   bool
)

func init() {
	flag.IntVar(&cycles, "cycles", 10000, "maximum number of instructions to execute")
	flag.StringVar(&ramList, "ram", "0-15", `RAM addresses to print, e.g. "0-2,16"`)
	flag.StringVar(&setList, "set", "", `initial memory values, e.g. "0=3,1=5"`)
	flag.BoolVar(&trace, "trace", false, "print the registers after each instruction")
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "usage: test [flags] program.hack")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}
	if err := run(flag.Arg(0), os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "test:", err)
		os.Exit(1)
	}
}

// run executes the program in the file name and writes the result to w.
func run(name string, w io.Writer) error {
	addresses, err := parseAddresses(ramList)
	if err != nil {
		return fmt.Errorf("-ram: %w", err)
	}
	sets, err := parseSets(setList)
	if err != nil {
		return fmt.Errorf("-set: %w", err)
	}

	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	program, err := parseProgram(f)
	if err != nil {
		return fmt.Errorf("%s: %w", name, err)
	}
	if len(program) > 1<<15 {
		return fmt.Errorf("%s: program is longer than 32768 words", name)
	}

	c := hack.New(program)
	for _, s := range sets {
		c.Memory.Clock(builtin.NewWord(int16(s.value)), builtin.ON, address(s.address))
	}

	executed, halted := execute(c, cycles, func(n int, pc builtin.Word) {
		if trace {
			fmt.Fprintf(w, "%6d: ROM[%d]  ", n, toInt(pc))
			printRegisters(w, c)
		}
	})

	if halted {
		fmt.Fprintf(w, "halted after %d cycles\n", executed)
	} else {
		fmt.Fprintf(w, "stopped after %d cycles (limit)\n", executed)
	}
	printRegisters(w, c)
	for _, a := range addresses {
		v := c.Memory.Clock(builtin.Word{}, builtin.OFF, address(a))
		fmt.Fprintf(w, "RAM[%d] = %d\n", a, int16(toInt(v)))
	}
	return nil
}

// state is the registers that determine the next instruction.
type state struct{ pc, a, d builtin.Word }

func snapshot(c *hack.Computer) state {
	return state{c.PC.Bits(), c.A.Bits(), c.D.Bits()}
}

// execute runs c for at most limit instructions, calling step after each one.
// It reports the number of executed instructions and whether c reached a halt loop.
//
// A halt loop is detected when the registers return to the state of one or two
// instructions before without any memory write in between; from then on the
// computer repeats the same instructions forever (as long as Keyboard does not change).
func execute(c *hack.Computer, limit int, step func(n int, pc builtin.Word)) (n int, halted bool) {
	if limit <= 0 {
		return 0, false
	}
	history := []state{snapshot(c)}
	wrote := []bool{false}
	for pc := range c.Clock() {
		n++
		step(n, pc)

		s := snapshot(c)
		history = append(history, s)
		wrote = append(wrote, writesMemory(c, pc))
		if len(history) > 3 {
			history, wrote = history[1:], wrote[1:]
		}
		last := len(history) - 1
		for period := 1; period <= 2 && period <= last; period++ {
			if history[last-period] == s && !anyTrue(wrote[last-period+1:]) {
				return n, true
			}
		}
		if n >= limit {
			break
		}
	}
	return n, false
}

// writesMemory reports whether the instruction at pc writes to M.
func writesMemory(c *hack.Computer, pc builtin.Word) bool {
	in := c.ROM32K.Clock(builtin.Word{}, builtin.OFF, [15]builtin.Pin(pc[0:15]))
	return in[15] == builtin.ON && in[3] == builtin.ON
}

func anyTrue(bs []bool) bool {
	for _, b := range bs {
		if b {
			return true
		}
	}
	return false
}

func printRegisters(w io.Writer, c *hack.Computer) {
	fmt.Fprintf(w, "PC=%d A=%d D=%d\n",
		toInt(c.PC.Bits()), int16(toInt(c.A.Bits())), int16(toInt(c.D.Bits())))
}

// parseProgram reads Hack machine code: one 16-digit binary word per line.
// Blank lines and "//" comments are ignored.
func parseProgram(r io.Reader) ([][16]builtin.Pin, error) {
	var program [][16]builtin.Pin
	sc := bufio.NewScanner(r)
	for line := 1; sc.Scan(); line++ {
		s, _, _ := strings.Cut(sc.Text(), "//")
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		if len(s) != 16 {
			return nil, fmt.Errorf("line %d: %q is not a 16-bit word", line, s)
		}
		v, err := strconv.ParseUint(s, 2, 16)
		if err != nil {
			return nil, fmt.Errorf("line %d: %q is not a 16-bit word", line, s)
		}
		program = append(program, builtin.NewWord(int16(v)))
	}
	return program, sc.Err()
}

// parseAddresses parses a comma-separated list of addresses and ranges such as "0-2,16".
func parseAddresses(s string) ([]int, error) {
	var addresses []int
	if strings.TrimSpace(s) == "" {
		return nil, nil
	}
	for item := range strings.SplitSeq(s, ",") {
		lo, hi, isRange := strings.Cut(strings.TrimSpace(item), "-")
		from, err := parseAddress(lo)
		if err != nil {
			return nil, err
		}
		to := from
		if isRange {
			if to, err = parseAddress(hi); err != nil {
				return nil, err
			}
			if to < from {
				return nil, fmt.Errorf("invalid range %q", item)
			}
		}
		for a := from; a <= to; a++ {
			addresses = append(addresses, a)
		}
	}
	return addresses, nil
}

type set struct{ address, value int }

// parseSets parses a comma-separated list of address=value pairs such as "0=3,1=-5".
func parseSets(s string) ([]set, error) {
	var sets []set
	if strings.TrimSpace(s) == "" {
		return nil, nil
	}
	for item := range strings.SplitSeq(s, ",") {
		a, v, ok := strings.Cut(strings.TrimSpace(item), "=")
		if !ok {
			return nil, fmt.Errorf("%q is not address=value", item)
		}
		address, err := parseAddress(a)
		if err != nil {
			return nil, err
		}
		value, err := strconv.ParseInt(strings.TrimSpace(v), 0, 32)
		if err != nil || value < -0x8000 || value > 0xFFFF {
			return nil, fmt.Errorf("%q is not a 16-bit value", v)
		}
		sets = append(sets, set{address, int(value)})
	}
	return sets, nil
}

func parseAddress(s string) (int, error) {
	a, err := strconv.ParseUint(strings.TrimSpace(s), 0, 16)
	if err != nil || a >= memorySize {
		return 0, fmt.Errorf("%q is not a memory address (0-%d)", s, memorySize-1)
	}
	return int(a), nil
}

// address returns the low 15 bits of v as a memory address.
func address(v int) [15]builtin.Pin {
	w := builtin.NewWord(int16(v))
	return [15]builtin.Pin(w[0:15])
}

// toInt returns w as an unsigned integer.
func toInt(w builtin.Word) (v int) {
	for i, b := range w {
		v |= int(b&1) << i
	}
	return
}
