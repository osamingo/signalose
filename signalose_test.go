package signalose

import (
	"bytes"
	"errors"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type ErrCloser struct{}

func (ec *ErrCloser) Close() error {
	return errors.New("dummy error")
}

func TestAddCloser(t *testing.T) {

	t.Parallel()

	writer := bytes.NewBuffer(nil)
	Writer = writer

	_, err := AddCloser("", nil)
	assert.Error(t, err)

	_, err = AddCloser("testing", nil)
	assert.Error(t, err)

	sigChan, err := AddCloser("testing", &ErrCloser{}, syscall.SIGUSR2)
	assert.NoError(t, err)

	sigChan <- syscall.SIGUSR2

	time.Sleep(time.Second)

	assert.Equal(t, "signalose: testing is close, err = dummy error\n", writer.String())
}
