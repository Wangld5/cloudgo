package service
import(
	"net/http"
	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)
func initRouter(r *mux.Router, formatter *render.Render) {
	//handleFunc第一个参数为请求路径，第二个参数为处理函数
	r.HandleFunc("/hello/{name}", helloHandler(formatter)).Methods("GET")
}

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options {
	    IndentJSON: true,
	})
	//创建一个路由
	r := mux.NewRouter()
	server := negroni.Classic()
	//初始化路由
	initRouter(r, formatter)
	server.UseHandler(r)
	return server
}
//请求处理函数，这是一个函数闭包，返回http.HandlerFunc类型。
func helloHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)//获得请求
		name := vars["name"]//解析出请求的内容
		formatter.JSON(w, http.StatusOK, map[string]string{"name": name})//模板渲染
	}
}