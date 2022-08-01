package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kzdv/api/internal/v1/overflight"
	"github.com/kzdv/api/internal/v1/user"
	"github.com/kzdv/api/pkg/logger"
)

var routeGroups map[string]func(*gin.RouterGroup)

var log = logger.Logger.WithField("component", "router/v1")

func init() {
	routeGroups = make(map[string]func(*gin.RouterGroup))
	routeGroups["/overflight"] = overflight.Routes
	routeGroups["/user"] = user.Routes
}

func SetupRoutes(r *gin.Engine) {
	log.Infof("Setting up old overflight redirect")
	// Setup redirect for old overflight endpoint
	r.GET("/live/:fac", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/v1/overflight/"+c.Param("fac"))
	})

	v1 := r.Group("/v1")
	for prefix, f := range routeGroups {
		log.Infof("Loading route prefix: %s", prefix)
		grp := v1.Group(prefix)
		f(grp)
	}
}
