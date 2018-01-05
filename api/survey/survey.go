package survey

import (
	"ehelp/x/rest"
	"fmt"
	"g/x/web"

	"ehelp/o/survey"

	"github.com/gin-gonic/gin"
)

type SurveyServer struct {
	*gin.RouterGroup
	rest.JsonRender
}

func NewSurveyServer(parent *gin.RouterGroup, name string) *SurveyServer {
	var s = SurveyServer{
		RouterGroup: parent.Group(name),
	}
	s.GET("/list", s.GetAll)
	s.POST("/icon/upload", s.UploadIcon)
	s.POST("/device/add", s.AddSurveyDevice)
	s.POST("/add", s.AddSurvey)
	return &s
}

func (s *SurveyServer) AddSurvey(ctx *gin.Context) {
	var srv *survey.Survey
	web.AssertNil(ctx.ShouldBindJSON(&srv))
	fmt.Println(srv)
	web.AssertNil(srv.Create())
	s.SendData(ctx, srv)
}

func (s *SurveyServer) AddSurveyDevice(ctx *gin.Context) {
	var surveyDevice = struct {
		SurveyID string `json:"survey_id"`
		DeviceID string `json:"device_id"`
	}{}
	web.AssertNil(ctx.BindJSON(&surveyDevice))
	web.AssertNil(survey.AddDeviceToSurvey(surveyDevice.DeviceID, surveyDevice.SurveyID))
	s.Success(ctx)
}

func (s *SurveyServer) GetAll(ctx *gin.Context) {
	surveys, _ := survey.ListSurvey()
	s.SendData(ctx, surveys)
}

func (s *SurveyServer) UploadIcon(ctx *gin.Context) {
	var file, err = ctx.FormFile("icon")
	web.AssertNil(err)
	web.AssertNil(ctx.SaveUploadedFile(file, "./static/smiley/abc"))
	s.Success(ctx)
}
