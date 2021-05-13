package datadog

import (
	"context"
	"fmt"
	"os"
	"time"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/datadog", new(Datadog))
}

// Datadog is the k6 datadog extension.
type Datadog struct {
}

// MetricQueryResponse is used to return the response to the user.
type MetricQueryResponse struct {
	Expression string
	Unit       string
	Points     []float64
}

// Client is the Datadog client wrapper.
type Client struct {
	ctx    *context.Context
	client *datadog.APIClient
}

// NewClient creates a new Datadog client with the provided options.
func (r *Datadog) NewClient(apiKey string, appKey string) *Client {
	ctx := context.WithValue(
		context.Background(),
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"apiKeyAuth": {
				Key: apiKey,
			},
			"appKeyAuth": {
				Key: appKey,
			},
		},
	)
	configuration := datadog.NewConfiguration()
	return &Client{ctx: &ctx, client: datadog.NewAPIClient(configuration)}
}

// MetricQuery queries Datadog metrics.
func (c *Client) MetricQuery(query string, since int64) []*MetricQueryResponse {
	now := time.Now().Unix()

	if since == 0 {
		since = 3600
	}

	resp, r, err := c.client.MetricsApi.QueryMetrics(*c.ctx).From(now - since).To(now).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.QueryMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	var series []*MetricQueryResponse
	if resp.GetStatus() == "ok" {
		for _, s := range resp.GetSeries() {
			var flatPoints []float64
			for _, point := range s.GetPointlist() {
				flatPoints = append(flatPoints, point[1])
			}
			unitName := s.GetUnit()[0].GetShortName()
			if len(unitName) == 0 {
				unitName = "unknown"
			}
			series = append(series, &MetricQueryResponse{Expression: s.GetExpression(), Unit: unitName, Points: flatPoints})
		}
	}
	return series
}

// Avg calculates the avg over a set of points.
func (m *MetricQueryResponse) Avg() float64 {
	var total float64 = 0
	for _, value := range m.Points {
		total += value
	}
	return total / float64(len(m.Points))
}

// Max retrieves the max from a set of points.
func (m *MetricQueryResponse) Max() float64 {
	max := m.Points[0]
	for _, v := range m.Points {
		if v > max {
			max = v
		}
	}
	return max
}

// Min retrieves the min from a set of points.
func (m *MetricQueryResponse) Min() float64 {
	min := m.Points[0]
	for _, v := range m.Points {
		if v < min {
			min = v
		}
	}
	return min
}
