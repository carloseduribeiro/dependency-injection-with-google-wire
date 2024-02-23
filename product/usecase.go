package product

type UseCase struct {
	repository RepositoryInterface
}

func NewProductUseCase(repository RepositoryInterface) *UseCase {
	return &UseCase{repository: repository}
}

func (u *UseCase) Execute(id int) (Product, error) {
	return u.repository.GetProductById(id)
}
