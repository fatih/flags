package flags

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

func TestStringList(t *testing.T) {
	regions := StringList(nil, "to", "Regions to be used")

	os.Args = []string{"images", "-to", "us-east-1,eu-west-2"}
	flag.Parse()

	want := []string{"us-east-1", "eu-west-2"}
	if !reflect.DeepEqual(*regions, want) {
		t.Errorf("Regions = %q, want %q", regions, want)
	}
}

func TestStringListVar(t *testing.T) {
	var regions []string
	StringListVar(&regions, nil, "tos", "Regions to be used")

	os.Args = []string{"images", "-tos", "us-east-1,eu-west-2"}
	flag.Parse()

	want := []string{"us-east-1", "eu-west-2"}
	if !reflect.DeepEqual(regions, want) {
		t.Errorf("Regions = %q, want %q", regions, want)
	}
}

func TestStringListFlagSet(t *testing.T) {
	f := flag.NewFlagSet("TestTags", flag.PanicOnError)

	var regions []string
	f.Var(NewStringList(nil, &regions), "to", "Regions to be used")
	f.Parse([]string{"-to", "us-east-1,eu-west-2"})

	want := []string{"us-east-1", "eu-west-2"}
	if !reflect.DeepEqual(regions, want) {
		t.Errorf("Regions = %q, want %q", regions, want)
	}
}

func TestIntList(t *testing.T) {
	f := flag.NewFlagSet("TestTags", flag.PanicOnError)

	var ids []int
	f.Var(IntListVar(&ids), "ids", "Ids to be used")
	f.Parse([]string{"-ids", "123,456"})

	want := []int{123, 456}
	if !reflect.DeepEqual(ids, want) {
		t.Errorf("Ids = %q, want %q", ids, want)
	}
}
