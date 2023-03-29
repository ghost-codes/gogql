package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ghost-codes/gogql/dataloader"
	db "github.com/ghost-codes/gogql/db/sqlc"
	"github.com/ghost-codes/gogql/graph"
	"github.com/ghost-codes/gogql/util"
	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}


    config,err:= util.LoadConfig(".");

    if err!=nil{
        log.Fatal("error occured while loading config vars",err);
    }

    conn,err:= sql.Open("postgres",config.DBSource)
    if err!=nil{
        log.Fatal("connection could not be established with db:",err)
    }

    store:=db.NewStore(conn);
    srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
        Store: *store,
        Config: config,
        DataLoaders: dataloader.NewRetriever(),
    }}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
    
    dlMiddleware:= dataloader.Middleware(*store);
    http.Handle("/graphql", dlMiddleware(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
