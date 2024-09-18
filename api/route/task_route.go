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

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
