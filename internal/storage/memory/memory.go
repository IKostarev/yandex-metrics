package memory

type MemoryStorage struct {
	gauge   map[string]float64
	counter map[string]int64
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		gauge:   make(map[string]float64),
		counter: make(map[string]int64),
	}
}

func (m *MemoryStorage) AddGauge(name string, value float64) {
	m.gauge[name] = value
}

func (m *MemoryStorage) AddCounter(name string, value int64) {
	if v, ok := m.counter[name]; ok {
		m.counter[name] = v + value
	} else {
		m.counter[name] = value
	}
}
