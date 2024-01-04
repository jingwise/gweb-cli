package response

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"log"
	"net/http"
)

var Default = &response{}

/*
* Error
*  @author: [it Chen Huawei]
*  @version[v1.0.0.1,2023-12-20 15:58:54]
*  @param c
*  @param code
*  @param err
*  @Description: 失败数据处理
 */

func Error(c *gin.Context, code int, err error) {
	log.Println("请求错误", "错误代码:", code, "错误信息:", err, c.Request.URL, c.Request.Method)
	res := Default.Clone()
	res.SetCode(code)
	res.SetMsg("No")
	res.SetData(map[string]interface{}{"errmsg": cast.ToString(code) + err.Error() + "error,hint: [1701063929056941497427837], from ip: " + c.ClientIP() + ",more info at https://dev.itcyy.cn/devtool/query?e=" + cast.ToString(code)})
	// 记录日志
	// 写入上下文
	c.Set("result", res)
	// 返回结果集
	c.AbortWithStatusJSON(http.StatusOK, res)

}

/*
* OK
*  @author: [it Chen Huawei]
*  @version[v1.0.0.1,2023-12-20 15:58:44]
*  @param c
*  @param data
*  @Description: 通常成功数据处理
 */
func OK(c *gin.Context, data interface{}) {
	log.Println("请求成功", c.Request.URL, c.Request.Method)
	res := Default.Clone()
	res.SetCode(StatusSuccess)
	res.SetData(data)
	res.SetMsg("Ok")
	c.AbortWithStatusJSON(http.StatusOK, res)
}

/*
* PageOK
*  @author: [it Chen Huawei]
*  @version[v1.0.0.1,2023-12-20 15:59:02]
*  @param c
*  @param result
*  @param count
*  @param pageIndex
*  @param pageSize
*  @Description: 分页数据处理
 */

func PageOK(c *gin.Context, result interface{}, count int, pageIndex int, pageSize int) {
	var res page
	res.List = result
	res.Count = count
	res.PageIndex = pageIndex
	res.PageSize = pageSize
	OK(c, res)
}

// Custum 兼容函数
func Custum(c *gin.Context, data gin.H) {
	c.Set("result", data)
	c.AbortWithStatusJSON(http.StatusOK, data)
}

/*
* HTMLOK
*  @author: [it Chen Huawei]
*  @version[v1.0.0.1,2023-12-20 15:58:41]
*  @param c
*  @param data
*  @Description:
 */
func HTMLOK(c *gin.Context, data interface{}) {
	logger.Info("请求成功", c.Request.URL, c.Request.Method)
	res := Default.Clone()
	res.SetCode(StatusSuccess)
	res.SetData(data)
	res.SetMsg("Ok")
	c.PureJSON(http.StatusOK, res)
}
