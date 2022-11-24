package commands

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	uuid "github.com/satori/go.uuid"

	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/write_service/internal/shared/test_fixtures/unit_test"
	"github.com/stretchr/testify/suite"
)

type updateProductUnitTests struct {
	*unit_test.UnitTestSharedFixture
	*unit_test.UnitTestMockFixture
}

func TestUpdateProductUnit(t *testing.T) {
	suite.Run(t, &updateProductUnitTests{UnitTestSharedFixture: unit_test.NewUnitTestSharedFixture(t)})
}

func (c *updateProductUnitTests) SetupTest() {
	// create new mocks or clear mocks before executing
	c.UnitTestMockFixture = unit_test.NewUnitTestMockFixture(c.T())
}

func (c *updateProductUnitTests) Test_New_Update_Product_Should_Return_No_Error_For_Valid_Input() {
	id := uuid.NewV4()
	name := gofakeit.Name()
	description := gofakeit.EmojiDescription()
	price := gofakeit.Price(150, 6000)

	updateProduct, err := NewUpdateProduct(id, name, description, price)

	c.Assert().NotNil(updateProduct)
	c.Assert().Equal(id, updateProduct.ProductID)
	c.Assert().Equal(name, updateProduct.Name)
	c.Assert().Equal(price, updateProduct.Price)

	c.Require().NoError(err)
}

func (c *updateProductUnitTests) Test_New_Update_Product_Should_Return_Error_For_Invalid_Price() {
	command, err := NewUpdateProduct(uuid.NewV4(), gofakeit.Name(), gofakeit.EmojiDescription(), 0)

	c.Require().Error(err)
	c.Assert().Nil(command)
}

func (c *updateProductUnitTests) Test_New_Update_Product_Should_Return_Error_For_Empty_Name() {
	command, err := NewUpdateProduct(uuid.NewV4(), "", gofakeit.EmojiDescription(), 120)

	c.Require().Error(err)
	c.Assert().Nil(command)
}

func (c *updateProductUnitTests) Test_New_Update_Product_Should_Return_Error_For_Empty_Description() {
	command, err := NewUpdateProduct(uuid.NewV4(), gofakeit.Name(), "", 120)

	c.Require().Error(err)
	c.Assert().Nil(command)
}