package GMP

import (
	internal "github.com/evilnode/cposc-go/GMP_internal"
)

// Sets the SSL mode.  Requests default to SSL,
// so setting this true will have no effect
func SSLMode(ssl bool) {
	internal.ENGA_NO_SSL = !ssl
}
