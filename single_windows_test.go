package singleproc_test

import (
	"os"
	"testing"

	"github.com/NovikovRoman/singleproc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSingle(t *testing.T) {
	s, err := singleproc.New("unittest")
	require.NoError(t, err)
	require.NotNil(t, s)

	t.Logf("Lockfile: %s", s.Lockfile())

	err = s.Lock()
	assert.NoError(t, err)

	assert.EqualError(t, checkLock(s), singleproc.ErrAlreadyRunning.Error())

	err = s.Unlock()
	assert.NoError(t, err)
}

func checkLock(s *singleproc.Single) error {
	if err := os.Remove(s.Lockfile()); err != nil && !os.IsNotExist(err) {
		return singleproc.ErrAlreadyRunning
	}

	return nil
}
