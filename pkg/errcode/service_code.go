package errcode

// 业务模块错误码

var (
	ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(20010002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(20010003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(20010004, "删除标签失败")
	ErrorCountTagFail   = NewError(20010005, "统计标签失败")
)
var (
	ErrorGetArticleFail    = NewError(20020001, "获取单个文章失败")
	ErrorGetArticlesFail   = NewError(20020002, "获取多个文章失败")
	ErrorCreateArticleFail = NewError(20020003, "创建文章失败")
	ErrorUpdateArticleFail = NewError(20020004, "更新文章失败")
	ErrorDeleteArticleFail = NewError(20020005, "删除文章失败")
)

var (
	ErrorCreateUserFail    = NewError(20030001, "创建用户失败")
	ErrorGetUserFail       = NewError(20030002, "获取用户失败")
	ErrorUpdateUserFail    = NewError(20030003, "更新用户失败")
	ErrorDeleteUserFail    = NewError(20030004, "删除用户失败")
	ErrorGetUserListFail   = NewError(20030005, "获取用户列表失败")
	ErrorCheckUserFail     = NewError(20030006, "验证用户失败")
	ErrorCheckUsernameFail = NewError(20030007, "验证用户名是否存在失败")
	ErrorGetUserInfoFail = NewError(20030008, "获取用户信息失败")
)

var (
	ErrorUploadFileFail = NewError(20040001, "上传文件失败")
)
