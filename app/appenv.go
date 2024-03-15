package main

type AppEnv string

const (
	Production  AppEnv = "production"
	Test        AppEnv = "test"
	Development AppEnv = "development"
)

func (e AppEnv) IsProduction() bool {
	return e == Production
}

func (e AppEnv) IsTest() bool {
	return e == Test
}

func (e AppEnv) IsDevelopment() bool {
	return e == Development
}
