package controllers

import (
	"github.com/anjush-bhargavan/library-management/middleware"
	"github.com/gin-gonic/gin"
)

//RegisterUserRoutes will implement the user routes
func RegisterUserRoutes(r *gin.Engine) {
	r.GET("/",IndexPage)
	r.POST("/login", UserLogin)
	r.POST("/signup", UserSignup)
	r.POST("/verifyotp", VerifyOTP)

	userGroup := r.Group("/user")
	userGroup.Use(middleware.Authorization("user"))
	{
		userGroup.GET("/home", HomePage)

		userGroup.GET("/home/book/:id", GetBook)
		userGroup.GET("/home/books", ViewBooks)
		userGroup.GET("/search",SearchBooks)
		userGroup.GET("book/category/:id",BookByCategory)
		userGroup.GET("/book/author/:id",BookByAuthor)
		userGroup.GET("/book/sort",SortByRating)

		userGroup.GET("/profile", UserProfile)
		userGroup.PUT("/profile", ProfileUpdate)
		userGroup.PATCH("/profile/changepassword", ChangePassword)
		userGroup.GET("/myplan",ViewMyPlan)

		userGroup.GET("profile/plans", ShowPlans)
		userGroup.POST("profile/plans", GetPlan)

		userGroup.GET("/profile/viewfine",ViewFine)
		r.GET("/profile/payfine",PayFine)
		
		userGroup.GET("/profile/viewhistory", ViewHistory)

		userGroup.POST("/wishlist/:id", AddToWishlist)
		userGroup.GET("/wishlist", ShowWishlist)
		userGroup.DELETE("/wishlist/:id", DeleteWishlist)

		r.GET("/profile/membership", Membership)
		r.GET("/payment/success", RazorpaySuccess)
		r.GET("/success", SuccessPage)
		r.GET("/invoice/download",InvoiceDownload)

		userGroup.GET("/checkout/:id", DeliveryDetails)
		userGroup.POST("/checkout", Delivery)
		userGroup.PATCH("/cancel",CancelOrder)
		userGroup.GET("/inhold",InholdBook)
		userGroup.POST("/return",ReturnBook)

		userGroup.POST("/review/:id",AddReview)
		userGroup.GET("/review/:id",ShowReview)
		userGroup.PATCH("/review/:id",EditReview)
		userGroup.DELETE("/review/:id",DeleteReview)

		userGroup.POST("/feedback",Feedback)
		userGroup.GET("/events",ViewEvents)

	}
}