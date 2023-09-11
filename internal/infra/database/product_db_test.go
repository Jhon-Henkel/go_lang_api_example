package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/Jhon-Henkel/go_lang_api_example/tree/main/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10)
	assert.Nil(t, err)
	productDB := NewProduct(db)

	err = productDB.Create(product)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProductsWithPagination(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.Nil(t, err)
		db.Create(product)
		assert.Nil(t, err)
	}
	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.Equal(t, 10, len(products))
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.Nil(t, err)
	assert.Equal(t, 10, len(products))
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(products))
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestFindAllProductsWithoutPagination(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.Nil(t, err)
		db.Create(product)
		assert.Nil(t, err)
	}
	productDB := NewProduct(db)
	products, err := productDB.FindAll(0, 0, "asc")
	assert.Nil(t, err)
	assert.Equal(t, 23, len(products))
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 23", products[22].Name)
}

func TestFindProductById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10)
	assert.Nil(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	productFound, err := productDB.FindById(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestUpdataProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10)
	assert.Nil(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	product.Name = "Product 2"
	product.Price = 20
	err = productDB.Update(product)
	assert.Nil(t, err)
	productFound, err := productDB.FindById(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10)
	assert.Nil(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	err = productDB.Delete(product.ID.String())
	assert.Nil(t, err)
	_, err = productDB.FindById(product.ID.String())
	assert.Error(t, err)
}
