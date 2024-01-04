//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializeService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository,
		NewSimpleService,
	)
	return nil, nil
}

func InitializeDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMongoDB,
		NewDatabasePostgreSQL,
		NewDatabaseRepository,
	)
	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializeFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializeSayHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil
}

var fooBarSet = wire.NewSet(
	NewFoo,
	NewBar,
)

func IntitializeFooBar() *FooBar {
	wire.Build(
		fooBarSet,
		wire.Struct(new(FooBar), "Foo", "Bar"),
	)
	return nil
}

var fooBarValueSet = wire.NewSet(
	wire.Value(&Foo{}),
	wire.Value(&Bar{}),
)

func IntitializeFooBarUsingvalue() *FooBar {
	wire.Build(
		fooBarSet,
		wire.Struct(new(FooBar), "*"),
	)
	return nil
}
