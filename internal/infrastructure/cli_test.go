package infrastructure

import (
	"flag"
	"os"
	"testing"

	"github.com/mhayk/GO-Expert-stress-test/internal/domain"
	mock_domain "github.com/mhayk/GO-Expert-stress-test/internal/domain/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CLITestSuite struct {
	suite.Suite
	loadTesterMock      *mock_domain.LoadTester
	reportGeneratorMock *mock_domain.ReportGenerator
	cli                 *cli
	exitCode            int
	usageCalled         bool
}

func (suite *CLITestSuite) SetupTest() {
	suite.loadTesterMock = mock_domain.NewLoadTester(suite.T())
	suite.reportGeneratorMock = mock_domain.NewReportGenerator(suite.T())

	suite.cli = &cli{
		loadTester:      suite.loadTesterMock,
		reportGenerator: suite.reportGeneratorMock,
	}

	suite.exitCode = 0
	suite.usageCalled = false
}

func (suite *CLITestSuite) TestCLI_Execute() {
	testConfig := domain.TestConfig{
		URL:           "http://test.com",
		TotalRequests: 100,
		Concurrency:   10,
	}
	testResult := domain.TestResult{
		Duration:      0,
		TotalRequests: 100,
		StatusCounts:  map[int]int{200: 100},
	}

	suite.loadTesterMock.On("RunLoadTest", testConfig).Return(testResult)
	suite.reportGeneratorMock.On("GenerateReport", testResult).Return()

	suite.T().Run("Valid Arguments", func(t *testing.T) {
		os.Args = []string{"cmd", "--url=http://test.com", "--requests=100", "--concurrency=10"}
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError) // Reset flags

		suite.cli.Execute()

		suite.loadTesterMock.AssertExpectations(t)
		suite.reportGeneratorMock.AssertExpectations(t)
		assert.Equal(t, 0, suite.exitCode)
		assert.False(t, suite.usageCalled)
	})

}

func TestCLITestSuite(t *testing.T) {
	suite.Run(t, new(CLITestSuite))
}
