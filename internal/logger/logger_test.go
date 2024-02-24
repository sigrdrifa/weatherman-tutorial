package logger_test

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sigrdrifa/weatherman/internal/logger"
	"testing"
)

func TestLogger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Logger Suite")
}

var _ = Describe("Logger Suite", func() {

	var loggerInstance *logger.Logger

	BeforeEach(func() {
		loggerInstance = logger.NewLogger(true)
	})

	var _ = Describe("Logging", func() {

		It("Should store DEBUG log messages", func() {
			loggerInstance.Debug("This is a test")
			loggerInstance.Debug("This is another test")

			Expect(len(loggerInstance.Flush())).To(Equal(2))
		})

		It("Should store ERROR messages", func() {
			loggerInstance.Debug("This is a test")
			loggerInstance.Error(errors.New("This is a test error"))

			msgs := loggerInstance.Flush()

			Expect(len(msgs)).To(Equal(2))
			Expect(msgs[1]).To(Equal("This is a test error"))
		})
	})

})
