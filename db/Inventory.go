package db

type Inventory struct {
	ID        int    `db:"id" json:"id"`
	Name      string `db:"name" json:"name" binding:"required"`
	ProjectID int    `db:"project_id" json:"project_id"`
	Inventory string `db:"inventory" json:"inventory"`

	// accesses dynamic inventory
	KeyID *int      `db:"key_id" json:"key_id"`
	Key   AccessKey `db:"-" json:"-"`
	// accesses hosts in inventory
	SshKeyID *int      `db:"ssh_key_id" json:"ssh_key_id"`
	SshKey   AccessKey `db:"-" json:"-"`

	// static/aws/do/gcloud
	Type string `db:"type" json:"type"`

	Removed bool `db:"removed" json:"removed"`
}
