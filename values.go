package flags

import (
	"flag"
	"strconv"
	"strings"
)

type StringSliceValue []string

// New returns a new StringList satisfies the flag.Value interface. This is
// useful to be used with flag.FlagSet. For the global flag instance,
// StringSlice() and StringSliceVar() can be used.
func NewStringSlice(val []string, p *[]string) *StringSliceValue {
	return (*StringSliceValue)(p)
}

func (s *StringSliceValue) Set(val string) error {
	// if empty default is used
	if val == "" {
		return nil
	}

	*s = StringSliceValue(strings.Split(val, ","))
	return nil
}

func (s *StringSliceValue) Get() interface{} { return []string(*s) }

func (s *StringSliceValue) String() string { return strings.Join(*s, ",") }

// StringSlice defines a []string flag with specified name, default value, and usage
// string. The return value is the address of a []string variable that stores
// the value of the flag.
func StringSlice(value []string, name, usage string) *[]string {
	p := new([]string)
	flag.Var(NewStringSlice(value, p), name, usage)
	return p
}

// StringSliceVar defines a []string flag with specified name, default value,
// and usage string. The argument p points to a []string variable in which to
// store the value of the flag.
func StringSliceVar(p *[]string, value []string, name, usage string) {
	flag.Var(NewStringSlice(value, p), name, usage)
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
