package network

import (
	"fmt"
	"github.com/3boku/community-forum/types"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	forumRouterInit     sync.Once
	forumRouterInstance *forumRouter
)

type forumRouter struct {
	r *Network
}

func newForumRouter(r *Network) *forumRouter {
	forumRouterInit.Do(func() {
		forumRouterInstance = &forumRouter{
			r: r,
		}

		r.registerCreate("/", forumRouterInstance.create)
		r.registerGet("/", forumRouterInstance.get)
		r.registerDelete("/", forumRouterInstance.delete)
		r.registerUpdate("/", forumRouterInstance.update)
	})

	return forumRouterInstance
}

func (f *forumRouter) create(c *gin.Context) {
	fmt.Println("create")

	f.r.okResponse(c, &types.ForumResponse{
		ApiResponse: &types.ApiResponse{
			Result:      1,
			Description: "성공",
		},
		Forum: nil,
	})
}

func (f *forumRouter) get(c *gin.Context) {
	fmt.Println("get")
}

func (f *forumRouter) update(c *gin.Context) {
	fmt.Println("update")
}

func (f *forumRouter) delete(c *gin.Context) {
	fmt.Println("delete")
}
