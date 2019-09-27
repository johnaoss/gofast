package main

import (
	"sort"
	"log"

	"github.com/ddo/go-fast"
)

var (
	client = initClient()
)

type Client struct {
	f *fast.Fast
	urls []string
}

// Result is the data obtained from a speedtest.
type Result struct {
	Data []float64
	Average float64
	Min float64
	Max float64

}

func initClient() *Client {

	f := fast.New()

	var err error
	if err := f.Init(); err != nil {
		log.Print("Failed to init fast", err.Error())
		return nil
	}

	urls, err := f.GetUrls()
	if err != nil {
		log.Print("Failed to get urls:", err)
		return nil
	}

	return &Client{
		urls: urls,
		f: f,
	}
}

func (c *Client) Measure() Result {
	if c == nil {
		return Result{}
	}

	kbpsChan := make(chan float64)
	// Read from the channel (this is so bad)
	results := make([]float64, 0, 24)
	go func() {
		for Kbps := range kbpsChan {
			results = append(results, Kbps)
		}
	}()
	
	if err := c.f.Measure(c.urls, kbpsChan); err != nil {
		log.Print("Failed to measure speed:", err)
	}

	var average float64
	for _, res := range results {
		average += res / float64(len(results))
	}

	sort.Float64s(results)

	return Result{
		Data: results,
		Average: average,
		Min: results[0],
		Max: results[len(results)-1],
	}
}