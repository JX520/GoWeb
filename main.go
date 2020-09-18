package main

import (
	"Goweb/initRouter"
)

func main()  {
	router := initRouter.SetRouter()

	_ = router.Run()

}