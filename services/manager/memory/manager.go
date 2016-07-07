package memory

import (
	"github.com/funkygao/swf/services/manager"
	"github.com/hashicorp/go-memdb"
)

type Manager struct {
	schema *memdb.DBSchema
	db     *memdb.MemDB
}

func New() manager.Service {
	return &Manager{
		schema: &memdb.DBSchema{
			Tables: map[string]*memdb.TableSchema{
				"ActivityType": &memdb.TableSchema{
					Name: "ActivityType",
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

				"WorkflowType": &memdb.TableSchema{
					Name: "WorkflowType",
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

func (this *Manager) Start() (err error) {
	if err = this.schema.Validate(); err != nil {
		return
	}

	this.db, err = memdb.NewMemDB(this.schema)
	if err != nil {
		return
	}

	return
}

func (this *Manager) Stop() {}
