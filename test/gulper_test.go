package main

import (
	"Gadget/pkg/k8sdeploy"
	"testing"
)

func TestHello(t *testing.T) {
	want := k8sdeploy.NewDefinition("test", "test")

	if got := k8sdeploy.NewDefinition("test", "test"); got.Name != want.Name {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
