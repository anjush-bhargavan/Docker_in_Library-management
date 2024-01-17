package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/anjush-bhargavan/library-management/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

var authToken = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjYsImVtYWlsIjoiY2F0aGVyaW5lYWJyYWhhbUBlbWFpbC5jb20iLCJyb2xlIjoidXNlciIsImV4cCI6MTcwNTY3NDc3OSwiaWF0IjoxNzA0ODEwNzc5fQ.qqmWD07eKJ0dAw3Fv5hm672XNZC8ADH6eRkUZ1XNiAA"

func TestProfile(t *testing.T) {
	test := []tests{
		{
			name:        "success",
			route:       "/user/profile",
			errorResult: nil,
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			getEmail = func(c *gin.Context) (string, bool) {
				return "test@email.com",true
			}
			getUser = func(val string, db *gorm.DB) (*models.User, error) {
				user := &models.User{}
				return user, nil
			}

			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			engine := gin.Default()

			RegisterUserRoutes(engine)
			req, err := http.NewRequest(http.MethodGet, tc.route, nil)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Authorization", authToken)
			engine.ServeHTTP(w, req)
			if tc.errorResult != nil {
				fmt.Println(tc.errorResult)
				errValue, _ := json.Marshal(tc.errorResult)
				require.JSONEq(t, w.Body.String(), string(errValue))
			} else {
				require.Equal(t, w.Code, 200)
			}

		})
	}
}

func TestProfileContextError(t *testing.T) {
	test := []tests{
		{
			name:        "Email error",
			route:       "/user/profile",
			errorResult: map[string]string{ "status":"Failed","message":"email not found from context",},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			getEmail = func(c *gin.Context) (string, bool) {
				return "",false
			}
			getUser = func(val string, db *gorm.DB) (*models.User, error) {
				user := &models.User{}
				return user, nil
			}

			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			engine := gin.Default()

			RegisterUserRoutes(engine)
			req, err := http.NewRequest(http.MethodGet, tc.route, nil)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Authorization", authToken)
			engine.ServeHTTP(w, req)
			if tc.errorResult != nil {
				fmt.Println(tc.errorResult)
				errValue, _ := json.Marshal(tc.errorResult)
				require.JSONEq(t, w.Body.String(), string(errValue))
			} else {
				require.Equal(t, w.Code, 200)
			}

		})
	}
}

func TestProfileUserError(t *testing.T) {
	test := []tests{
		{
			name:        "User error",
			route:       "/user/profile",
			errorResult: map[string]string{  "status":"Failed","message":"User not found",},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			getEmail = func(c *gin.Context) (string, bool) {
				return "test@email.com",true
			}
			getUser = func(val string, db *gorm.DB) (*models.User, error) {
				user := &models.User{}
				return user, errors.New("mocked error")
			}

			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			engine := gin.Default()

			RegisterUserRoutes(engine)
			req, err := http.NewRequest(http.MethodGet, tc.route, nil)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Authorization", authToken)

			engine.ServeHTTP(w, req)
			if tc.errorResult != nil {
				fmt.Println(tc.errorResult)
				errValue, _ := json.Marshal(tc.errorResult)
				require.JSONEq(t, w.Body.String(), string(errValue))
			} else {
				require.Equal(t, w.Code, 200)
			}

		})
	}
}

func TestUpadateProfile(t *testing.T) {
	test := []tests {
		{
		name : "success",
		body: models.User{
			FirstName: "test",
			LastName:  "name",
			DoB:       "01/01/2000",
			Gender:    "M",
			Email:     "test@gmail.com",
			Phone:     "7521750433",
			Address:   "15/2, Oz Villa,New California",
			Password:  "12345",
		},
		route : "/user/profile",
		errorResult : nil,
		},
	}

	for _,tc := range test {
		t.Run(tc.name,func(t *testing.T) {
			getEmail = func(c *gin.Context) (string, bool) {
				return "test@email.com",true
			}
			getUser = func(val string, db *gorm.DB) (*models.User, error) {
				user := &models.User{}
				return user, nil
			}
			updateUser = func(db *gorm.DB) error {
				return nil
			}
			body, err := json.Marshal(tc.body)
			if err != nil {
				require.NoError(t,err)
			}
			r := strings.NewReader(string(body))

			gin.SetMode(gin.TestMode) // we are setting gin in test mode

			engine := gin.Default() //starting a gin engine

			RegisterUserRoutes(engine) //using the engine to pass in user routes to 

			w := httptest.NewRecorder()
			req,err := http.NewRequest(http.MethodPut,tc.route,r)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Authorization",authToken)

			engine.ServeHTTP(w,req)

			if tc.errorResult != nil {
				errValue,_ := json.Marshal(tc.errorResult)
				require.JSONEq(t,w.Body.String(),string(errValue))
			}else{
				require.Equal(t,w.Code,200)
			}

		})
	}
}

