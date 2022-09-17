package bus

import (
	"context"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/logger"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/messaging/bus"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/messaging/consumer"
	"sync"
)

type rabbitMQBus struct {
	consumers []consumer.Consumer
	logger    logger.Logger
}

func NewRabbitMQBus(log logger.Logger, consumers []consumer.Consumer) bus.Bus {
	return &rabbitMQBus{logger: log, consumers: consumers}
}

func (r *rabbitMQBus) Start(ctx context.Context) error {
	for _, c := range r.consumers {
		err := c.Consume(ctx)
		if err != nil {
			return r.Stop(ctx)
		}
	}

	return r.Stop(ctx)
}

func (r *rabbitMQBus) Stop(ctx context.Context) error {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(r.consumers))

	for _, c := range r.consumers {
		go func(c consumer.Consumer) {
			defer waitGroup.Done()

			err := c.UnConsume(ctx)
			if err != nil {
				r.logger.Error("error in the unconsuming")
			}
		}(c)
	}
	waitGroup.Wait()

	return nil
}