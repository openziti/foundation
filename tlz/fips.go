package tlz

import "crypto/fips140"

// FipsEnabled reports whether the cryptography libraries are operating in FIPS
// 140-3 mode.
//
// Note that this also works with the Microsoft build of Go when using an OpenSSL backend.
// In that case, it reports whether the OpenSSL library is operating in FIPS mode.
func FipsEnabled() bool {
	return fips140.Enabled()
}
