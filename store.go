package raftd

import (
	"path"
	"errors"
	"encoding/json"
)
type Store sturct {
	Nodes map[string]string `json:"nodes"`
}

func createStore()*Store{
	s:=new(Store)
	s.Nodes=make(map[string]string)
	return s
}

func (s *Store) Set(key string, value []byte) (string bool){

	key =path.Clean(key)

	oldValue,ok :=s.Nodes[key]

	if ok {
		s.Nodes[key] = value
		return oldValue,true
	}else{
		s.Nodes[key] =value
		return "",false
    }

}

func (s *Store) Get(key string) (string, error) {
	key = path.Clean(key)

	value, ok := s.Nodes[key]

	if ok {
		return value, nil
	} else {
		return "", errors.New("Key does not exist")
	}
}

func (s *Store) Delete(key string) (string, error) {
	key = path.Clean(key)

	oldValue, ok := s.Nodes[key]

	if ok {
		delete(s.Nodes, key)
		return oldValue, nil
	} else {
		return "", errors.New("Key does not exist")
	}
}

func (s *Store) Save() ([]byte, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (s *Store) Recovery(state []byte) error {
	err := json.Unmarshal(state, s)
	return err
}