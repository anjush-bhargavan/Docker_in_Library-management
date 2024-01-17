package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/anjush-bhargavan/library-management/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestFeedback(t *testing.T) {
	test := []tests {
		{
			name: "success",
			body: models.FeedBack{
				Subject: "good",
				Content: "variety of books",
			},
			route: "/user/feedback",
			errorResult: nil,
		},
		{
			name: "Validation error",
			body: models.FeedBack{
				Subject: "",
				Content: "variety of books",
			},
			route: "/user/feedback",
			errorResult: map[string]string{"status":"Failed",
			"message":"Please fill all fields"},
		},
	}

	for _,tc := range test {
		t.Run(tc.name,func(t *testing.T) {
			getUserID = func(c *gin.Context) (uint64, bool) {
				return 1, true
			}
			createFeedback = func(feedback *models.FeedBack, db *gorm.DB) error {
				return nil
			}

			body,err := json.Marshal(tc.body)
			if err != nil {
				require.NoError(t, err)
			}

			gin.SetMode(gin.TestMode)

			engine := gin.Default()

			RegisterUserRoutes(engine)

			r := strings.NewReader(string(body))

			w := httptest.NewRecorder()

			req,err := http.NewRequest(http.MethodPost,tc.route,r)
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
