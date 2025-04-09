package initiator

import (
	"github.com/BoruTamena/jobboard/internal/handler"
	"github.com/BoruTamena/jobboard/internal/handler/application"
	"github.com/BoruTamena/jobboard/internal/handler/auth"
	"github.com/BoruTamena/jobboard/internal/handler/job"
	"github.com/BoruTamena/jobboard/internal/handler/user"
)

type Handler struct {
	authHandler        handler.Auth
	userProfileHandler handler.UserProfile
	jobHandler         handler.Job
	jobApplication     handler.JobApplication
}

func InitHandler(md Module) Handler {

	return Handler{
		authHandler:        auth.AuthHandler(md.authModule),
		userProfileHandler: user.InitUserProfileHandler(md.userProfileModule),
		jobHandler:         job.NewJobHandler(md.jobModule),
		jobApplication:     application.NewJobApplication(md.applicationModule),
	}
}