type password struct {
	Old  string 
	New  string 
	CNew string 
}

func TestChangePassword(t *testing.T) {
	test := [] tests {
		{
			name : "success",
			body: password{
				Old  : "12345",
				New  : "123",
				CNew : "123",
			},
			route: "/user/profile/changepassword",
			errorResult: nil,
		},
	}

	for _,tc := range test {
		t.Run(tc.name,func(t *testing.T) {
			getEmail = func(c *gin.Context) (string, bool) {
				return "test@email.com",true
			}
			getUser = func(val string, db *gorm.DB) (*models.User, error) {
				return &globalUser,nil
			}
			checkPassword = func(user *models.User,providedPassword string) error {
				return nil
			}
			hashPassword = func(user *models.User,password string) error {
				return nil
			}
			updateUser = func(db *gorm.DB) error {
				return nil
			}

			body,err := json.Marshal(tc.body) 
			if err != nil {
				require.NoError(t,err)
			}

			r := strings.NewReader(string(body))

			gin.SetMode(gin.TestMode)

			engine := gin.Default()

			RegisterUserRoutes(engine)

			w := httptest.NewRecorder()
			req,err := http.NewRequest(http.MethodPatch,tc.route,r)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Authorization",authToken)
			engine.ServeHTTP(w,req)

			if tc.errorResult != nil {
				errValue,_ := json.Marshal(tc.errorResult)
				require.JSONEq(t,w.Body.String(),string(errValue))
				
			}else{
				require.Equal(t,w.Code,200)
			}
		})
	}
}


func TestHistory(t *testing.T) {
	test := []tests{
		{
			name:        "success",
			route:       "/user/profile/viewhistory",
			errorResult: nil,
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			getUserID = func(c *gin.Context) (uint64, bool) {
				return 1,false
			}
			getHistory = func(val uint64, db *gorm.DB) (*[]models.History, error) {
				return &[]models.History{},nil
			}

			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			engine := gin.Default()

			RegisterUserRoutes(engine)
			req, err := http.NewRequest(http.MethodGet, tc.route, nil)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Authorization", authToken)
			engine.ServeHTTP(w, req)
			if tc.errorResult != nil {
				errValue, _ := json.Marshal(tc.errorResult)
				require.JSONEq(t, w.Body.String(), string(errValue))
			} else {
				require.Equal(t, w.Code, 200)
			}

		})
	}
}


func TestViewPlan(t *testing.T) {
	test := []tests {
		{
			name: "success",
			route: "/user/myplan",
			errorResult: nil,
		},
	}

	for _,tc := range test {
		t.Run(tc.name,func(t *testing.T) {
			getUserID = func(c *gin.Context) (uint64, bool) {
				return 1, true
			}
			getMembership = func(val uint64, db *gorm.DB) (*models.Membership, error) {
				return &models.Membership{},nil
			}

			gin.SetMode(gin.TestMode)

			engine := gin.Default()

			w := httptest.NewRecorder()

			RegisterUserRoutes(engine)

			req,err := http.NewRequest(http.MethodGet,tc.route,nil)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Authorization",authToken)
			engine.ServeHTTP(w,req)
			if tc.errorResult != nil {
				errValue,_ := json.Marshal(tc.errorResult)
				require.JSONEq(t,w.Body.String(),string(errValue))
			}else{
				require.Equal(t,w.Code,200)
			}
		})
			
		
	}
}