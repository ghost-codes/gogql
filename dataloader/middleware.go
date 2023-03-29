package dataloader

import (
	"context"
	"net/http"

	db "github.com/ghost-codes/gogql/db/sqlc"
)

func Middleware(store db.Store) (func (http.Handler) http.Handler){
    return func (next http.Handler) http.Handler{
        return http.HandlerFunc(func (w http.ResponseWriter,r *http.Request){
            ctx:=r.Context()
            loaders:=newLoaders(ctx,store);
            augmentedCtx := context.WithValue(ctx,key,loaders)

            r= r.WithContext(augmentedCtx)
            next.ServeHTTP(w,r)
        })
    }
}
