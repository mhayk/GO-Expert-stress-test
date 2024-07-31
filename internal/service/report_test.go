package service

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/mhayk/GO-Expert-stress-test/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type reportServiceTestSuite struct {
	suite.Suite
	rs *reportService
}

func (suite *reportServiceTestSuite) SetupTest() {
	suite.rs = &reportService{}
}

func (suite *reportServiceTestSuite) TestGenerateReport() {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	testResult := domain.TestResult{
		Duration:      10 * time.Second,
		TotalRequests: 100,
		StatusCounts:  map[int]int{200: 90, 404: 5, 500: 5},
	}

	expectedOutput := fmt.Sprintf(" ___  ____  ____  ____  ___  ___    ____  ____  ___  ____\n")
	expectedOutput += fmt.Sprintf("/ __)(_  _)(  _ \\( ___)/ __)/ __)  (_  _)( ___)/ __)(_  _)\n")
	expectedOutput += fmt.Sprintf("\\__ \\  )(   )   / )__) \\__ \\\\__ \\    )(   )__) \\__ \\  )(\n")
	expectedOutput += fmt.Sprintf("(___/ (__) (_)\\_)(____)(___/(___/   (__) (____)(___/ (__)\n")
	expectedOutput += fmt.Sprintf("------------------------------------------------------------\n")
	expectedOutput += fmt.Sprintf("Relatório de teste de stress\n")
	expectedOutput += fmt.Sprintf("Tempo total gasto: %v\n", testResult.Duration)
	expectedOutput += fmt.Sprintf("Total de requests: %d\n", testResult.TotalRequests)
	expectedOutput += fmt.Sprintf("Requests com status 200: %d\n", testResult.StatusCounts[200])
	expectedOutput += fmt.Sprintf("Status 404: %d\n", testResult.StatusCounts[404])
	expectedOutput += fmt.Sprintf("Status 500: %d\n", testResult.StatusCounts[500])
	expectedOutput += fmt.Sprintf("------------------------------------------------------------\n")

	suite.rs.GenerateReport(testResult)

	w.Close()
	out := captureOutput(r)
	os.Stdout = old

	assert.Equal(suite.T(), expectedOutput, out)
}

func captureOutput(r *os.File) string {
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestReportService(t *testing.T) {
	suite.Run(t, new(reportServiceTestSuite))
}
