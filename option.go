package singleproc

// Option configures Single
type Option func(*Single)

// WithLockPath configures the path for the lockfile
func WithLockPath(lockpath string) Option {
	return func(s *Single) {
		s.lockpath = lockpath
	}
}

// WithPidPath configures the path for the lockfile
func WithPidPath(pidpath string) Option {
	return func(s *Single) {
		s.pidpath = pidpath
	}
}
