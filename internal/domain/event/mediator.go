package event

import (
	"context"

	"{{.Module}}/internal/logger"
	"{{.Module}}/internal/trace"
	"go.uber.org/zap"
)

type (
	// DomainEventType 事件类型
	DomainEventType string

	// DomainEvent 领域事件
	DomainEvent interface {
		Type() DomainEventType
	}

	// DomainEventMediator 领域事件调度者
	DomainEventMediator interface {
		Publish(ctx context.Context, event DomainEvent)
		Subscribe(DomainEventType, DomainEventSubscriber)
	}

	// DomainEventSubscriber 领域事件订阅方
	DomainEventSubscriber interface {
		Handle(ctx context.Context, event DomainEvent)
	}

	mediator struct {
		subscribers map[DomainEventType][]DomainEventSubscriber
		concurrency chan struct{}
	}
)

var defaultMediator DomainEventMediator = (*mediator)(nil)

func newMediator(n int) *mediator {
	return &mediator{
		subscribers: make(map[DomainEventType][]DomainEventSubscriber),
		concurrency: make(chan struct{}, n),
	}
}

func (m *mediator) Publish(ctx context.Context, event DomainEvent) {
	if subs, exists := m.subscribers[event.Type()]; exists {
		m.concurrency <- struct{}{}
		go func(ctx context.Context, event DomainEvent, subs ...DomainEventSubscriber) {
			defer func() { <-m.concurrency }()

			select {
			case <-ctx.Done():
				logger.Logger.Error("failed to handle current domain event", zap.String("event_type", string(event.Type())),
					zap.Any("event_details", event), zap.Error(ctx.Err()), trace.MustExtractRequestIndexFromCtxAsField(ctx))
			default:
				for _, sub := range subs {
					sub.Handle(ctx, event)
				}
			}
		}(ctx, event, subs...)
	}
}

func (m *mediator) Subscribe(t DomainEventType, sub DomainEventSubscriber) {
	if _, exists := m.subscribers[t]; !exists {
		m.subscribers[t] = []DomainEventSubscriber{sub}
		return
	}
	m.subscribers[t] = append(m.subscribers[t], sub)
}

func Publish(ctx context.Context, event DomainEvent) {
	defaultMediator.Publish(ctx, event)
}

func Subscribe(t DomainEventType, sub DomainEventSubscriber) {
	defaultMediator.Subscribe(t, sub)
}

func ResetMediator(m DomainEventMediator) {
	defaultMediator = m
}

func init() {
	ResetMediator(newMediator(10))
}
