package option

import (
	kconf "github.com/go-kratos/kratos/v2/config"
	"github.com/jacexh/gopkg/config"
)

type (
	fakeSource struct {
		s kconf.Source
	}

	fakeWatcher struct {
		w kconf.Watcher
	}
)

func (f fakeSource) Load() ([]*config.KeyValue, error) {
	kvs, err := f.s.Load()
	if err != nil {
		return nil, err
	}
	var ret = make([]*config.KeyValue, len(kvs))
	for i, kv := range kvs {
		ret[i] = &config.KeyValue{
			Key:    kv.Key,
			Value:  kv.Value,
			Format: kv.Format,
		}
	}
	return ret, nil
}

func (f fakeSource) Watch() (config.Watcher, error) {
	w, err := f.s.Watch()
	if err != nil {
		return nil, err
	}
	return fakeWatcher{w: w}, err
}

func (fw fakeWatcher) Next() ([]*config.KeyValue, error) {
	kvs, err := fw.w.Next()
	if err != nil {
		return nil, err
	}
	var ret = make([]*config.KeyValue, len(kvs))
	for i, kv := range kvs {
		ret[i] = &config.KeyValue{
			Key:    kv.Key,
			Value:  kv.Value,
			Format: kv.Format,
		}
	}
	return ret, nil
}

func (fw fakeWatcher) Stop() error {
	return fw.w.Stop()
}

func convertSource(source kconf.Source) config.Source {
	return fakeSource{
		s: source,
	}
}
