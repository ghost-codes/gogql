package dataloader

//go:generate go run github.com/vektah/dataloaden UserLoader int64 "*github.com/ghost-codes/gogql/db/sqlc.User"

import (
	"context"
	"time"

	db "github.com/ghost-codes/gogql/db/sqlc"
)


type contextKey string

const key= contextKey("dataloader")

//Loaders holds reference to the individual dataloaders
type Loaders struct{
    //Individual loaders will be initialized
    UserByVideoID *UserLoader
}

func newLoaders(ctx context.Context,store db.Store)*Loaders{
    return &Loaders{
        // inidividual loaders will be initialized here
        UserByVideoID: newUserByVideoID(ctx,store),
    }
}

func newUserByVideoID(ctx context.Context, store db.Store) *UserLoader{
    return NewUserLoader(UserLoaderConfig{
        MaxBatch: 100,
        Wait: 5*time.Millisecond,
        Fetch: func(keys []int64) ([]*db.User, []error) {
            // db query
            res,err := store.ListOfUsers(ctx,keys)
            if err!= nil{
                return nil, []error{err}
            }
            //map
            groupByVideoID :=make(map[int64]*db.User,len(keys));
            for _,r:= range res{
                groupByVideoID[int64(r.VideoID)]=&db.User{
                    ID: r.ID,
                    Name: r.Name,
                }
        }

        //order
        result :=make([]*db.User,len(keys))

        for i,re:=range keys{
            result[i]=groupByVideoID[re]
        }

        return result,nil
    },
})
}


// Retriever retrieves dataloaders from the request context.
type Retriever interface{
    Retrieve(context.Context) *Loaders
}

type retriever struct{
    key contextKey 
}

func (r *retriever) Retrieve(ctx context.Context) *Loaders{
    return ctx.Value(r.key).(*Loaders)
}


func NewRetriever() Retriever{
    return &retriever{key:key}
}
