package metrics

func NewMockStore() Store {
	return &mockStore{}
}

type mockStore struct{}

func (*mockStore) Get(family, name string) float64 {
	return 0
}
func (*mockStore) Watch(family, name string) error {
	return nil
}
func (*mockStore) Display(family, name string) string {
	return "0"
}
func (*mockStore) Prepare() error {
	return nil
}
func (*mockStore) Collect(pid int) error {
	return nil
}
func (*mockStore) AddSource(name string, config map[string]string) (Source, error) {
	return nil, nil
}
func (*mockStore) Families() []string {
	return []string{"cpu"}
}
func (*mockStore) MetricNames(family string) []string {
	return []string{"user"}
}

func (*mockStore) Each(func(family, name string, metric Metric)) {
}
func (*mockStore) Save(family, name string, value float64) {
}
func (*mockStore) DeclareCounter(family, name string, xform TransformFunc, display DisplayFunc) {
}
func (*mockStore) DeclareGauge(family, name string, display DisplayFunc) {
}
func (*mockStore) Metric(family, name string) Metric {
	return &gauge{}
}
