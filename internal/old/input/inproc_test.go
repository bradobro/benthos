package input_test

import (
	"testing"
	"time"

	"github.com/Jeffail/benthos/v3/internal/component/metrics"
	"github.com/Jeffail/benthos/v3/internal/log"
	"github.com/Jeffail/benthos/v3/internal/manager"
	"github.com/Jeffail/benthos/v3/internal/message"
	"github.com/Jeffail/benthos/v3/internal/old/input"
	"github.com/stretchr/testify/require"
)

//------------------------------------------------------------------------------

func TestInprocDryRun(t *testing.T) {
	t.Parallel()

	mgr, err := manager.NewV2(manager.NewResourceConfig(), nil, log.Noop(), metrics.Noop())
	if err != nil {
		t.Fatal(err)
	}

	mgr.SetPipe("foo", make(chan message.Transaction))

	conf := input.NewConfig()
	conf.Inproc = "foo"

	ip, err := input.NewInproc(conf, mgr, log.Noop(), metrics.Noop())
	require.NoError(t, err)

	<-time.After(time.Millisecond * 100)

	ip.CloseAsync()
	if err = ip.WaitForClose(time.Second); err != nil {
		t.Error(err)
	}
}

func TestInprocDryRunNoConn(t *testing.T) {
	t.Parallel()

	mgr, err := manager.NewV2(manager.NewResourceConfig(), nil, log.Noop(), metrics.Noop())
	if err != nil {
		t.Fatal(err)
	}

	conf := input.NewConfig()
	conf.Inproc = "foo"

	ip, err := input.NewInproc(conf, mgr, log.Noop(), metrics.Noop())
	require.NoError(t, err)

	<-time.After(time.Millisecond * 100)

	ip.CloseAsync()
	if err = ip.WaitForClose(time.Second); err != nil {
		t.Error(err)
	}
}

//------------------------------------------------------------------------------