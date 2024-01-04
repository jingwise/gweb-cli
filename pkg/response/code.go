package response

const (
	StringNull            = ""
	StatusSuccess         = 0      //请求成功
	StatusNotFound        = 40404  //找不到数据
	StatusUnAuth          = 40101  //用户在黑名单
	StatusUnAuthFail      = 40102  //未携带有效access_token
	StatusUnAuthFail2     = 40103  //access_token已过期
	StatusUnAuthFail3     = 404104 //IP地址错误
	StatusUnAuthFail4     = 405106 //数据查询失败
	StatusUnAuthFail5     = 406107 //并发量过大
	StatusUnAuthFail6     = 407109 //请求失败，错误未知
	StatusUnAuthFail7     = 408110 //文件过大
	StatusUnAuthFail8     = 40930  //未知错误
	StatusUnAuthFail9     = 41090  //用户信息错误
	StatusUnAuthFail10    = 41130  //密文有误或公钥有误
	StatusUnAuthFail11    = 44466  //用户不是管理员
	StatusUnAuthFail12    = 44467  //未携带access_token
	StatusUnAuthFail13    = 44468  //用户不存在
	StatusUnAuthFail14    = 44444  //access_token生成失败
	StatusUnAuthFail19    = 66666  //缺少图片
	StatusUnAuthFail20    = 666666 //缺少参数
	StatusUnAuthFail21    = 666    //你要么就别写参数，要么就给个值，给个空字符串干嘛
	StatusUnAuthFail22    = 40222  //尺寸不正确
	StatusUnAuthFail23    = 40223  //参数不正确
	StatusUnAuthFail24    = 40224  //数据不正确或未下架
	StatusUnAuthFail25    = 40225  //评论内容含有违规内容
	StatusRequestTimeout  = 40001  //请求超时
	StatusRequestNotPage  = 40006  //路径不存在
	StatusCommentTextLong = 44077  //评论文字过长
	StatusSignError       = 44078  //签名生成错误
)
