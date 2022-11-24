package endpoints

import (
	"context"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/test/utils"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/orders/internal/orders/delivery"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/orders/internal/shared/test_fixtures/e2e"
)

// we could also run the server on docker and then send rest call to the api
func Test_Order_By_Id_E2E(t *testing.T) {
	testUtils.SkipCI(t)
	fixture := e2e.NewE2ETestFixture()

	e := NewGetOrderByIdEndpoint(delivery.NewOrderEndpointBase(fixture.InfrastructureConfigurations, fixture.V1.OrdersGroup, fixture.Bus, fixture.OrdersMetrics))
	e.MapRoute()

	fixture.Run()
	defer fixture.Cleanup()

	// create httpexpect instance
	expect := httpexpect.New(t, fixture.HttpServer.URL)

	expect.GET("/api/v1/orders/{id}").
		WithPath("id", "c8018f1e-787b-4d5e-98fd-4b4e072d56b2").
		WithContext(context.Background()).
		Expect().
		Status(http.StatusOK)
}