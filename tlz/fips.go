//go:build goexperiment.opensslcrypto

package tlz

import "crypto/boring"

// returns true if the binary was built with FIPS mode enabled
func FipsEnabled() bool {
    return boring.Enabled()
}
