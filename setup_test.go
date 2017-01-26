package ace

import (
	"testing"

	"github.com/mholt/caddy"
	"github.com/mholt/caddy/caddyhttp/httpserver"
)

func TestSetup(t *testing.T) {
	c := caddy.NewTestController("http", `ace`)
	err := setup(c)
	if err != nil {
		t.Errorf("Expected no errors, but got: %v", err)
	}
	mids := httpserver.GetConfig(c).Middleware()
	if len(mids) == 0 {
		t.Fatal("Expected middleware, got 0 instead")
	}

	handler := mids[0](httpserver.EmptyNext)
	myHandler, ok := handler.(Ace)
	if !ok {
		t.Fatalf("Expected handler to be type BasicAuth, got: %#v", handler)
	}

	if !httpserver.SameNext(myHandler.Next, httpserver.EmptyNext) {
		t.Error("'Next' field of handler was not set properly")
	}
}

func TestAceParse(t *testing.T) {

	tests := []struct {
		inputAceConfig    string
		shouldErr         bool
		expectedAceConfig []Config
	}{
		{`ace`, false, []Config{
			{},
		}},
		{`ace`, true, []Config{
			{},
		}},
	}

	for i, test := range tests {
		c := caddy.NewTestController("http", test.inputAceConfig)
		actualAceConfigs, err := aceParse(c)
		if err == nil && test.shouldErr {
			t.Errorf("Test %d didn't error, but it should have", i)
		} else if err != nil && !test.shouldErr {
			t.Errorf("Test %d errored, but it shouldn't have; got '%v'", i, err)
		}
		for j, actualAceConfig := range actualAceConfigs {
			t.Errorf("d% s% s%", i, j, actualAceConfig)
		}

	}

}
