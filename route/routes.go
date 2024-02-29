package route

import (
	handler "registration/handler"
	"registration/logger"
	repo "registration/repo"
)

func LoggerInit(loglevel string) *logger.Logger {
	log := logger.New(loglevel)
	return log
}
func Routes(db *repo.DB, log *logger.Logger) (router *handler.Router, err1 error) {

	RegRepo := repo.NewRegRepository(db)
	RegHandler := handler.NewRegHandler(*RegRepo)

	router, err := handler.NewRouter(*RegHandler)
	return router, err

}
