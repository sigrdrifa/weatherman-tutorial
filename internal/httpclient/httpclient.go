package httpclient

import "time"

type HttpClient struct {
}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

func (h *HttpClient) GetWeatherData(longitude float64, latitude float64) (string, error) {
	time.Sleep(time.Millisecond * 25)

	if longitude > -20 && longitude < 50 && latitude > 50 {
		// this looks like northern europe..
		return `Rain and wind are in your immediate future, my friend. Today it will likely rain 
    and tomorrow it will probably rain too. Also, have you heard of this thing called Spring? 
    Not for you! Winter is coming, eternal winter. Strength, friend.`, nil
	}
	return `Even if the sun isn't beaming right now, there is a good chance it will soon! There 
  is hope for you and your loved ones, you may even be able to go to the beach at some point 
  and sip cocktails while listening to the gentle sounds of the ocean waves washing ashore. 
  The world is a beautiful place...`, nil
}
