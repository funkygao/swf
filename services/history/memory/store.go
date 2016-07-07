package memory

import (
	"github.com/funkygao/swf/services/historystore"
	"github.com/hashicorp/go-memdb"
)

type History struct {
	schema *memdb.DBSchema
	db     *memdb.MemDB
}

func New() historystore.Service {
	return &History{
		schema: &memdb.DBSchema{
			Tables: map[string]*memdb.TableSchema{
				"history": &memdb.TableSchema{
					Name:    "history",
					Indexes: map[string]*memdb.IndexSchema{},
				},
			},
		},
	}

}

func (this *History) Start() (err error) {
	if err == this.schema.Validate(); err != nil {
		return
	}

	this.db, err = memdb.NewMemDB(this.schema)
	if err != nil {
		return
	}

	return
}

func (this *History) Stop() {
}
