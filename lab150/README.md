# [iris](https://github.com/kataras/iris)
安装:`go get -u -v github.com/kataras/iris`，这个库太大了，就不用单独的vendor了

|实验|简介|说明|
|---|---|---|
|lab001|demo| |
|lab002|tutorial| |
|lab003|websocket| |
|lab004|swagger| |

## NOTICE
 - func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
 - rp.Describe("router: %v", app.Router.BuildRouter(app.ContextPool, routerHandler, app.APIBuilder))
 - unc (h *routerHandler) HandleRequest(ctx context.Context) {
 - func (nodes Nodes) findChild(path string, params []string) (*node, []string) {