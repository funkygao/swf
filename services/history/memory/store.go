package memory

import (
	"github.com/funkygao/swf/services/history"
	"github.com/hashicorp/go-memdb"
)

type History struct {
	schema *memdb.DBSchema
	db     *memdb.MemDB
}

func New() history.Service {
	return &History{
		schema: &memdb.DBSchema{
			Tables: map[string]*memdb.TableSchema{
				"history": &memdb.TableSchema{
					Name: "history",
					Indexes: map[string]*memdb.IndexSchema{
						"id": &memdb.IndexSchema{
							Name:   "id",
							Unique: true,
							Indexer: &memdb.CompoundIndex{
								AllowMissing: false,
								Indexes: []memdb.Indexer{
									&memdb.StringFieldIndex{
										Lowercase: false,
										Field:     "Name",
									},
									&memdb.StringFieldIndex{
										Lowercase: false,
										Field:     "Version",
									},
								},
							},
						},
					},
				},
			},
		},
	}

}

func (this *History) Name() string {
	return "history"
}

func (this *History) Start() (err error) {
	if err = this.schema.Validate(); err != nil {
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
