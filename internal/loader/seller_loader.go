package loader

type SellerLoader interface {
	Load() (v map[int]models.Vehicle, err error)
}
