package main

import (
	"github.com/Dendrafrz/test_gin/mappings"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mappings.CreateUrlMappings()
	//listen server on 0.0.0.0:8090
	mappings.Router.Run(":8080")
}
