package infrastructure

import (
	"github.com/mhayk/GO-Expert-stress-test/internal/domain"
	"github.com/mhayk/GO-Expert-stress-test/internal/service"
)

type Container struct {
	loadTester      domain.LoadTester
	reportGenerator domain.ReportGenerator
	cli             CLI
}

func NewContainer() *Container {
	container := &Container{
		loadTester:      service.NewLoadTesterService(),
		reportGenerator: service.NewReportService(),
	}
	container.cli = NewCLI(container.loadTester, container.reportGenerator)
	return container
}

func (c *Container) GetLoadTester() domain.LoadTester {
	return c.loadTester
}

func (c *Container) GetReportGenerator() domain.ReportGenerator {
	return c.reportGenerator
}

func (c *Container) GetCLI() CLI {
	return c.cli
}
