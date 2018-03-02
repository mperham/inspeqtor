package expose

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"text/template"

	"github.com/mperham/inspeqtor"
	"github.com/mperham/inspeqtor/util"
)

//go:generate go-bindata -pkg expose -o static.go memory/... static/...

func init() {
	inspeqtor.BuildExpose = exposeBuilder
}

var (
	exposedServices = map[string]*memstatsService{}
	// Such a gross hack, unclear why ServeMux doesn't gracefully
	// handle multiple registrations.
	memstatsRegistered = false
)

func exposeBuilder(global *inspeqtor.ConfigFile, chk inspeqtor.Checkable, elements []string, options map[string]string) error {
	for _, name := range elements {
		if name != "memstats" {
			return fmt.Errorf("Unknown expose option: %s", name)
		}
		src, err := chk.Metrics().AddSource(name, options)
		if err != nil {
			return err
		}

		exposedServices[chk.Name()] = &memstatsService{
			service: chk,
			source:  src.(*memstatsSource),
		}

		if !memstatsRegistered {
			http.Handle("/", http.FileServer(
				&AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "static"}))
			// renders the index page
			http.HandleFunc("/memory/", memoryRenderer)
			// renders real-time memstats metrics directly from service
			http.HandleFunc("/memstats.json", memstatsHandler)
			// renders hourly data from the metrics store
			http.HandleFunc("/metrics.json", metricsHandler)
			memstatsRegistered = true
		}
	}
	return nil
}

type memstatsService struct {
	service inspeqtor.Checkable
	source  *memstatsSource
}

func Bootstrap(ins *inspeqtor.Inspeqtor) error {
	return nil
}

var (
	idxTmpl *template.Template
)

type MemoryPage struct {
	AvailableServices []string
	Version           string
	Selected          string
}

func memoryRenderer(w http.ResponseWriter, r *http.Request) {
	def := ""
	svcs := []string{}
	for name := range exposedServices {
		if def == "" {
			def = name
		}
		svcs = append(svcs, name)
	}

	values := r.URL.Query()
	name := values.Get("service")
	if name == "" {
		name = def
	}

	if _, ok := exposedServices[name]; !ok {
		http.Error(w, fmt.Sprintf("No such service: %s", name), 404)
		return
	}
	sort.Strings(svcs)

	//var idxTmpl *template.Template
	if idxTmpl == nil {
		//asset, err := ioutil.ReadFile("expose/memory/index.html.tmpl")
		asset, err := Asset("memory/index.html.tmpl")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading template: %s", err.Error()), 503)
			return
		}

		tmp, err := template.New("index.html").Parse(string(asset))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing template: %s", err.Error()), 503)
			return
		}
		idxTmpl = tmp
	}

	ctxt := &MemoryPage{svcs, inspeqtor.VERSION, name}
	err := idxTmpl.Execute(w, ctxt)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error rendering template: %s", err.Error()), 500)
		return
	}
}

func memstatsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		name := r.URL.Query().Get("service")
		if name == "" {
			http.Error(w, "Missing input", 401)
			return
		}
		svc := exposedServices[name]
		if svc == nil {
			http.Error(w, fmt.Sprintf("%s Not Found", name), 404)
			return
		}
		i, err := easyMemstatsClient(svc.source.Location(), func(resp *http.Response) (int64, error) {
			w.Header().Set("Content-Type", "application/json")
			return io.Copy(w, resp.Body)
		})
		if i == 0 && err != nil {
			http.Error(w, err.Error(), 503)
			return
		}

	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func easyMemstatsClient(url string, msp func(*http.Response) (int64, error)) (int64, error) {
	util.DebugDebug("Fetching memstats from %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return msp(resp)
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		values := r.URL.Query()
		name := values.Get("service")
		if name == "" {
			http.Error(w, "Missing input", 401)
			return
		}
		svc := exposedServices[name]
		if svc == nil {
			http.Error(w, fmt.Sprintf("%s Not Found", name), 404)
			return
		}

		fam := values.Get("family")
		if fam == "" {
			http.Error(w, "Missing input", 401)
			return
		}

		names := svc.service.Metrics().MetricNames(fam)
		vals := map[string]map[string]float64{
			fam: map[string]float64{},
		}

		for _, met := range names {
			vals[fam][met] = svc.service.Metrics().Get(fam, met)
		}

		err := json.NewEncoder(w).Encode(vals)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error marshalling JSON: %s", err.Error()), 503)
		}
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}
