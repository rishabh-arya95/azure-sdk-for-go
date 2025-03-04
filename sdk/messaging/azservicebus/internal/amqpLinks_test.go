// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package internal

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/test"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/utils"
	"github.com/Azure/go-amqp"
	"github.com/stretchr/testify/require"
)

var retryOptionsOnlyOnce = utils.RetryOptions{
	MaxRetries: 0,
}

type fakeNetError struct {
	temp    bool
	timeout bool
}

func (pe fakeNetError) Timeout() bool   { return pe.timeout }
func (pe fakeNetError) Temporary() bool { return pe.temp }
func (pe fakeNetError) Error() string   { return "Fake but very permanent error" }

func assertFailedLinks(t *testing.T, lwid *LinksWithID, expectedErr error) {
	err := lwid.Sender.Send(context.TODO(), &amqp.Message{
		Data: [][]byte{
			{0},
		},
	})
	require.ErrorIs(t, err, expectedErr)

	_, err = PeekMessages(context.TODO(), lwid.RPC, 0, 1)
	require.ErrorIs(t, err, expectedErr)

	msg, err := lwid.Receiver.Receive(context.TODO())
	require.ErrorIs(t, err, expectedErr)
	require.Nil(t, msg)

}

func assertLinks(t *testing.T, lwid *LinksWithID) {
	err := lwid.Sender.Send(context.TODO(), &amqp.Message{
		Data: [][]byte{
			{0},
		},
	})
	require.NoError(t, err)

	_, err = PeekMessages(context.TODO(), lwid.RPC, 0, 1)
	require.NoError(t, err)

	require.NoError(t, lwid.Receiver.IssueCredit(1))
	msg, err := lwid.Receiver.Receive(context.TODO())
	require.NoError(t, err)
	require.NotNil(t, msg)
}

func TestAMQPLinksBasic(t *testing.T) {
	entityPath, cleanup := test.CreateExpiringQueue(t, nil)
	defer cleanup()

	cs := test.GetConnectionString(t)
	ns, err := NewNamespace(NamespaceWithConnectionString(cs))
	require.NoError(t, err)

	links := NewAMQPLinks(ns, entityPath, func(ctx context.Context, session AMQPSession) (AMQPSenderCloser, AMQPReceiverCloser, error) {
		return newLinksForAMQPLinksTest(entityPath, session)
	})

	lwr, err := links.Get(context.Background())
	require.NoError(t, err)

	assertLinks(t, lwr)

	require.EqualValues(t, entityPath, links.EntityPath())
}

func TestAMQPLinksLive(t *testing.T) {
	// we're not going to use this client for tehse tests.
	entityPath, cleanup := test.CreateExpiringQueue(t, nil)
	defer cleanup()

	cs := test.GetConnectionString(t)
	ns, err := NewNamespace(NamespaceWithConnectionString(cs))
	require.NoError(t, err)

	defer func() { _ = ns.Close(context.Background()) }()

	createLinksCalled := 0

	links := NewAMQPLinks(ns, entityPath, func(ctx context.Context, session AMQPSession) (AMQPSenderCloser, AMQPReceiverCloser, error) {
		createLinksCalled++
		return newLinksForAMQPLinksTest(entityPath, session)
	})

	require.EqualValues(t, 0, createLinksCalled)
	require.NoError(t, links.RecoverIfNeeded(context.Background(), LinkID{}, amqp.ErrConnClosed))
	require.EqualValues(t, 1, createLinksCalled)

	lwr, err := links.Get(context.Background())
	require.NoError(t, err)

	amqpClient, clientRev, err := ns.GetAMQPClientImpl(context.Background())
	require.NoError(t, err)
	require.EqualValues(t, 1, clientRev)
	require.NoError(t, amqpClient.Close())

	// all the links are dead because the connection is dead.
	assertFailedLinks(t, lwr, amqp.ErrConnClosed)

	// now we'll recover, which should recreate everything
	require.NoError(t, links.RecoverIfNeeded(context.Background(), lwr.ID, amqp.ErrConnClosed))
	require.EqualValues(t, 2, createLinksCalled)

	lwr, err = links.Get(context.Background())
	require.NoError(t, err)

	// should work now, connection should be reopened
	assertLinks(t, lwr)

	// cheat a bit and close the links out from under us (but leave them in place)
	actualLinks := links.(*AMQPLinksImpl)
	_ = actualLinks.Sender.Close(context.Background())
	_ = actualLinks.Receiver.Close(context.Background())
	_ = actualLinks.RPCLink.Close(context.Background())

	assertFailedLinks(t, lwr, amqp.ErrLinkClosed)

	lwr, err = links.Get(context.Background())
	require.NoError(t, err)

	require.NoError(t, links.RecoverIfNeeded(context.Background(), lwr.ID, amqp.ErrLinkClosed))
	require.EqualValues(t, 3, createLinksCalled)

	lwr, err = links.Get(context.Background())
	require.NoError(t, err)

	assertLinks(t, lwr)
}

