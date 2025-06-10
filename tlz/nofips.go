//go:build !fips

package tlz

// returns false if the binary was built with FIPS mode disabled
func FipsEnabled() bool {
    return false
}
