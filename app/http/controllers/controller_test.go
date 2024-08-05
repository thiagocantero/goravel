package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type ControllerTest struct {
	//Dependent services
}

func NewControllerTest() *ControllerTest {
	return &ControllerTest{
		//Inject services
	}
}

func (r *ControllerTest) Index(ctx http.Context) http.Response {
	return nil
}	
