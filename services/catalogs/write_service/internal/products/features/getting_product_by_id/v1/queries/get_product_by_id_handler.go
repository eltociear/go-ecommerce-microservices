package getProductByIdQuery

import (
	"context"
	"fmt"
	"net/http"

	attribute2 "go.opentelemetry.io/otel/attribute"

	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/http/http_errors/custom_errors"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/logger"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/mapper"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/otel/tracing"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/otel/tracing/attribute"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/write_service/config"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/write_service/internal/products/contracts/data"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/write_service/internal/products/dto"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/write_service/internal/products/features/getting_product_by_id/v1/dtos"
)

type GetProductByIdHandler struct {
	log    logger.Logger
	cfg    *config.Config
	pgRepo data.ProductRepository
}

func NewGetProductByIdHandler(log logger.Logger, cfg *config.Config, pgRepo data.ProductRepository) *GetProductByIdHandler {
	return &GetProductByIdHandler{log: log, cfg: cfg, pgRepo: pgRepo}
}

func (q *GetProductByIdHandler) Handle(ctx context.Context, query *GetProductById) (*dtos.GetProductByIdResponseDto, error) {
	ctx, span := tracing.Tracer.Start(ctx, "GetProductByIdHandler.Handle")
	span.SetAttributes(attribute.Object("Query", query))
	span.SetAttributes(attribute2.String("ProductId", query.ProductID.String()))
	defer span.End()

	product, err := q.pgRepo.GetProductById(ctx, query.ProductID)
	if err != nil {
		return nil, tracing.TraceErrFromSpan(span, customErrors.NewApplicationErrorWrapWithCode(err, http.StatusNotFound, fmt.Sprintf("[GetProductByIdHandler_Handle.GetProductById] error in getting product with id %s in the repository", query.ProductID.String())))
	}

	productDto, err := mapper.Map[*dto.ProductDto](product)
	if err != nil {
		return nil, tracing.TraceErrFromSpan(span, customErrors.NewApplicationErrorWrap(err, "[GetProductByIdHandler_Handle.Map] error in the mapping product"))
	}

	q.log.Infow(fmt.Sprintf("[GetProductByIdHandler.Handle] product with id: {%s} fetched", query.ProductID), logger.Fields{"ProductId": query.ProductID.String()})

	return &dtos.GetProductByIdResponseDto{Product: productDto}, nil
}