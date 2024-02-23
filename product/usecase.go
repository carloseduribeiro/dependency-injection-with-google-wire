package product

type UseCase struct {
	repository *Repository
}

func NewProductUseCase(repository *Repository) *UseCase {
	return &UseCase{repository: repository}
}

func (u *UseCase) Execute(id int) (Product, error) {
	return u.repository.GetProductById(id)
}
