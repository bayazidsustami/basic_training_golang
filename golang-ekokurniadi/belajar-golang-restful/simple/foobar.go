package simple

type FooBarService struct {
	*FooService
	*BarService
}

func NewFooBarService(fooService *FooService, barservice *BarService) *FooBarService {
	return &FooBarService{
		FooService: fooService,
		BarService: barservice,
	}
}
