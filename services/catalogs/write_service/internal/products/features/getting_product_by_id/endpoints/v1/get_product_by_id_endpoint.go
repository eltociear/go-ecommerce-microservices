package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/mediatr"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/tracing"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/utils"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/write_service/internal/products/contracts"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/write_service/internal/products/features/getting_product_by_id"
	shared_configurations "github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/write_service/internal/shared/configurations"
	"net/http"
)

type getProductByIdEndpoint struct {
	mediator          *mediatr.Mediator
	productRepository contracts.ProductRepository
	productsGroup     *echo.Group
	infrastructure    *shared_configurations.Infrastructure
}

func NewGetProductByIdEndpoint(infra *shared_configurations.Infrastructure, mediator *mediatr.Mediator, productsGroup *echo.Group, productRepository contracts.ProductRepository) *getProductByIdEndpoint {
	return &getProductByIdEndpoint{mediator: mediator, productRepository: productRepository, productsGroup: productsGroup, infrastructure: infra}
}

func (ep *getProductByIdEndpoint) MapRoute() {
	ep.productsGroup.GET("/:id", ep.getProductByID())
}

// GetProductByID
// @Tags Products
// @Summary Get product
// @Description Get product by id
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} getting_product_by_id.GetProductByIdResponseDto
// @Router /products/{id} [get]
func (ep *getProductByIdEndpoint) getProductByID() echo.HandlerFunc {
	return func(c echo.Context) error {

		ep.infrastructure.Metrics.GetProductByIdHttpRequests.Inc()
		ctx, span := tracing.StartHttpServerTracerSpan(c, "productsHandlers.getProductByID")
		defer span.Finish()

		request := &getting_product_by_id.GetProductByIdRequestDto{}
		if err := c.Bind(request); err != nil {
			ep.infrastructure.Log.WarnMsg("Bind", err)
			tracing.TraceErr(span, err)
			return err
		}

		query := getting_product_by_id.NewGetProductById(request.ProductId)

		if err := ep.infrastructure.Validator.StructCtx(ctx, query); err != nil {
			ep.infrastructure.Log.WarnMsg("validate", err)
			tracing.TraceErr(span, err)
			return err
		}

		queryResult, err := ep.mediator.Send(ctx, query)

		if err != nil {
			ep.infrastructure.Log.WarnMsg("GetProductById", err)
			tracing.TraceErr(span, err)
			return err
		}

		response, ok := queryResult.(*getting_product_by_id.GetProductByIdResponseDto)
		err = utils.CheckType(ok)
		if err != nil {
			tracing.TraceErr(span, err)
			return err
		}

		return c.JSON(http.StatusOK, response)
	}
}