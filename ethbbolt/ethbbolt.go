package ethbbolt

import (
  "fmt"
  "path/filepath"

  "github.com/etcd-io/bbolt"
)

type EDB struct {
  d DB
}

func (edb *EDB) Close() {
  edb.d.Close()
}
