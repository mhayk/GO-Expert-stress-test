// Code generated by mockery v2.43.2. DO NOT EDIT.

package mock

import (
	domain "github.com/mhayk/GO-Expert-stress-test/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// ReportGenerator is an autogenerated mock type for the ReportGenerator type
type ReportGenerator struct {
	mock.Mock
}

type ReportGenerator_Expecter struct {
	mock *mock.Mock
}

func (_m *ReportGenerator) EXPECT() *ReportGenerator_Expecter {
	return &ReportGenerator_Expecter{mock: &_m.Mock}
}

// GenerateReport provides a mock function with given fields: result
func (_m *ReportGenerator) GenerateReport(result domain.TestResult) {
	_m.Called(result)
}

// ReportGenerator_GenerateReport_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateReport'
type ReportGenerator_GenerateReport_Call struct {
	*mock.Call
}

// GenerateReport is a helper method to define mock.On call
//   - result domain.TestResult
func (_e *ReportGenerator_Expecter) GenerateReport(result interface{}) *ReportGenerator_GenerateReport_Call {
	return &ReportGenerator_GenerateReport_Call{Call: _e.mock.On("GenerateReport", result)}
}

func (_c *ReportGenerator_GenerateReport_Call) Run(run func(result domain.TestResult)) *ReportGenerator_GenerateReport_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.TestResult))
	})
	return _c
}

func (_c *ReportGenerator_GenerateReport_Call) Return() *ReportGenerator_GenerateReport_Call {
	_c.Call.Return()
	return _c
}

func (_c *ReportGenerator_GenerateReport_Call) RunAndReturn(run func(domain.TestResult)) *ReportGenerator_GenerateReport_Call {
	_c.Call.Return(run)
	return _c
}

// NewReportGenerator creates a new instance of ReportGenerator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReportGenerator(t interface {
	mock.TestingT
	Cleanup(func())
}) *ReportGenerator {
	mock := &ReportGenerator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
