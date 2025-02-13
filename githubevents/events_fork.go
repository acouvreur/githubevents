// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"context"
	"fmt"
	"github.com/google/go-github/v69/github"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
)

// Actions are used to identify registered callbacks.
const (
	// ForkEvent is the event name of github.ForkEvent's
	ForkEvent = "fork"

	// ForkEventAnyAction is used to identify callbacks
	// listening to all events of type github.ForkEvent
	ForkEventAnyAction = "*"
)

// ForkEventHandleFunc represents a callback function triggered on github.ForkEvent's.
// 'deliveryID' (type: string) is the unique webhook delivery ID.
// 'eventName' (type: string) is the name of the event.
// 'event' (type: *github.ForkEvent) is the webhook payload.
type ForkEventHandleFunc func(ctx context.Context, deliveryID string, eventName string, event *github.ForkEvent) error

// OnForkEventAny registers callbacks listening to any events of type github.ForkEvent
//
// This function appends the callbacks passed as arguments to already existing ones.
// If already existing callbacks are to be overwritten, SetOnForkEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#fork
func (g *EventHandler) OnForkEventAny(callbacks ...ForkEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onForkEvent == nil {
		g.onForkEvent = make(map[string][]ForkEventHandleFunc)
	}
	g.onForkEvent[ForkEventAnyAction] = append(
		g.onForkEvent[ForkEventAnyAction],
		callbacks...,
	)
}

// SetOnForkEventAny registers callbacks listening to any events of type github.ForkEvent
// and overwrites already registered callbacks.
//
// This function overwrites all previously registered callbacks.
// If already registered callbacks are not to be overwritten, OnForkEventAny must be used.
//
// Callbacks are executed in parallel. This function blocks until all callbacks executed in parallel have returned,
// then returns the first non-nil error (if any) from them. If OnError callbacks have been set, they will be called when an error occurs.
//
// Reference: https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#fork
func (g *EventHandler) SetOnForkEventAny(callbacks ...ForkEventHandleFunc) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if callbacks == nil || len(callbacks) == 0 {
		panic("callbacks is nil or empty")
	}
	if g.onForkEvent == nil {
		g.onForkEvent = make(map[string][]ForkEventHandleFunc)
	}
	g.onForkEvent[ForkEventAnyAction] = callbacks
}

func (g *EventHandler) handleForkEventAny(ctx context.Context, deliveryID string, eventName string, event *github.ForkEvent) error {
	ctx, span := g.Tracer.Start(ctx, "handleForkEventAny", trace.WithAttributes(
		attribute.String("deliveryID", deliveryID),
		attribute.String("event", eventName),
	))
	defer span.End()
	if event == nil {
		err := fmt.Errorf("event was empty or nil")
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	if _, ok := g.onForkEvent[ForkEventAnyAction]; !ok {
		return nil
	}
	eg := new(errgroup.Group)
	for _, h := range g.onForkEvent[ForkEventAnyAction] {
		handle := h
		eg.Go(func() error {
			err := handle(ctx, deliveryID, eventName, event)
			if err != nil {
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

// ForkEvent handles github.ForkEvent.
//
// Callbacks are executed in the following order:
//
// 1) All callbacks registered with OnBeforeAny are executed in parallel.
// 2) All callbacks registered with OnForkEvent... are executed in parallel in case the Event has actions.
// 3) All callbacks registered with OnAfterAny are executed in parallel.
//
// on any error all callbacks registered with OnError are executed in parallel.
func (g *EventHandler) ForkEvent(ctx context.Context, deliveryID string, eventName string, event *github.ForkEvent) error {
	ctx, span := g.Tracer.Start(ctx, "ForkEvent", trace.WithAttributes(
		attribute.String("deliveryID", deliveryID),
		attribute.String("event", eventName),
	))
	defer span.End()

	if event == nil {
		err := fmt.Errorf("event action was empty or nil")
		span.SetStatus(codes.Error, err.Error())
		return err
	}

	err := g.handleBeforeAny(ctx, deliveryID, eventName, event)
	if err != nil {
		return g.handleError(ctx, deliveryID, eventName, event, err)
	}

	err = g.handleForkEventAny(ctx, deliveryID, eventName, event)
	if err != nil {
		return g.handleError(ctx, deliveryID, eventName, event, err)
	}

	err = g.handleAfterAny(ctx, deliveryID, eventName, event)
	if err != nil {
		return g.handleError(ctx, deliveryID, eventName, event, err)
	}
	return nil
}
