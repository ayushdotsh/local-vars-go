package projector_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ayushdotsh/local-vars-go/pkg/projector"
)

func getOpts(args []string) *projector.Opts {
	opts := &projector.Opts{
		Args:   args,
		Config: "",
		Pwd:    "",
	}
	return opts
}

func testConfig(t *testing.T, args []string, extectedArgs []string, operation projector.Operation) {
	opts := getOpts(args)
	if opts == nil {
		t.Fatal("Expected opts to be non-nil")
	}
	fmt.Printf("Testing with opts: %+v\n", opts)
	config, err := projector.NewConfig(opts)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !reflect.DeepEqual(extectedArgs, config.Args) {
		t.Errorf("Expected args to %+v but got %+v", extectedArgs, config.Args)
	}

	if config.Operation != operation {
		t.Errorf("operation expected was %v but got %v", operation, config.Operation)
	}
}

func TestConfigPrint(t *testing.T) {
	testConfig(t, []string{}, []string{}, projector.Print)
}

func TestConfigPrintKey(t *testing.T) {
	testConfig(t, []string{"foo"}, []string{"foo"}, projector.Print)
}

func TestConfigAddKeyValue(t *testing.T) {
	testConfig(t, []string{"add", "foo", "bar"}, []string{"foo", "bar"}, projector.Add)
}

func TestConfigRemoveKey(t *testing.T) {
	testConfig(t, []string{"rm", "foo"}, []string{"foo"}, projector.Remove)
}
