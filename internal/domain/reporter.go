package domain

//go:generate mockery --name ReportGenerator --outpkg mock --output mock --filename report_generator.go --with-expecter=true

type ReportGenerator interface {
	GenerateReport(result TestResult)
}
