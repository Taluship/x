package cloud

import (
	"os"

	"github.com/zcharym/pocketbase-client"
)

// BaseClient
// Currently the base client is single node oriented, but it should be extended to support multi-node
type BaseClient struct {
	Pocketbase *pocketbase.Client
}

func NewClient() *BaseClient {
	c := &BaseClient{}
	c.Pocketbase = pocketbase.NewClient(
		os.Getenv("POCKETBASE_URL"),
		pocketbase.WithAdminEmailPassword(
			os.Getenv("POCKETBASE_EMAIL"),
			os.Getenv("POCKETBASE_PWD"),
		),
	)

	return c
}
