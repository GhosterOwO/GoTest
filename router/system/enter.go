package system

type RouterGroup struct {
	LoginRouter
	UserRouter
	ImgRouter
}

var RouterGroupApp = new(RouterGroup)
