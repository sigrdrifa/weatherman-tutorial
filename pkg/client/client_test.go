package client_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	mocks "github.com/sigrdrifa/weatherman/mocks/pkg/client"
	"github.com/sigrdrifa/weatherman/pkg/client"
	"github.com/stretchr/testify/mock"
)

var testT *testing.T

func TestClient(t *testing.T) {
	RegisterFailHandler(Fail)
	testT = t
	RunSpecs(t, "Client Suite")
}

var _ = Describe("Client Suite", func() {

	var mockHttpClient *mocks.HttpClient
	var mockLogger *mocks.Logger

	BeforeEach(func() {
		mockLogger = mocks.NewLogger(testT)
		mockHttpClient = mocks.NewHttpClient(testT)
	})

	AfterEach(func() {
		mockLogger = nil
		mockHttpClient = nil
	})

	var _ = Describe("Getting Weather", func() {

		When("Passing it Northern Europe coordinates", func() {

			It("Should tell us that the weather is bad", func() {
				long := 15.67621
				lat := 48.15215

				mockLogger.On("Debug", mock.Anything)

				mockHttpClient.On("GetWeatherData", mock.Anything, mock.Anything).Return(
					`Rain and wind are in your immedeate future, my friend. Today it will likely rain 
    and tomorrow it will probably rain too. Also, have you heard of this thing called Spring? 
    Not for you! Winter is coming, eternal winter. Strength, friend.`, nil)

				cl, err := client.NewClient(mockLogger, mockHttpClient)
				Expect(err).NotTo(HaveOccurred())

				res, err := cl.GetWeather(long, lat)
				Expect(err).NotTo(HaveOccurred())
				Expect(res).To(Equal(
					`Rain and wind are in your immedeate future, my friend. Today it will likely rain 
    and tomorrow it will probably rain too. Also, have you heard of this thing called Spring? 
    Not for you! Winter is coming, eternal winter. Strength, friend.`))
			})
		})
	})
})
