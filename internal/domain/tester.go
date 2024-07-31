package domain

//go:generate mockery --name LoadTester --outpkg mock --output mock --filename load_tester.go --with-expecter=true

type LoadTester interface {
	RunLoadTest(config TestConfig) TestResult
}
