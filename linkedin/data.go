package linkedin

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"

	"github.com/spudtrooper/goutil/io"
)

type data struct {
	dir string
}

type state struct {
	Pending map[string]bool `json:"urls"`
	Trying  map[string]bool `json:"trying"`
	Done    map[string]bool `json:"done"`
	Failed  map[string]bool `json:"failed"`
}

func makeData(dataDir string) (*data, error) {
	dir, err := io.MkdirAll(dataDir)
	if err != nil {
		return nil, err
	}
	return &data{dir}, nil
}

func (d *data) Dir() string {
	return d.dir
}

func (d *data) CreateQueue(urls []string) error {
	s, err := d.readState()
	if err != nil {
		return err
	}
	log.Printf("urls: %+v", urls)
	for _, u := range urls {
		log.Println(u)
		if !s.Done[u] {
			s.Pending[u] = true
		}
	}
	if err := d.saveState(s); err != nil {
		return err
	}
	return nil
}

func (d *data) HasMore() (bool, error) {
	s, err := d.readState()
	if err != nil {
		return false, err
	}
	return len(s.Pending) > 0, nil
}

func (d *data) Next() (string, bool, error) {
	s, err := d.readState()
	if err != nil {
		return "", false, err
	}
	for u := range s.Pending {
		if !s.Done[u] {
			s.Pending[u] = false
			s.Trying[u] = true
			if err := d.saveState(s); err != nil {
				return "", false, err
			}
			return u, false, nil
		}
	}
	return "", true, nil
}

func (d *data) Try(url string) error {
	s, err := d.readState()
	if err != nil {
		return err
	}
	s.Pending[url] = false
	s.Trying[url] = true
	if err := d.saveState(s); err != nil {
		return err
	}
	return nil
}

func (d *data) saveState(s *state) error {
	b, err := json.Marshal(s)
	if err != nil {
		return err
	}
	f := path.Join(d.dir, "state.json")
	log.Printf("saving to %s", f)
	if ioutil.WriteFile(f, b, 0644); err != nil {
		return err
	}
	return nil
}

func (d *data) readState() (*state, error) {
	s, err := d.readStateInternal()
	if err != nil {
		return nil, err
	}
	if s.Done == nil {
		s.Done = map[string]bool{}
	}
	if s.Trying == nil {
		s.Trying = map[string]bool{}
	}
	if s.Pending == nil {
		s.Pending = map[string]bool{}
	}
	if s.Failed == nil {
		s.Failed = map[string]bool{}
	}
	return s, nil
}
func (d *data) readStateInternal() (*state, error) {
	f := path.Join(d.dir, "state.json")
	if !io.FileExists(f) {
		var s state
		b, err := json.Marshal(&s)
		if err != nil {
			return nil, err
		}
		log.Printf("starting with %s", f)
		if ioutil.WriteFile(f, b, 0644); err != nil {
			return nil, err
		}
		return &s, nil
	}
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	var s state
	if err := json.Unmarshal(b, &s); err != nil {
		return nil, err
	}
	return &s, nil
}
