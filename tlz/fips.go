//go:build fips

package tlz

// returns true if the binary was built with FIPS mode enabled
func FipsEnabled() bool {
    return true
}
