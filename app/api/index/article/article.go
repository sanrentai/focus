package article

import (
	"github.com/gogf/gf/net/ghttp"
)

type C struct{}

// @summary 展示文章板块页面
// @tags    文章
// @produce html
// @router  /article [GET]
// @success 200 {string} html "页面HTML"
func (c *C) Index(r *ghttp.Request) {

}

// @summary 展示文章内容
// @tags    文章
// @produce html
// @param   id path int false "文章ID"
// @router  /article/detail/{id} [GET]
// @success 200 {string} html "页面HTML"
func (c *C) Detail(r *ghttp.Request) {

}

func (c *C) Create(r *ghttp.Request) {

}

func (c *C) Update(r *ghttp.Request) {

}

func (c *C) Delete(r *ghttp.Request) {

}
