package main

import (
	"github.com/YaminThiriAung21/UniversityGolang/Database"
	"github.com/YaminThiriAung21/UniversityGolang/Route"
)

func main() {
	Database.Connectdb()
	Route.LoadRoute()
}
