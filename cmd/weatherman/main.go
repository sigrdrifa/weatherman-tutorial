package main

import (
	"fmt"
	"os"

	"github.com/sigrdrifa/weatherman/internal/httpclient"
	"github.com/sigrdrifa/weatherman/internal/logger"
	"github.com/sigrdrifa/weatherman/pkg/client"
)

func main() {
  logger := logger.NewLogger(true)
  logger.Debug("Starting weatherman..")

  httpClient := httpclient.NewHttpClient()

  cl, err := client.NewClient(logger, httpClient)
  if err != nil {
    logger.Error(err)
    os.Exit(1)
  }

  res, err := cl.GetWeather(5.999802, 56.308920)
  if err != nil {
    logger.Debug("Failed to GetWeather()")
    logger.Error(err)
    os.Exit(1)
  }
  logger.Debug("Got response from the Oracle!")
  fmt.Println(res)
}
