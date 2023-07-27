package createOrderCommandV1

import (
	"context"
	"fmt"

	attribute2 "go.opentelemetry.io/otel/attribute"

	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/es/contracts/store"
	customErrors "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/http/http_errors/custom_errors"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/logger"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/mapper"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/otel/tracing"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/otel/tracing/attribute"

	"github.com/mehdihadeli/go-ecommerce-microservices/internal/services/orders/config"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/services/orders/internal/orders/features/creating_order/v1/dtos"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/services/orders/internal/orders/models/orders/aggregate"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/services/orders/internal/orders/models/orders/value_objects"
)

type CreateOrderHandler struct {
    log logger.Logger
    cfg *config.Config
    // goland can't detect this generic type, but it is ok in vscode
    aggregateStore store.AggregateStore[*aggregate.Order]
}

func NewCreateOrderHandler(log logger.Logger, cfg *config.Config, aggregateStore store.AggregateStore[*aggregate.Order]) *CreateOrderHandler {
    return &CreateOrderHandler{log: log, cfg: cfg, aggregateStore: aggregateStore}
}

func (c *CreateOrderHandler) Handle(ctx context.Context, command *CreateOrder) (*dtos.CreateOrderResponseDto, error) {

    ctx, span := tracing.Tracer.Start(ctx, "CreateOrderHandler.Handle")
    span.SetAttributes(attribute2.String("OrderId", command.OrderId.String()))
    span.SetAttributes(attribute.Object("Command", command))
    defer span.End()

    shopItems, err := mapper.Map[[]*value_objects.ShopItem](command.ShopItems)
    if err != nil {
        return nil, tracing.TraceErrFromSpan(span, customErrors.NewApplicationErrorWrap(err, "[CreateOrderHandler_Handle.Map] error in the mapping shopItems"))
    }

    order, err := aggregate.NewOrder(command.OrderId, shopItems, command.AccountEmail, command.DeliveryAddress, command.DeliveryTime, command.CreatedAt)
    if err != nil {
        return nil, tracing.TraceErrFromSpan(span, customErrors.NewApplicationErrorWrap(err, "[CreateOrderHandler_Handle.NewOrder] error in creating new order"))
    }

    _, err = c.aggregateStore.Store(order, nil, ctx)
    if err != nil {
        return nil, tracing.TraceErrFromSpan(span, customErrors.NewApplicationErrorWrap(err, "[CreateOrderHandler_Handle.Store] error in storing order aggregate"))
    }

    response := &dtos.CreateOrderResponseDto{OrderId: order.Id()}

    span.SetAttributes(attribute.Object("CreateOrderResponseDto", response))

    c.log.Infow(fmt.Sprintf("[CreateOrderHandler.Handle] order with id: {%s} created", command.OrderId), logger.Fields{"ProductId": command.OrderId})

    return response, nil
}