func TestAMQPLinksLiveRecoverLink(t *testing.T) {
	// we're not going to use this client for tehse tests.
	entityPath, cleanup := test.CreateExpiringQueue(t, nil)
	defer cleanup()

	cs := test.GetConnectionString(t)
	ns, err := NewNamespace(NamespaceWithConnectionString(cs))
	require.NoError(t, err)

	defer func() { _ = ns.Close(context.Background()) }()

	createLinksCalled := 0

	links := NewAMQPLinks(ns, entityPath, func(ctx context.Context, session AMQPSession) (AMQPSenderCloser, AMQPReceiverCloser, error) {
		createLinksCalled++
		return newLinksForAMQPLinksTest(entityPath, session)
	})

	require.EqualValues(t, 0, createLinksCalled)
	require.NoError(t, links.RecoverIfNeeded(context.Background(), LinkID{}, amqp.ErrConnClosed))
	require.EqualValues(t, 1, createLinksCalled)

	lwr, err := links.Get(context.Background())
	require.NoError(t, err)

	require.NoError(t, links.RecoverIfNeeded(context.Background(), lwr.ID, amqp.ErrLinkClosed))
	require.EqualValues(t, 2, createLinksCalled)
}

func TestAMQPLinksLiveRace(t *testing.T) {
	entityPath, cleanup := test.CreateExpiringQueue(t, nil)
	defer cleanup()

	cs := test.GetConnectionString(t)
	ns, err := NewNamespace(NamespaceWithConnectionString(cs))
	require.NoError(t, err)

	defer func() { _ = ns.Close(context.Background()) }()

	createLinksCalled := 0

	links := NewAMQPLinks(ns, entityPath, func(ctx context.Context, session AMQPSession) (AMQPSenderCloser, AMQPReceiverCloser, error) {
		createLinksCalled++
		return newLinksForAMQPLinksTest(entityPath, session)
	})

	wg := sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := links.RecoverIfNeeded(context.Background(), LinkID{}, amqp.ErrConnClosed)
			require.NoError(t, err)
		}()
	}

	wg.Wait()

	// TODO: also check that the connection hasn't recycled multiple times.
	require.EqualValues(t, 1, createLinksCalled)
}

func TestAMQPLinksLiveRaceLink(t *testing.T) {
	entityPath, cleanup := test.CreateExpiringQueue(t, nil)
	defer cleanup()

	cs := test.GetConnectionString(t)
	ns, err := NewNamespace(NamespaceWithConnectionString(cs))
	require.NoError(t, err)

	defer func() { _ = ns.Close(context.Background()) }()

	createLinksCalled := 0

	test.EnableStdoutLogging()

	links := NewAMQPLinks(ns, entityPath, func(ctx context.Context, session AMQPSession) (AMQPSenderCloser, AMQPReceiverCloser, error) {
		createLinksCalled++
		return newLinksForAMQPLinksTest(entityPath, session)
	})

	wg := sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := links.RecoverIfNeeded(context.Background(), LinkID{}, &amqp.DetachError{})
			require.NoError(t, err)
		}()
	}

	wg.Wait()

	// TODO: also check that the connection hasn't recycled multiple times.
	require.EqualValues(t, 1, createLinksCalled)
}

