package router

import (
	"fmt"
	ctl "lecture/WBA-BC-Project-03/backend/controller"

	"lecture/WBA-BC-Project-03/backend/docs"

	"github.com/gin-gonic/gin"
	swgFiles "github.com/swaggo/files"
	ginSwg "github.com/swaggo/gin-swagger"
	"golang.org/x/sync/errgroup"
)

type Router struct {
	ct *ctl.Controller
}

func NewRouter(ctl *ctl.Controller) (*Router, error) {
	r := &Router{ct: ctl} //controller 포인터를 ct로 복사, 할당

	return r, nil
}

var (
	g errgroup.Group
)

// 임의 인증을 위한 함수
func liteAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c == nil {
			c.Abort() // 미들웨어에서 사용, 이후 요청 중지
			return
		}
		//http 헤더내 "Authorization" 폼의 데이터를 조회
		auth := c.GetHeader("Authorization")
		//실제 인증기능이 올수있다. 단순히 출력기능만 처리 현재는 출력예시
		fmt.Println("Authorization-word ", auth)

		c.Next() // 다음 요청 진행
	}
}

// 실제 라우팅
func (p *Router) Idx() *gin.Engine {
	//~생략
	e := gin.Default()
	e.Use(gin.Recovery())

	e.GET("/swagger/:any", ginSwg.WrapHandler(swgFiles.Handler))
	docs.SwaggerInfo.Host = "localhost:8080" //swagger 정보 등록

	dex := e.Group("", liteAuth())
	{
		// 사용자
		dex.GET("/all", p.ct.GetLiquidityList) // 1. 사용 가능한 풀 목록 조회
		// dex.GET("/my", p.ct.GetMyPools)        // 2. 자산이 예치된 풀 목록 조회
		pool := dex.Group("/pool")
		{
			pool.PUT("", p.ct.AddLiquidity)
			pool.DELETE("", p.ct.RemoveLiquidity)
			pool.GET("balance", p.ct.BalanceOf)

			//관리자
			pool.POST("", p.ct.CreateLiquidity) // 1. 새로운 풀 생성
			// pool.DELETE("", p.ct.DeleteLiquidity)  // 2. 풀 삭제
		}
		dex.POST("/swap", p.ct.Swap)

	}

	return e
}
