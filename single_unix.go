//go:build !windows
// +build !windows

package singleproc

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"
)

// Lock tries to obtain an excluded lock on a lockfile and exits the program if an error occurs
func (s *Single) Lock() (err error) {
	// open/create lock file
	f, err := os.OpenFile(s.Lockfile(), os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return
	}

	s.file = f
	// set the lock type to F_WRLCK, therefor the file has to be opened writable
	flock := syscall.Flock_t{
		Type: syscall.F_WRLCK,
		Pid:  int32(os.Getpid()),
	}

	if s.pidpath != "" { // error if not root
		_ = ioutil.WriteFile(s.Pidfile(), []byte(strconv.Itoa(int(flock.Pid))), 0644)
	}

	// try to obtain an exclusive lock - FcntlFlock seems to be the portable *ix way
	if err = syscall.FcntlFlock(s.file.Fd(), syscall.F_SETLK, &flock); err != nil {
		return ErrAlreadyRunning
	}

	return
}

// Unlock releases the lock, closes and removes the lockfile
func (s *Single) Unlock() (err error) {
	// set the lock type to F_UNLCK
	flock := syscall.Flock_t{
		Type: syscall.F_UNLCK,
		Pid:  int32(os.Getpid()),
	}

	if err = syscall.FcntlFlock(s.file.Fd(), syscall.F_SETLK, &flock); err != nil {
		return fmt.Errorf("failed to unlock the lock file: %w", err)
	}

	if err = s.file.Close(); err != nil {
		return fmt.Errorf("failed to close the lock file: %w", err)
	}

	if err = os.Remove(s.Lockfile()); err != nil {
		return fmt.Errorf("failed to remove the lock file: %w", err)
	}

	return
}
