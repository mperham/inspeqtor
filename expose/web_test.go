package expose

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mperham/inspeqtor"
	"github.com/stretchr/testify/assert"
)

func TestExpose(t *testing.T) {
	// global state used below, can't run in parallel
	err := exposeBuilder(nil, nil, []string{"foo"}, map[string]string{})
	assert.NotNil(t, err)

	assert.False(t, memstatsRegistered)
	svc := inspeqtor.MockCheckable("something")
	assert.NotContains(t, svc.Metrics().Families(), "memstats")
	err = exposeBuilder(nil, svc, []string{"memstats"}, map[string]string{})
	assert.Contains(t, svc.Metrics().Families(), "memstats")
	assert.Equal(t, 5, len(svc.Metrics().MetricNames("memstats")))
	assert.True(t, memstatsRegistered)

	// reset global state
	memstatsRegistered = false
}

func TestMemoryRender(t *testing.T) {
	exposedServices = map[string]*memstatsService{}

	req, err := http.NewRequest("GET", "http://localhost:4677/memory/?service=nonesuch", nil)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	memoryRenderer(w, req)
	assert.Equal(t, 404, w.Code)

	src, err := buildMemstatsSource(map[string]string{})
	assert.Nil(t, err)
	exposedServices["inspeqtor"] = &memstatsService{
		inspeqtor.MockCheckable("inspeqtor"),
		src.(*memstatsSource),
	}

	req, err = http.NewRequest("GET", "http://localhost:4677/memory/", nil)
	assert.Nil(t, err)

	w = httptest.NewRecorder()
	memoryRenderer(w, req)
	assert.Equal(t, 200, w.Code)

	req, err = http.NewRequest("GET", "http://localhost:4677/memory/?service=inspeqtor", nil)
	assert.Nil(t, err)

	w = httptest.NewRecorder()
	memoryRenderer(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestMemstatsRender(t *testing.T) {
	exposedServices = map[string]*memstatsService{}
	req, _ := http.NewRequest("GET", "http://localhost:4677/memstats.json", nil)
	w := httptest.NewRecorder()
	memstatsHandler(w, req)
	assert.Equal(t, 401, w.Code)

	req, _ = http.NewRequest("GET", "http://localhost:4677/memstats.json?service=inspeqtor", nil)
	w = httptest.NewRecorder()
	memstatsHandler(w, req)
	assert.Equal(t, 404, w.Code)

	src, err := buildMemstatsSource(map[string]string{"port": "4677"})
	assert.Nil(t, err)
	exposedServices["inspeqtor"] = &memstatsService{
		inspeqtor.MockCheckable("inspeqtor"),
		src.(*memstatsSource),
	}

	req, _ = http.NewRequest("GET", "http://localhost:4677/memstats.json?service=inspeqtor", nil)
	w = httptest.NewRecorder()
	memstatsHandler(w, req)
	assert.Equal(t, 503, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "connection refused"), w.Body.String())
}

func TestMetricsRender(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost:4677/metrics.json?family=memstats", nil)
	w := httptest.NewRecorder()
	metricsHandler(w, req)
	assert.Equal(t, 401, w.Code)

	req, _ = http.NewRequest("GET", "http://localhost:4677/metrics.json?service=inspeqtor", nil)
	w = httptest.NewRecorder()
	metricsHandler(w, req)
	assert.Equal(t, 401, w.Code)

	req, _ = http.NewRequest("GET", "http://localhost:4677/metrics.json?service=nonesuch", nil)
	w = httptest.NewRecorder()
	metricsHandler(w, req)
	assert.Equal(t, 404, w.Code)

	req, _ = http.NewRequest("GET", "http://localhost:4677/metrics.json?service=inspeqtor&family=memstats", nil)
	w = httptest.NewRecorder()
	metricsHandler(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"memstats\":{}}\n", w.Body.String())
}