func TestAMQPLinksRetry(t *testing.T) {
	entityPath, cleanup := test.CreateExpiringQueue(t, nil)
	defer cleanup()

	cs := test.GetConnectionString(t)
	ns, err := NewNamespace(NamespaceWithConnectionString(cs))
	require.NoError(t, err)

	defer func() { _ = ns.Close(context.Background()) }()

	createLinksCalled := 0

	links := NewAMQPLinks(ns, entityPath, func(ctx context.Context, session AMQPSession) (AMQPSenderCloser, AMQPReceiverCloser, error) {
		createLinksCalled++
		return newLinksForAMQPLinksTest(entityPath, session)
	})

	err = links.Retry(context.Background(), "retryOp", func(ctx context.Context, lwid *LinksWithID, args *utils.RetryFnArgs) error {
		// force recoveries
		return amqp.ErrConnClosed
	}, utils.RetryOptions{
		MaxRetries: 2,
		// note: omitting MaxRetries just to give a sanity check that
		// we do setDefaults() before we run.
		RetryDelay:    time.Millisecond,
		MaxRetryDelay: time.Millisecond,
	})

	require.ErrorIs(t, err, amqp.ErrConnClosed)
	require.EqualValues(t, 3, createLinksCalled)
}

func TestAMQPLinksMultipleWithSameConnection(t *testing.T) {
	entityPath, cleanup := test.CreateExpiringQueue(t, nil)
	defer cleanup()

	cs := test.GetConnectionString(t)
	ns, err := NewNamespace(NamespaceWithConnectionString(cs))
	require.NoError(t, err)

	defer func() { _ = ns.Close(context.Background()) }()

	createLinksCalled := 0

	links := NewAMQPLinks(ns, entityPath, func(ctx context.Context, session AMQPSession) (AMQPSenderCloser, AMQPReceiverCloser, error) {
		createLinksCalled++
		return newLinksForAMQPLinksTest(entityPath, session)
	})

	createLinksCalled2 := 0

	links2 := NewAMQPLinks(ns, entityPath, func(ctx context.Context, session AMQPSession) (AMQPSenderCloser, AMQPReceiverCloser, error) {
		createLinksCalled2++
		return newLinksForAMQPLinksTest(entityPath, session)
	})

	wg := sync.WaitGroup{}

	lwr, err := links.Get(context.Background())
	require.NoError(t, err)
	require.EqualValues(t, 1, createLinksCalled)
	require.EqualValues(t, 1, lwr.ID.Link)

	lwr2, err := links2.Get(context.Background())
	require.NoError(t, err)
	require.EqualValues(t, 1, createLinksCalled2)
	require.EqualValues(t, 1, lwr2.ID.Link)

	wg.Add(1)

	go func() {
		defer wg.Done()
		err = links.RecoverIfNeeded(context.Background(), lwr.ID, &amqp.DetachError{})
		require.NoError(t, err)
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()

		err = links2.RecoverIfNeeded(context.Background(), lwr2.ID, &amqp.DetachError{})
		require.NoError(t, err)
	}()

	wg.Wait()

	// TODO: also check that the connection hasn't recycled multiple times.
	require.EqualValues(t, 2, createLinksCalled)
	require.EqualValues(t, 2, createLinksCalled2)

	_, clientRev, err := ns.GetAMQPClientImpl(context.Background())
	require.NoError(t, err)
	require.EqualValues(t, 1, clientRev)

	recovered, err := ns.Recover(context.Background(), clientRev)
	require.NoError(t, err)
	require.True(t, recovered)

	_, clientRev, err = ns.GetAMQPClientImpl(context.Background())
	require.NoError(t, err)
	require.EqualValues(t, 2, clientRev)

	// now attempt a recover but with an older revision (won't do anything since we've
	// already recovered past that older rev. They should just call Get())
	recovered, err = ns.Recover(context.Background(), clientRev-1)
	require.NoError(t, err)
	require.False(t, recovered)

	_, clientRev, err = ns.GetAMQPClientImpl(context.Background())
	require.NoError(t, err)
	require.EqualValues(t, 2, clientRev)
}

