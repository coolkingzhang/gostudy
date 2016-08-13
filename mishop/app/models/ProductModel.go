package models

type ProductModel interface {
	GetProduct() string
}
type ProductModelImpl struct {
}

func (productModelImpl ProductModelImpl) GetProduct() string {
	return "from model product"
}
