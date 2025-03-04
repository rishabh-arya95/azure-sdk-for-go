// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azservicebus

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"sync"
	"sync/atomic"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/log"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/utils"
)

// Client provides methods to create Sender and Receiver
// instances to send and receive messages from Service Bus.
type Client struct {
	// NOTE: values need to be 64-bit aligned. Simplest way to make sure this happens
	// is just to make it the first value in the struct
	// See:
	//   Godoc: https://pkg.go.dev/sync/atomic#pkg-note-BUG
	//   PR: https://github.com/Azure/azure-sdk-for-go/pull/16847
	linkCounter uint64

	linksMu   *sync.Mutex
	links     map[uint64]internal.Closeable
	creds     clientCreds
	namespace interface {
		// used internally by `Client`
		internal.NamespaceWithNewAMQPLinks
		// for child clients
		internal.NamespaceForAMQPLinks
	}
	retryOptions RetryOptions
}

// ClientOptions contains options for the `NewClient` and `NewClientFromConnectionString`
// functions.
type ClientOptions struct {
	// TLSConfig configures a client with a custom *tls.Config.
	TLSConfig *tls.Config

	// Application ID that will be passed to the namespace.
	ApplicationID string

	// NewWebSocketConn is a function that can create a net.Conn for use with websockets.
	// For an example, see ExampleNewClient_usingWebsockets() function in example_client_test.go.
	NewWebSocketConn func(ctx context.Context, args NewWebSocketConnArgs) (net.Conn, error)

	// RetryOptions controls how often operations are retried from this client and any
	// Receivers and Senders created from this client.
	RetryOptions RetryOptions
}

// RetryOptions controls how often operations are retried from this client and any
// Receivers and Senders created from this client.
type RetryOptions = utils.RetryOptions

// NewWebSocketConnArgs are passed to your web socket creation function (ClientOptions.NewWebSocketConn)
type NewWebSocketConnArgs = internal.NewWebSocketConnArgs

// NewClient creates a new Client for a Service Bus namespace, using a TokenCredential.
// A Client allows you create receivers (for queues or subscriptions) and senders (for queues and topics).
// fullyQualifiedNamespace is the Service Bus namespace name (ex: myservicebus.servicebus.windows.net)
// credential is one of the credentials in the `github.com/Azure/azure-sdk-for-go/sdk/azidentity` package.
func NewClient(fullyQualifiedNamespace string, credential azcore.TokenCredential, options *ClientOptions) (*Client, error) {
	if fullyQualifiedNamespace == "" {
		return nil, errors.New("fullyQualifiedNamespace must not be empty")
	}

	if credential == nil {
		return nil, errors.New("credential was nil")
	}

	return newClientImpl(clientCreds{
		credential:              credential,
		fullyQualifiedNamespace: fullyQualifiedNamespace,
	}, options)
}

// NewClientFromConnectionString creates a new Client for a Service Bus namespace using a connection string.
// A Client allows you create receivers (for queues or subscriptions) and senders (for queues and topics).
// connectionString is a Service Bus connection string for the namespace or for an entity.
func NewClientFromConnectionString(connectionString string, options *ClientOptions) (*Client, error) {
	if connectionString == "" {
		return nil, errors.New("connectionString must not be empty")
	}

	return newClientImpl(clientCreds{
		connectionString: connectionString,
	}, options)
}

// Next overloads (ie, credential sticks with the client)
// func NewClientWithNamedKeyCredential(fullyQualifiedNamespace string, credential azcore.TokenCredential, options *ClientOptions) (*Client, error) {
// }

type clientCreds struct {
	connectionString string

	// the Service Bus namespace name (ex: myservicebus.servicebus.windows.net)
	fullyQualifiedNamespace string
	credential              azcore.TokenCredential
}

func newClientImpl(creds clientCreds, options *ClientOptions) (*Client, error) {
	client := &Client{
		linksMu: &sync.Mutex{},
		creds:   creds,
		links:   map[uint64]internal.Closeable{},
	}

	var err error
	var nsOptions []internal.NamespaceOption

	if client.creds.connectionString != "" {
		nsOptions = append(nsOptions, internal.NamespaceWithConnectionString(client.creds.connectionString))
	} else if client.creds.credential != nil {
		option := internal.NamespaceWithTokenCredential(
			client.creds.fullyQualifiedNamespace,
			client.creds.credential)

		nsOptions = append(nsOptions, option)
	}

	if options != nil {
		if options.TLSConfig != nil {
			nsOptions = append(nsOptions, internal.NamespaceWithTLSConfig(options.TLSConfig))
		}

		if options.NewWebSocketConn != nil {
			nsOptions = append(nsOptions, internal.NamespaceWithWebSocket(options.NewWebSocketConn))
		}

		if options.ApplicationID != "" {
			nsOptions = append(nsOptions, internal.NamespaceWithUserAgent(options.ApplicationID))
		}

		nsOptions = append(nsOptions, internal.NamespaceWithRetryOptions((utils.RetryOptions)(options.RetryOptions)))
	}

	client.namespace, err = internal.NewNamespace(nsOptions...)
	return client, err
}

// NewReceiver creates a Receiver for a queue. A receiver allows you to receive messages.
func (client *Client) NewReceiverForQueue(queueName string, options *ReceiverOptions) (*Receiver, error) {
	id, cleanupOnClose := client.getCleanupForCloseable()
	receiver, err := newReceiver(newReceiverArgs{
		cleanupOnClose: cleanupOnClose,
		ns:             client.namespace,
		entity:         entity{Queue: queueName},
	}, options)

	if err != nil {
		return nil, err
	}

	client.addCloseable(id, receiver)
	return receiver, nil
}