func TestAMQPLinksCloseIfNeeded(t *testing.T) {
	t.Run("fatal", func(t *testing.T) {
		for _, fatalErr := range []error{NewErrNonRetriable(""), context.Canceled, context.DeadlineExceeded} {
			receiver := &FakeAMQPReceiver{}
			sender := &FakeAMQPSender{}
			ns := &FakeNS{}

			links := NewAMQPLinks(ns, "entityPath", func(ctx context.Context, session AMQPSession) (AMQPSenderCloser, AMQPReceiverCloser, error) {
				return sender, receiver, nil
			})

			_, err := links.Get(context.Background())
			require.NoError(t, err)

			rk := links.CloseIfNeeded(context.Background(), fatalErr)
			require.Equal(t, RecoveryKindFatal, rk)
			require.Equal(t, 0, receiver.Closed)
			require.Equal(t, 0, sender.Closed)
			require.Equal(t, 0, ns.CloseCalled)
		}
	})

	t.Run("link", func(t *testing.T) {
		receiver := &FakeAMQPReceiver{}
		sender := &FakeAMQPSender{}
		ns := &FakeNS{}

		links := NewAMQPLinks(ns, "entityPath", func(ctx context.Context, session AMQPSession) (AMQPSenderCloser, AMQPReceiverCloser, error) {
			return sender, receiver, nil
		})

		_, err := links.Get(context.Background())
		require.NoError(t, err)

		rk := links.CloseIfNeeded(context.Background(), amqp.ErrLinkClosed)
		require.Equal(t, RecoveryKindLink, rk)
		require.Equal(t, 1, receiver.Closed)
		require.Equal(t, 1, sender.Closed)
		require.Equal(t, 0, ns.CloseCalled)
	})

	t.Run("conn", func(t *testing.T) {
		receiver := &FakeAMQPReceiver{}
		sender := &FakeAMQPSender{}
		ns := &FakeNS{}

		links := NewAMQPLinks(ns, "entityPath", func(ctx context.Context, session AMQPSession) (AMQPSenderCloser, AMQPReceiverCloser, error) {
			return sender, receiver, nil
		})

		_, err := links.Get(context.Background())
		require.NoError(t, err)

		rk := links.CloseIfNeeded(context.Background(), amqp.ErrConnClosed)
		require.Equal(t, RecoveryKindConn, rk)
		require.Equal(t, 1, receiver.Closed)
		require.Equal(t, 1, sender.Closed)
		require.Equal(t, 1, ns.CloseCalled)
	})

	t.Run("none", func(t *testing.T) {
		receiver := &FakeAMQPReceiver{}
		sender := &FakeAMQPSender{}
		ns := &FakeNS{}

		links := NewAMQPLinks(ns, "entityPath", func(ctx context.Context, session AMQPSession) (AMQPSenderCloser, AMQPReceiverCloser, error) {
			return sender, receiver, nil
		})

		_, err := links.Get(context.Background())
		require.NoError(t, err)

		rk := links.CloseIfNeeded(context.Background(), nil)
		require.Equal(t, RecoveryKindNone, rk)
		require.Equal(t, 0, receiver.Closed)
		require.Equal(t, 0, sender.Closed)
		require.Equal(t, 0, ns.CloseCalled)
	})
}

