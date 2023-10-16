package interfaces

import "context"

type Publisher interface {
	PublishBinlog(ctx context.Context, payload Payload) error
	PublishOutbox(ctx context.Context, outbox Outbox) error
}

type BackendPublisher interface {
	PublishBinlog(ctx context.Context, event Event, payload SendPayload) (string, error)
	PublishOutbox(ctx context.Context, outbox Outbox) (string, error)
	IsTarget(ctx context.Context, payload SendPayload) bool
}
