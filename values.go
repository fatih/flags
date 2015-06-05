package flags

import (
	"flag"
	"strconv"
	"strings"
)

type stringList []string

// New returns a new StringList satisfies the flag.Value interface. This is
// useful to be used with flag.FlagSet. For global flag, StringList() and
// StringListVar() can be used.
func NewStringList(val []string, p *[]string) *stringList {
	return (*stringList)(p)
}

func (s *stringList) Set(val string) error {
	// if empty default is used
	if val == "" {
		return nil
	}

	*s = stringList(strings.Split(val, ","))
	return nil
}

func (s *stringList) Get() interface{} { return []string(*s) }

func (s *stringList) String() string { return strings.Join(*s, ",") }

// String defines a []string flag with specified name, default value, and usage
// string. The return value is the address of a []string variable that stores
// the value of the flag.
func StringList(value []string, name, usage string) *[]string {
	p := new([]string)
	flag.Var(NewStringList(value, p), name, usage)
	return p
}

// StringListVar defines a []string flag with specified name, default value,
// and usage string. The argument p points to a []string variable in which to
// store the value of the flag.
func StringListVar(p *[]string, value []string, name, usage string) {
	flag.Var(NewStringList(value, p), name, usage)
}

// IntList is an implementation of flag.Value interface that accepts a comma
// separated list of integers and populates a []int slice.
//
// Example:
//
//  var ids []int
//  flag.Var(flags.IntListVar(&ids), "server", "IDs to be used")
type IntList []int

// New returns a new *IntList.
func IntListVar(p *[]int) *IntList {
	return (*IntList)(p)
}

func (i *IntList) Set(val string) error {
	// if empty default is used
	if val == "" {
		return nil
	}

	var list []int
	for _, in := range strings.Split(val, ",") {
		i, err := strconv.Atoi(in)
		if err != nil {
			return err
		}

		list = append(list, i)
	}

	*i = IntList(list)
	return nil
}

func (i *IntList) Get() interface{} { return []int(*i) }

func (i *IntList) String() string {
	var list []string
	for _, in := range *i {
		list = append(list, strconv.Itoa(in))
	}
	return strings.Join(list, ",")
}
