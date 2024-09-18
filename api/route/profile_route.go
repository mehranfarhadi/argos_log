package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mehranfarhadi/argos_log/api/controller"
	"github.com/mehranfarhadi/argos_log/bootstrap"
	"github.com/mehranfarhadi/argos_log/domain"
	"github.com/mehranfarhadi/argos_log/mongo"
	"github.com/mehranfarhadi/argos_log/repository"
	"github.com/mehranfarhadi/argos_log/usecase"
)

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
