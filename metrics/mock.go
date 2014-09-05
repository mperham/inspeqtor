package metrics

func NewMockStore() Store {
	return &mockStore{}
}

type mockStore struct{}

func (*mockStore) Get(family string, name string) int64 {
	return 0
}
func (*mockStore) Display(family string, name string) string {
	return "0"
}
func (*mockStore) PrepareRule(family string, name string, threshold int64) (int64, error) {
	return 0, nil
}
func (*mockStore) Collect(pid int) error {
	return nil
}

func (*mockStore) Families() []string {
	return []string{"cpu"}
}
func (*mockStore) Metrics(family string) []string {
	return []string{"user"}
}

func (*mockStore) Save(family, name string, value int64) {
}
func (*mockStore) DeclareCounter(family, name string, xform TransformFunc, display DisplayFunc) {
}
func (*mockStore) DeclareGauge(family, name string, prep PrepareFunc, display DisplayFunc) {
}
