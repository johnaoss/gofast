package main

import (
	"log"
	"sort"

	"github.com/ddo/go-fast"
)

var (
	client = initClient()
)

type Client struct {
	f     *fast.Fast
	urls  []string
	ready bool
}

// Result is the data obtained from a speedtest.
type Result struct {
	Data    []float64
	Average float64
	Min     float64
	Max     float64
}

func initClient() *Client {
	f := fast.New()

	var failed bool
	if err := f.Init(); err != nil {
		log.Print("Failed to init fast", err.Error())
		failed = true
	}

	urls, err := f.GetUrls()
	if err != nil {
		log.Print("Failed to get urls:", err)
		failed = true
	}

	return &Client{
		urls:  urls,
		f:     f,
		ready: !failed,
	}
}

// Measure measures the network speed from the computer to `fast.com`'s servers.
func (c *Client) Measure() Result {
	if c == nil || !c.ready {
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
		Data:    results,
		Average: average,
		Min:     results[0],
		Max:     results[len(results)-1],
	}
}

// Ready returns whether or not the client is actually ready to be used.
// This may return false if there are network issues, to which it will attempt
// to reconnect.
func (c *Client) Ready() bool {
	if c.ready {
		return true
	}

	// Otherwise, we need to reinstantiate the client.
	if err := c.f.Init(); err != nil {
		log.Printf("Failed to reinitialize the client: %v", err)
		return false
	}

	c.ready = true
	return c.ready

}