func TestAMQPLinksRetriesUnit(t *testing.T) {
	tests := []struct {
		Err         error
		Attempts    []int32
		ExpectReset bool
	}{
		// nothing goes wrong, only need the one attempt
		{Err: nil, Attempts: []int32{0}},

		// connection related or unknown failures happen, all attempts exhausted
		{Err: amqp.ErrConnClosed, Attempts: []int32{0, 1, 2, 3}},
		{Err: errors.New("unknown error"), Attempts: []int32{0, 1, 2, 3}},

		// fatal errors don't retry at all.
		{Err: NewErrNonRetriable("non retriable error"), Attempts: []int32{0}},

		// detach error happens - we have slightly special behavior here in that we do a quick
		// retry for attempt '0', to avoid sleeping if the error was stale. This mostly happens
		// in situations like sending, where you might have long times in between sends and your
		// link is closed due to idling.
		{Err: &amqp.DetachError{}, Attempts: []int32{0, 0, 1, 2, 3}, ExpectReset: true},
	}

	for _, testData := range tests {
		var testName string = ""

		if testData.Err != nil {
			testName = testData.Err.Error()
		}

		t.Run(testName, func(t *testing.T) {
			receiver := &FakeAMQPReceiver{}
			sender := &FakeAMQPSender{}
			ns := &FakeNS{}

			links := NewAMQPLinks(ns, "entityPath", func(ctx context.Context, session AMQPSession) (AMQPSenderCloser, AMQPReceiverCloser, error) {
				return sender, receiver, nil
			})

			var attempts []int32

			var logMessages []string
			removeLogging := test.CaptureLogsForTest(&logMessages)
			defer removeLogging()

			err := links.Retry(context.Background(), "test", func(ctx context.Context, lwid *LinksWithID, args *utils.RetryFnArgs) error {
				attempts = append(attempts, args.I)
				return testData.Err
			}, utils.RetryOptions{
				RetryDelay: time.Millisecond,
			})

			require.Equal(t, testData.Err, err)
			require.Equal(t, testData.Attempts, attempts)

			if testData.ExpectReset {
				require.Contains(t, logMessages, fmt.Sprintf("[azsb.Conn] Link was previously detached. Attempting quick reconnect to recover from error: %s", err.Error()))
			} else {
				for _, msg := range logMessages {
					require.NotContains(t, msg, "Link was previously detached")
				}
			}
		})
	}
}

func TestAMQPLinks_Logging(t *testing.T) {
	t.Run("link", func(t *testing.T) {
		receiver := &FakeAMQPReceiver{}
		ns := &FakeNS{}

		links := NewAMQPLinks(ns, "entityPath", func(ctx context.Context, session AMQPSession) (AMQPSenderCloser, AMQPReceiverCloser, error) {
			return nil, receiver, nil
		})

		var messages []string

		cleanup := test.CaptureLogsForTest(&messages)
		defer cleanup()

		err := links.RecoverIfNeeded(context.Background(), LinkID{}, &amqp.DetachError{})
		require.NoError(t, err)

		require.Equal(t, []string{
			"[azsb.Conn] Recovering link for error link detached, reason: *Error(nil)",
			"[azsb.Conn] Recovering link only",
			"[azsb.Conn] Recovered links",
		}, messages)
	})

	t.Run("connection", func(t *testing.T) {
		receiver := &FakeAMQPReceiver{}
		ns := &FakeNS{}

		links := NewAMQPLinks(ns, "entityPath", func(ctx context.Context, session AMQPSession) (AMQPSenderCloser, AMQPReceiverCloser, error) {
			return nil, receiver, nil
		})

		var messages []string

		cleanup := test.CaptureLogsForTest(&messages)
		defer cleanup()

		err := links.RecoverIfNeeded(context.Background(), LinkID{}, amqp.ErrConnClosed)
		require.NoError(t, err)

		require.Equal(t, []string{
			"[azsb.Conn] Recovering link for error amqp: connection closed",
			"[azsb.Conn] Recovering connection (and links)",
			"[azsb.Conn] recreating link: c: true, current:{0 0}, old:{0 0}", "[azsb.Conn] Recovered connection and links",
		}, messages)
	})
}

func newLinksForAMQPLinksTest(entityPath string, session AMQPSession) (AMQPSenderCloser, AMQPReceiverCloser, error) {
	receiveMode := amqp.ModeSecond

	opts := []amqp.LinkOption{
		amqp.LinkSourceAddress(entityPath),
		amqp.LinkReceiverSettle(receiveMode),
		amqp.LinkWithManualCredits(),
		amqp.LinkCredit(1000),
	}

	receiver, err := session.NewReceiver(opts...)

	if err != nil {
		return nil, nil, err
	}

	sender, err := session.NewSender(
		amqp.LinkSenderSettle(amqp.ModeMixed),
		amqp.LinkReceiverSettle(amqp.ModeFirst),
		amqp.LinkTargetAddress(entityPath))

	if err != nil {
		_ = receiver.Close(context.Background())
		return nil, nil, err
	}

	return sender, receiver, nil
}
