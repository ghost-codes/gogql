package graph

import (
	"github.com/ghost-codes/gogql/dataloader"
	db "github.com/ghost-codes/gogql/db/sqlc"
	"github.com/ghost-codes/gogql/util"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
    Config util.Config
    Store db.Store
    DataLoaders dataloader.Retriever
}
