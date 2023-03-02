package routes

import (
	"database/sql"
	"kstylehub/app/config"
	dl "kstylehub/features/dummylogin"
	ld "kstylehub/features/likereview/delivery"
	lr "kstylehub/features/likereview/repository"
	ls "kstylehub/features/likereview/service"
	md "kstylehub/features/member/delivery"
	mr "kstylehub/features/member/repository"
	ms "kstylehub/features/member/service"
	jm "kstylehub/features/middleware"
	pd "kstylehub/features/product/delivery"
	pr "kstylehub/features/product/repository"
	ps "kstylehub/features/product/service"
	rr "kstylehub/features/review/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func New(r *gin.Engine, db *sql.DB, config *config.AppConfig) {
	//repo
	memberRepo := mr.NewMemberRespository(db)
	reviewRepo := rr.NewReviewRepository(db, memberRepo)
	productRepo := pr.NewProductRespository(db, reviewRepo)
	likeReviewRepo := lr.NewLikeReviewRepository(db)

	//service
	validate := validator.New()
	productService := ps.NewProductService(productRepo)
	memberService := ms.NewMemberService(memberRepo, *validate)
	likeReviewService := ls.NewLikeReviewService(likeReviewRepo, *validate)

	//delivery
	productDelivery := pd.NewProductDelivery(productService)
	memberDelivery := md.NewMemberDelivery(memberService)
	likeReviewDelivery := ld.NewLikeReviewDelivery(likeReviewService)
	dummyAuth := dl.NewDummyLogin(db, config)

	//middleware
	jwtMiddleware := jm.NewJWTMiddleware(config)

	//route member
	r.GET("/members", memberDelivery.GetAll)
	r.POST("/members", memberDelivery.Insert)
	r.PUT("/members/:id", jwtMiddleware.Handle(memberDelivery.Update))
	r.DELETE("/members/:id", jwtMiddleware.Handle(memberDelivery.Delete))

	//route product
	r.GET("/products/:id", productDelivery.GetOneByID)

	//review
	r.POST("/reviews/:id/like", jwtMiddleware.Handle(likeReviewDelivery.Upsert))
	r.DELETE("/reviews/:id/like", jwtMiddleware.Handle(likeReviewDelivery.Delete))

	//auth
	r.POST("/auth/login", dummyAuth.Handler)
}
