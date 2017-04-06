package alphapoint

import (
	// "github.com/thrasher-/gocryptotrader/exchanges/alphapoint" //If we were outside of the folder we shoud declare the package here.
	"testing"
)

func TestSetDefaults(t *testing.T) {
	SetDefaults := Alphapoint{} // So I set an empty Alphapoint struct with member functions to setdefaults.
	SetDefaults.SetDefaults()   //so I acesss setdefaults struct into its member function and setdefaults should run.
}
