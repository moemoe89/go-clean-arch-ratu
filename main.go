//
//  Simple REST API
//
//  Copyright © 2020. All rights reserved.
//

package main

import (
	conf "github.com/moemoe89/simple-go-clean-arch/config"
	_ "github.com/moemoe89/simple-go-clean-arch/docs"
	"github.com/moemoe89/simple-go-clean-arch/routers"

	"fmt"

	"github.com/DeanThompson/ginpprof"
)

// @title Simple REST API
// @version 1.0.0
// @description This is a documentation of Simple REST API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url -
// @contact.email bismobaruno@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host
// @BasePath /api/v1
func main() {
	app := routers.GetRouter()
	ginpprof.Wrap(app)
	err := app.Run(":" + conf.Configuration.Port)
	if err != nil {
		panic(fmt.Sprintf("Can't start the app: %s", err.Error()))
	}
}
