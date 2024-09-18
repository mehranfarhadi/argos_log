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

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
