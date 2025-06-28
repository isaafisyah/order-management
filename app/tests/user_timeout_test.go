package tests

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/order-management/app/internal/user/handler"
	"github.com/isaafisyah/order-management/app/internal/user/response"
	"github.com/isaafisyah/order-management/app/middleware"
	"github.com/isaafisyah/order-management/app/mocks/user"
	"github.com/isaafisyah/order-management/app/utils/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserHandler_FindAll_Timeout(t *testing.T) {
    gin.SetMode(gin.TestMode)
	logger.InitLogger("test")
    start := time.Now()

    useCaseMock := new(user.MockUserUsecase)

    useCaseMock.On("FindAll", mock.Anything).Return([]response.UserResponse{
        {ID: 1, Name: "Test User", Email: "test@example.com"},
    }, nil)

    userHandler := handler.NewUserHandler(useCaseMock)

    r := gin.Default()
    r.Use(middleware.TimeoutMiddleware(5 * time.Second))
    prefixGroup := r.Group("/api/v1")
    prefixGroup.GET("/users", userHandler.FindAll)

    resp := ExecuteRequest(r, http.MethodGet, "/api/v1/users", nil)
    elapsed := time.Since(start)

    assert.GreaterOrEqual(t, elapsed.Seconds(), 5.0)
    assert.Equal(t, http.StatusGatewayTimeout, resp.Code)
}

func ExecuteRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
    req, _ := http.NewRequest(method, path, body)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    return w
}

