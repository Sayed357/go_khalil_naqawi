package controllers_test

import (
	"go-todos-api/controllers"
	"go-todos-api/services/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestDeleteTodoSuccess(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/todos/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockService := new(mocks.MockTodoService)
	mockService.On("Delete", "1").Return(nil)

	controller := controllers.TodoController{Service: mockService} // Use the mock struct
	err := controller.Delete(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}
