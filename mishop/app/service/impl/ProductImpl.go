package impl

type ProductImpl struct {
	Id   string
	Name string
}

func (product ProductImpl) Echo() string {
	return product.Id
}
