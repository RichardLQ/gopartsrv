package mini

import (
	"github.com/gin-gonic/gin"
	"gopartsrv/condition/controller/mini"
)

//首页路由
func MiniRouter(e *gin.Engine) {
	v1 := e.Group("/mini")
	{

		v1.GET("/banner", mini.Banner) //导航栏
		v1.GET("/getOpenid", mini.GetOpenid) //获取openid
		v1.GET("/getToken", mini.GetToken) //获取token
		v1.GET("/getScriptCmd", mini.GetScriptCmd) //执行cmd命令
		v1.GET("/updateUserInfo", mini.UpdateUserInfo) //执行cmd命令
		v1.GET("/addCollectInfo", mini.AddCollectInfo) //搜集
		v1.POST("/uploadImage", mini.UploadImage)//上传图片
		v1.GET("/getuploadImage", mini.GetUploadImage)//查看上传图片
		v1.GET("/deluploadImage", mini.DelUploadImage)//删除上传图片

		// 心路历程
		v1.Any("/getPicUrl", mini.GetPicUrl) //获取图片地址
		v1.GET("/getRecommend", mini.Recommend) //热门推荐
		v1.GET("/getClassification", mini.Classification) //分类
		v1.GET("/getClassifiedContent", mini.ClassifiedContent) //分类内容
		v1.GET("/updateComplaints", mini.UpdateComplaints) //分类内容
		v1.GET("/getSearch", mini.GetSearch) //搜索图片
		v1.GET("/addNoteContent", mini.AddNoteContent) //添加记事本
		v1.GET("/getNoteContent", mini.GetNoteContent) //查询记事本
		v1.GET("/delNoteContent", mini.DelNoteContent) //删除记事本
		//记忆流沙
		v1.GET("/getReview", mini.Review) //视频内容
		v1.Any("/getCompanyInfo", mini.CompanyInfo) //公司信息
		v1.Any("/getHotImageAndVideo", mini.GetHotImageAndVideo) //获取热门推荐
		v1.Any("/getTypeVideo", mini.GetTypeVideo) //舞蹈类型内容
		v1.Any("/getTeacherInfo", mini.GetTeacherInfo) //获取老师信息

	}
}
