package upgrade

import "syscall"

var (
	// SIG_UPGRADE ...
	SIG_UPGRADE = syscall.SIGUSR1
)
