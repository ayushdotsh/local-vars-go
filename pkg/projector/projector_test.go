package projector_test

import (
	"testing"

	"github.com/ayushdotsh/local-vars-go/pkg/projector"
)

func getData() *projector.Data {
	return &projector.Data{
		Projector: map[string]map[string]string{
			"/": {
				"foo": "bar1",
				"fem": "is_great",
			},
			"/foo": {
				"foo": "bar2",
			},
			"/foo/bar": {
				"foo": "bar3",
			},
		},
	}
}

func getProjector(pwd string, data *projector.Data) *projector.Projector {
	return projector.CreateProjector(
		&projector.Config{
			Args:      []string{},
			Operation: projector.Print,
			Pwd:       pwd,
			Config:    "Hello!",
		},
		data,
	)
}

func test(t *testing.T, proj *projector.Projector, key, value string) {
	v, ok := proj.GetValue(key)
	if !ok {
		t.Errorf("Expected to find value \"foo\"")
	}
	if value != v {
		t.Errorf("Expected value to be %v but received \"%v\"", value, v)
	}
}

func TestGetValue(t *testing.T) {
	data := getData()
	pwd := "/foo/bar"
	proj := getProjector(pwd, data)

	test(t, proj, "foo", "bar3")
	test(t, proj, "fem", "is_great")
}

func TestSetValue(t *testing.T) {
	data := getData()
	pwd := "/foo/bar"
	proj := getProjector(pwd, data)

	test(t, proj, "foo", "bar3")
	proj.SetValue("foo", "bar4")
	test(t, proj, "foo", "bar4")

	proj.SetValue("fem", "is_super_great")
	test(t, proj, "fem", "is_super_great")

	proj = getProjector("/", data)
	test(t, proj, "fem", "is_great")
}

func TestRemoveValue(t *testing.T) {
	data := getData()
	pwd := "/foo/bar"
	proj := getProjector(pwd, data)

	test(t, proj, "foo", "bar3")
	proj.RemoveValue("foo")
	test(t, proj, "foo", "bar2")
}
