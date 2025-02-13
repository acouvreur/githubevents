// Copyright 2022 The GithubEvents Authors. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package githubevents

// THIS FILE IS GENERATED - DO NOT EDIT DIRECTLY
// make edits in gen/generate.go

import (
	"context"
	"errors"
	"github.com/google/go-github/v69/github"
	"go.opentelemetry.io/otel/trace/noop"
	"sync"
	"testing"
)

func TestOnProjectV2EventAny(t *testing.T) {
	type args struct {
		callbacks []ProjectV2EventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectV2EventHandleFunc",
			args: args{
				[]ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectV2EventHandleFuncs",
			args: args{
				[]ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectV2EventAny(tt.args.callbacks...)
			if len(g.onProjectV2Event[ProjectV2EventAnyAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectV2Event[ProjectV2EventAnyAction]))
			}
		})
	}
}

func TestSetOnProjectV2EventAny(t *testing.T) {
	type args struct {
		callbacks []ProjectV2EventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectV2EventHandleFunc",
			args: args{
				[]ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectV2EventHandleFuncs",
			args: args{
				[]ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnProjectV2EventAny(func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
				return nil
			})
			g.SetOnProjectV2EventAny(tt.args.callbacks...)
			if len(g.onProjectV2Event[ProjectV2EventAnyAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectV2Event[ProjectV2EventAnyAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectV2EventAny(t *testing.T) {

	action := "*"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectV2Event
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",

				event: &github.ProjectV2Event{Action: &action},

				fail: false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",

				event: &github.ProjectV2Event{Action: &action},

				fail: true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectV2EventAny(func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectV2EventAny(context.Background(), tt.args.deliveryID, tt.args.deliveryID, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("TestHandleProjectV2EventAny() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectEventCreated(t *testing.T) {
	type args struct {
		callbacks []ProjectV2EventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectV2EventHandleFunc",
			args: args{
				callbacks: []ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectV2EventHandleFunc",
			args: args{
				callbacks: []ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventCreated(tt.args.callbacks...)
			if len(g.onProjectV2Event[ProjectEventCreatedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectV2Event[ProjectEventCreatedAction]))
			}
		})
	}
}

func TestSetOnProjectEventCreated(t *testing.T) {
	type args struct {
		callbacks []ProjectV2EventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectV2EventHandleFunc",
			args: args{
				[]ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectV2EventHandleFuncs",
			args: args{
				[]ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnProjectEventCreated(func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
				return nil
			})
			g.SetOnProjectEventCreated(tt.args.callbacks...)
			if len(g.onProjectV2Event[ProjectEventCreatedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectV2Event[ProjectEventCreatedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectEventCreated(t *testing.T) {
	action := ProjectEventCreatedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectV2Event
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventCreated(func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectEventCreated(context.Background(), tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectEventCreated() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectEventEdited(t *testing.T) {
	type args struct {
		callbacks []ProjectV2EventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectV2EventHandleFunc",
			args: args{
				callbacks: []ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectV2EventHandleFunc",
			args: args{
				callbacks: []ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventEdited(tt.args.callbacks...)
			if len(g.onProjectV2Event[ProjectEventEditedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectV2Event[ProjectEventEditedAction]))
			}
		})
	}
}

func TestSetOnProjectEventEdited(t *testing.T) {
	type args struct {
		callbacks []ProjectV2EventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectV2EventHandleFunc",
			args: args{
				[]ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectV2EventHandleFuncs",
			args: args{
				[]ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnProjectEventEdited(func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
				return nil
			})
			g.SetOnProjectEventEdited(tt.args.callbacks...)
			if len(g.onProjectV2Event[ProjectEventEditedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectV2Event[ProjectEventEditedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectEventEdited(t *testing.T) {
	action := ProjectEventEditedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectV2Event
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventEdited(func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectEventEdited(context.Background(), tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectEventEdited() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectEventClosed(t *testing.T) {
	type args struct {
		callbacks []ProjectV2EventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectV2EventHandleFunc",
			args: args{
				callbacks: []ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectV2EventHandleFunc",
			args: args{
				callbacks: []ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventClosed(tt.args.callbacks...)
			if len(g.onProjectV2Event[ProjectEventClosedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectV2Event[ProjectEventClosedAction]))
			}
		})
	}
}

func TestSetOnProjectEventClosed(t *testing.T) {
	type args struct {
		callbacks []ProjectV2EventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectV2EventHandleFunc",
			args: args{
				[]ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectV2EventHandleFuncs",
			args: args{
				[]ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnProjectEventClosed(func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
				return nil
			})
			g.SetOnProjectEventClosed(tt.args.callbacks...)
			if len(g.onProjectV2Event[ProjectEventClosedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectV2Event[ProjectEventClosedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectEventClosed(t *testing.T) {
	action := ProjectEventClosedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectV2Event
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventClosed(func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectEventClosed(context.Background(), tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectEventClosed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectEventReopened(t *testing.T) {
	type args struct {
		callbacks []ProjectV2EventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectV2EventHandleFunc",
			args: args{
				callbacks: []ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectV2EventHandleFunc",
			args: args{
				callbacks: []ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventReopened(tt.args.callbacks...)
			if len(g.onProjectV2Event[ProjectEventReopenedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectV2Event[ProjectEventReopenedAction]))
			}
		})
	}
}

func TestSetOnProjectEventReopened(t *testing.T) {
	type args struct {
		callbacks []ProjectV2EventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectV2EventHandleFunc",
			args: args{
				[]ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectV2EventHandleFuncs",
			args: args{
				[]ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnProjectEventReopened(func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
				return nil
			})
			g.SetOnProjectEventReopened(tt.args.callbacks...)
			if len(g.onProjectV2Event[ProjectEventReopenedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectV2Event[ProjectEventReopenedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectEventReopened(t *testing.T) {
	action := ProjectEventReopenedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectV2Event
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventReopened(func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectEventReopened(context.Background(), tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectEventReopened() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnProjectEventDeleted(t *testing.T) {
	type args struct {
		callbacks []ProjectV2EventHandleFunc
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "must add single ProjectV2EventHandleFunc",
			args: args{
				callbacks: []ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
		},
		{
			name: "must add multiple ProjectV2EventHandleFunc",
			args: args{
				callbacks: []ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventDeleted(tt.args.callbacks...)
			if len(g.onProjectV2Event[ProjectEventDeletedAction]) == 0 {
				t.Errorf("failed to add callbacks, got %d", len(g.onProjectV2Event[ProjectEventDeletedAction]))
			}
		})
	}
}

func TestSetOnProjectEventDeleted(t *testing.T) {
	type args struct {
		callbacks []ProjectV2EventHandleFunc
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "must add single ProjectV2EventHandleFunc",
			args: args{
				[]ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
			want: 1,
		},
		{
			name: "must add multiple ProjectV2EventHandleFuncs",
			args: args{
				[]ProjectV2EventHandleFunc{
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
					func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
						return nil
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			// add callbacks to be overwritten
			g.SetOnProjectEventDeleted(func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
				return nil
			})
			g.SetOnProjectEventDeleted(tt.args.callbacks...)
			if len(g.onProjectV2Event[ProjectEventDeletedAction]) != tt.want {
				t.Errorf("failed to add callbacks, got %d, want %d", len(g.onProjectV2Event[ProjectEventDeletedAction]), tt.want)
			}
		})
	}
}

func TestHandleProjectEventDeleted(t *testing.T) {
	action := ProjectEventDeletedAction

	emptyAction := ""
	fakeAction := "doesntexist"

	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectV2Event
		fail       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "must pass",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: &action},
				fail:       false,
			},
			wantErr: false,
		},
		{
			name: "must fail with error",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: &action},
				fail:       true,
			},
			wantErr: true,
		},
		{
			name: "must fail event nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      nil,
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail with wrong action",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: &fakeAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action empty",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: &emptyAction},
				fail:       false,
			},
			wantErr: true,
		},
		{
			name: "must fail event action nil",
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: nil},
				fail:       false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New("fake")
			g.OnProjectEventDeleted(func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
				if tt.args.fail {
					return errors.New("fake error")
				}
				return nil
			})
			if err := g.handleProjectEventDeleted(context.Background(), tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("handleProjectEventDeleted() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProjectV2Event(t *testing.T) {
	type fields struct {
		handler *EventHandler
	}
	type args struct {
		deliveryID string
		eventName  string
		event      *github.ProjectV2Event
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "must trigger ProjectV2EventAny with unknown event action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2Event: map[string][]ProjectV2EventHandleFunc{
						ProjectV2EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Log("onAny action called")
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  ProjectV2Event,

				event: &github.ProjectV2Event{Action: ptrString("unknown")},
			},
			wantErr: false,
		},

		{
			name: "must trigger ProjectEventCreated",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2Event: map[string][]ProjectV2EventHandleFunc{
						ProjectV2EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventCreatedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Logf("%s action called", ProjectEventCreatedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: ptrString(ProjectEventCreatedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectEventCreated with empty action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2Event: map[string][]ProjectV2EventHandleFunc{
						ProjectV2EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventCreatedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Logf("%s action called", ProjectEventCreatedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectEventCreated with nil action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2Event: map[string][]ProjectV2EventHandleFunc{
						ProjectV2EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventCreatedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Logf("%s action called", ProjectEventCreatedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger ProjectEventEdited",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2Event: map[string][]ProjectV2EventHandleFunc{
						ProjectV2EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventEditedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Logf("%s action called", ProjectEventEditedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: ptrString(ProjectEventEditedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectEventEdited with empty action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2Event: map[string][]ProjectV2EventHandleFunc{
						ProjectV2EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventEditedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Logf("%s action called", ProjectEventEditedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectEventEdited with nil action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2Event: map[string][]ProjectV2EventHandleFunc{
						ProjectV2EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventEditedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Logf("%s action called", ProjectEventEditedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger ProjectEventClosed",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2Event: map[string][]ProjectV2EventHandleFunc{
						ProjectV2EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventClosedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Logf("%s action called", ProjectEventClosedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: ptrString(ProjectEventClosedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectEventClosed with empty action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2Event: map[string][]ProjectV2EventHandleFunc{
						ProjectV2EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventClosedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Logf("%s action called", ProjectEventClosedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectEventClosed with nil action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2Event: map[string][]ProjectV2EventHandleFunc{
						ProjectV2EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventClosedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Logf("%s action called", ProjectEventClosedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger ProjectEventReopened",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2Event: map[string][]ProjectV2EventHandleFunc{
						ProjectV2EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventReopenedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Logf("%s action called", ProjectEventReopenedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: ptrString(ProjectEventReopenedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectEventReopened with empty action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2Event: map[string][]ProjectV2EventHandleFunc{
						ProjectV2EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventReopenedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Logf("%s action called", ProjectEventReopenedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectEventReopened with nil action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2Event: map[string][]ProjectV2EventHandleFunc{
						ProjectV2EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventReopenedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Logf("%s action called", ProjectEventReopenedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: nil},
			},
			wantErr: true,
		},

		{
			name: "must trigger ProjectEventDeleted",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2Event: map[string][]ProjectV2EventHandleFunc{
						ProjectV2EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventDeletedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Logf("%s action called", ProjectEventDeletedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: ptrString(ProjectEventDeletedAction)},
			},
			wantErr: false,
		},
		{
			name: "must fail ProjectEventDeleted with empty action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2Event: map[string][]ProjectV2EventHandleFunc{
						ProjectV2EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventDeletedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Logf("%s action called", ProjectEventDeletedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: ptrString("")},
			},
			wantErr: true,
		},
		{
			name: "must fail ProjectEventDeleted with nil action",
			fields: fields{
				handler: &EventHandler{
					WebhookSecret: "fake",
					onBeforeAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onBeforeAny called")
								return nil
							},
						},
					},
					onAfterAny: map[string][]EventHandleFunc{
						EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event any) error {
								t.Log("onAfterAny called")
								return nil
							},
						},
					},
					onProjectV2Event: map[string][]ProjectV2EventHandleFunc{
						ProjectV2EventAnyAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Log("onAny action called")
								return nil
							},
						},
						ProjectEventDeletedAction: {
							func(ctx context.Context, deliveryID string, eventName string, event *github.ProjectV2Event) error {
								t.Logf("%s action called", ProjectEventDeletedAction)
								return nil
							},
						},
					},
					Tracer: noop.Tracer{},
				},
			},
			args: args{
				deliveryID: "42",
				eventName:  "project_v2",
				event:      &github.ProjectV2Event{Action: nil},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &EventHandler{
				WebhookSecret: "fake",
				mu:            sync.RWMutex{},
				Tracer:        noop.Tracer{},
			}
			if err := g.ProjectV2Event(context.Background(), tt.args.deliveryID, tt.args.eventName, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("ProjectV2Event() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