// NewReceiver creates a Receiver for a subscription. A receiver allows you to receive messages.
func (client *Client) NewReceiverForSubscription(topicName string, subscriptionName string, options *ReceiverOptions) (*Receiver, error) {
	id, cleanupOnClose := client.getCleanupForCloseable()
	receiver, err := newReceiver(newReceiverArgs{
		cleanupOnClose: cleanupOnClose,
		ns:             client.namespace,
		entity:         entity{Topic: topicName, Subscription: subscriptionName},
	}, options)

	if err != nil {
		return nil, err
	}

	client.addCloseable(id, receiver)
	return receiver, nil
}

type NewSenderOptions struct {
	// For future expansion
}

// NewSender creates a Sender, which allows you to send messages or schedule messages.
func (client *Client) NewSender(queueOrTopic string, options *NewSenderOptions) (*Sender, error) {
	id, cleanupOnClose := client.getCleanupForCloseable()
	sender, err := newSender(newSenderArgs{
		ns:             client.namespace,
		queueOrTopic:   queueOrTopic,
		cleanupOnClose: cleanupOnClose,
	}, client.retryOptions)

	if err != nil {
		return nil, err
	}

	client.addCloseable(id, sender)
	return sender, nil
}

// AcceptSessionForQueue accepts a session from a queue with a specific session ID.
// NOTE: this receiver is initialized immediately, not lazily.
func (client *Client) AcceptSessionForQueue(ctx context.Context, queueName string, sessionID string, options *SessionReceiverOptions) (*SessionReceiver, error) {
	id, cleanupOnClose := client.getCleanupForCloseable()
	sessionReceiver, err := newSessionReceiver(
		ctx,
		&sessionID,
		client.namespace,
		entity{Queue: queueName},
		cleanupOnClose,
		toReceiverOptions(options))

	if err != nil {
		return nil, err
	}

	if err := sessionReceiver.init(ctx); err != nil {
		return nil, err
	}

	client.addCloseable(id, sessionReceiver)
	return sessionReceiver, nil
}

// AcceptSessionForSubscription accepts a session from a subscription with a specific session ID.
// NOTE: this receiver is initialized immediately, not lazily.
func (client *Client) AcceptSessionForSubscription(ctx context.Context, topicName string, subscriptionName string, sessionID string, options *SessionReceiverOptions) (*SessionReceiver, error) {
	id, cleanupOnClose := client.getCleanupForCloseable()
	sessionReceiver, err := newSessionReceiver(
		ctx,
		&sessionID,
		client.namespace,
		entity{Topic: topicName, Subscription: subscriptionName},
		cleanupOnClose,
		toReceiverOptions(options))

	if err != nil {
		return nil, err
	}

	if err := sessionReceiver.init(ctx); err != nil {
		return nil, err
	}

	client.addCloseable(id, sessionReceiver)
	return sessionReceiver, nil
}

// AcceptNextSessionForQueue accepts the next available session from a queue.
// NOTE: this receiver is initialized immediately, not lazily.
func (client *Client) AcceptNextSessionForQueue(ctx context.Context, queueName string, options *SessionReceiverOptions) (*SessionReceiver, error) {
	id, cleanupOnClose := client.getCleanupForCloseable()
	sessionReceiver, err := newSessionReceiver(
		ctx,
		nil,
		client.namespace,
		entity{Queue: queueName},
		cleanupOnClose,
		toReceiverOptions(options))

	if err != nil {
		return nil, err
	}

	if err := sessionReceiver.init(ctx); err != nil {
		return nil, err
	}

	client.addCloseable(id, sessionReceiver)
	return sessionReceiver, nil
}

// AcceptNextSessionForSubscription accepts the next available session from a subscription.
// NOTE: this receiver is initialized immediately, not lazily.
func (client *Client) AcceptNextSessionForSubscription(ctx context.Context, topicName string, subscriptionName string, options *SessionReceiverOptions) (*SessionReceiver, error) {
	id, cleanupOnClose := client.getCleanupForCloseable()
	sessionReceiver, err := newSessionReceiver(
		ctx,
		nil,
		client.namespace,
		entity{Topic: topicName, Subscription: subscriptionName},
		cleanupOnClose,
		toReceiverOptions(options))

	if err != nil {
		return nil, err
	}

	if err := sessionReceiver.init(ctx); err != nil {
		return nil, err
	}

	client.addCloseable(id, sessionReceiver)
	return sessionReceiver, nil
}

// Close closes the current connection Service Bus as well as any Senders or Receivers created
// using this client.
func (client *Client) Close(ctx context.Context) error {
	var lastError error

	var links []internal.Closeable

	client.linksMu.Lock()

	for _, link := range client.links {
		links = append(links, link)
	}

	client.linksMu.Unlock()

	for _, link := range links {
		if err := link.Close(ctx); err != nil {
			log.Writef(internal.EventConn, "Failed to close links (error might be cached): %s", err.Error())
			lastError = err
		}
	}

	if lastError != nil {
		return fmt.Errorf("errors while closing links: %w", lastError)
	}
	return nil
}

func (client *Client) addCloseable(id uint64, closeable internal.Closeable) {
	client.linksMu.Lock()
	client.links[id] = closeable
	client.linksMu.Unlock()
}

func (client *Client) getCleanupForCloseable() (uint64, func()) {
	id := atomic.AddUint64(&client.linkCounter, 1)

	return id, func() {
		client.linksMu.Lock()
		delete(client.links, id)
		client.linksMu.Unlock()
	}
}
