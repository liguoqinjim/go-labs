# [iris](https://github.com/kataras/iris)

|实验|简介|说明|
|---|---|---|
|lab001|demo| |
|lab002|tutorial| |
|lab003|websocket| |
|lab004|swagger| |
|lab005|zap,日志中间件| |
|lab006|单元测试| |
|lab007|casbin中间件| |
|lab008|testing| |
|lab009|上传文件| |
|lab010|下发文件| |
|lab011|得到所有router| |

## iris模板
 - 模板在github里面有一个专门的库，TODO
 - 之后模板可能会用到iris-cli，但是现在这个库还在测试阶段还没有完成(https://github.com/kataras/iris-cli)

## NOTICE
 - func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
 - rp.Describe("router: %v", app.Router.BuildRouter(app.ContextPool, routerHandler, app.APIBuilder))
 - unc (h *routerHandler) HandleRequest(ctx context.Context) {
 - func (nodes Nodes) findChild(path string, params []string) (*node, []string) {