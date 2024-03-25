package services

import (
	"bet-manager/mocks"
	"bet-manager/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type UnitTestSuite struct {
	suite.Suite
	repository          UserService
	iUserRepositoryMock *mocks.IUserRepository
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, &UnitTestSuite{})
}

func (uts *UnitTestSuite) SetupTest() {
	repoMock := mocks.IUserRepository{}
	api := NewUserService(&repoMock)

	uts.repository = api
	uts.iUserRepositoryMock = &repoMock
}

func (uts *UnitTestSuite) TestGetAllUsers() {
	uts.iUserRepositoryMock.On("GetAll", mock.Anything).Return([]models.User{
		{
			UserId:   1,
			Email:    "Test",
			Password: "asdasdsa",
			Role:     1,
		},
	}, nil)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	status, users, ginMessage := uts.repository.GetAllUsers(ctx)

	fmt.Println(status, users, ginMessage)
	uts.Equal(1, len(*users))
	uts.Equal(200, status)
	uts.Equal(nil, ginMessage)
}

func (uts *UnitTestSuite) TestRemoveUser() {
	uts.iUserRepositoryMock.On("GetUser", mock.Anything).Return(models.User{
		UserId:   1,
		Email:    "Test",
		Password: "Test",
		Role:     1,
	}, nil)

	uts.iUserRepositoryMock.On("RemoveUser", mock.Anything).Return(nil)

	var user = &models.User{
		Email:    "Test",
		Password: "Test",
		Role:     1,
	}

	jsonBody, _ := json.Marshal(user)
	req, _ := http.NewRequest("DELETE", "/remove-user", bytes.NewBuffer(jsonBody))

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	status, _ := uts.repository.RemoveUser(ctx)
	uts.Equal(200, status)
}

func (uts *UnitTestSuite) TestRegister() {
	uts.iUserRepositoryMock.On("GetUser", mock.Anything).Return(models.User{
		UserId:   0,
		Email:    "",
		Password: "",
		Role:     0,
	}, nil)
	uts.iUserRepositoryMock.On("CreateUser", mock.Anything).Return(nil)
	var user = &models.User{
		Email:    "Test11",
		Password: "Test11",
		Role:     1,
	}

	jsonBody, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	status, _ := uts.repository.Register(ctx)

	uts.Equal(200, status)

}

func (uts *UnitTestSuite) TestLogin() {
	uts.iUserRepositoryMock.On("GetUser", mock.Anything).Return(models.User{
		UserId:   1,
		Email:    "test",
		Password: "test",
		Role:     1,
	}, nil)

	var user = &models.User{
		Email:    "test",
		Password: "test",
	}

	jsonBody, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	status, _, _, _ := uts.repository.Login(ctx)

	uts.Equal(200, status)

}

/*type UnitTestSuite struct {
	suite.Suite
	repository          ApiHandlers
	iUserRepositoryMock *mocks.IUserRepository
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, &UnitTestSuite{})
}

func (uts *UnitTestSuite) SetupTest() {
	repoMock := mocks.IUserRepository{}
	api := NewApiHandlersRepository(&repoMock)

	uts.repository = api
	uts.iUserRepositoryMock = &repoMock
}

func (uts *UnitTestSuite) TestGetAllUsers() {
	uts.iUserRepositoryMock.On("GetAll", mock.Anything).Return([]models.User{
		{
			UserId:   1,
			Email:    "Test",
			Password: "asdasdsa",
			Role:     1,
		},
	}, nil)
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	uts.repository.getAllUsers(ctx)

	var response []models.User
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		return
	}

	uts.Equal(1, len(response))
}

func (uts *UnitTestSuite) TestLogin() {
	uts.iUserRepositoryMock.On("GetUser", mock.Anything).Return(models.User{
		UserId:   1,
		Email:    "Test",
		Password: "Test",
		Role:     1,
	}, nil)

	var user = &models.User{
		Email:    "Test",
		Password: "Test",
		Role:     1,
	}
	jsonBody, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	uts.repository.login(ctx)

	var response models.User
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		return
	}

	uts.Equal("Test", response.Email)
}

func (uts *UnitTestSuite) TestRemoveUser() {
	uts.iUserRepositoryMock.On("GetUser", mock.Anything).Return(models.User{
		UserId:   1,
		Email:    "Test",
		Password: "Test",
		Role:     1,
	}, nil)

	uts.iUserRepositoryMock.On("RemoveUser", mock.Anything).Return(errors.New("test"))

	var user = &models.User{
		Email:    "Test",
		Password: "Test",
		Role:     1,
	}

	jsonBody, _ := json.Marshal(user)
	req, _ := http.NewRequest("DELETE", "/remove-user", bytes.NewBuffer(jsonBody))

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	uts.repository.removeUser(ctx)

	var response string
	fmt.Println(response)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	fmt.Println("REST")
	fmt.Println(response)
	uts.Equal(nil, err)
}*/
