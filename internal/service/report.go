package service

import (
	"fmt"

	"github.com/mhayk/GO-Expert-stress-test/internal/domain"
)

type reportService struct{}

func NewReportService() domain.ReportGenerator {
	return &reportService{}
}

func (s *reportService) GenerateReport(result domain.TestResult) {

	fmt.Println(" ___  ____  ____  ____  ___  ___    ____  ____  ___  ____")
	fmt.Println("/ __)(_  _)(  _ \\( ___)/ __)/ __)  (_  _)( ___)/ __)(_  _)")
	fmt.Println("\\__ \\  )(   )   / )__) \\__ \\\\__ \\    )(   )__) \\__ \\  )(")
	fmt.Println("(___/ (__) (_)\\_)(____)(___/(___/   (__) (____)(___/ (__)")
	fmt.Println("------------------------------------------------------------")
	fmt.Println("Relatório de teste de stress")
	fmt.Printf("Tempo total gasto: %v\n", result.Duration)
	fmt.Printf("Total de requests: %d\n", result.TotalRequests)
	fmt.Printf("Requests com status 200: %d\n", result.StatusCounts[200])

	for status, count := range result.StatusCounts {
		if status != 200 {
			fmt.Printf("Status %d: %d\n", status, count)
		}
	}
	fmt.Println("------------------------------------------------------------")
}
