package models

import "errors"

type Tournament struct {
  Id int64
}

func (h Tournament) GetById(id int64) (*Tournament, error) {
  return nil, errors.New("Blah")
}

