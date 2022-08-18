package domains

type Database interface {
	CompanyRepoUseCase
	UserRepoUseCase
}
