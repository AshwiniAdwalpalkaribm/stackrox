package store

import (
	"context"

	"github.com/stackrox/rox/generated/storage"
)

// Store provides storage functionality for alerts.
//
//go:generate mockgen-wrapper
type Store interface {
	Get(ctx context.Context, id string) (*storage.Notifier, bool, error)
	Walk(ctx context.Context, fn func(obj *storage.Notifier) error) error
	GetMany(ctx context.Context, identifiers []string) ([]*storage.Notifier, []int, error)
	Exists(ctx context.Context, id string) (bool, error)
	Upsert(ctx context.Context, obj *storage.Notifier) error
	UpsertMany(ctx context.Context, objs []*storage.Notifier) error
	Delete(ctx context.Context, id string) error
}
