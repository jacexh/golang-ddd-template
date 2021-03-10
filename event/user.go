package event

import (
	"context"

	"github.com/jacexh/golang-ddd-template/domain/user"
	"github.com/jacexh/golang-ddd-template/logger"
	"github.com/jacexh/golang-ddd-template/trace"
	"go.uber.org/zap"
)

type userEventPublisher struct {
	subscribers map[string][]user.Subscriber
	concurrency chan struct{}
}

var Publisher user.EventPublisher = (*userEventPublisher)(nil)

func BuildUserEventPublisher() user.EventPublisher {
	Publisher = &userEventPublisher{
		subscribers: make(map[string][]user.Subscriber),
		concurrency: make(chan struct{}, 3),
	}
	return Publisher
}

func (p *userEventPublisher) Subscribe(t string, sub user.Subscriber) {
	if _, exists := p.subscribers[t]; !exists {
		p.subscribers[t] = []user.Subscriber{sub}
		return
	}
	p.subscribers[t] = append(p.subscribers[t], sub)
}

func (p *userEventPublisher) Publish(ctx context.Context, event user.Event) {
	if subs, exists := p.subscribers[event.Type()]; exists {
		p.concurrency <- struct{}{}
		go func(ctx context.Context, event user.Event, subs ...user.Subscriber) {
			defer func() { <-p.concurrency }()
			for _, sub := range subs {
				if err := sub.Handle(ctx, event); err != nil {
					logger.Logger.Error("failed to handle event",
						zap.String("event_type", event.Type()),
						trace.MustExtractRequestIndexFromCtxAsField(ctx), zap.Error(err))
				}
			}
		}(ctx, event, subs...)
	}
}
