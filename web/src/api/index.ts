import request from './request'

// 导入请求模块, 发送 HTTP 请求
export default {
  // 随即一言接口
  oneSentence: () => request.get('https://v1.hitokoto.cn?c=i'),

  // auth 路由请求
  login: (data: any) => request.post('/auth/login', data), // 登录
  logout: () => request.get('/auth/logout'), // 退出
  register: (data: any) => request.post('/auth/register', data), // 注册
  userInfo: async () => await request.get('/auth/info'), // 获取用户信息认证检查

  // permission 路由请求
  initDefaultRoles: () => request.get('/api/role/default'), // 初始化默认角色
  initDefaultPermission: () => request.get('/api/permission/default'), // 初始化默认权限
  createWebobjectPermissions: (name: string) => request.get(`/api/permission/object/${name}`), // 创建对象权限

  // article 路由请求
  // 获取分类
  getCategoryOptions: async () => {
    const result = await request.get('/api/category/all')
    return (result.items || []).map((item: any) => ({
      label: item.name,
      value: item.id,
    }))
  },
  // 获取标签
  getTagOptions: async () => {
    try {
      const result = await request.get('/api/tag/all')
      return (result.items || []).map((item: any) => ({
        label: item.name,
        value: item.id,
      }))
    }
    catch (err) {
      return []
    }
  },
  // permission options (tree structure)
  // 获取权限选项
  getPermissionOptions: async () => {
    try {
      const result = await request.post('/api/permission')
      return (result.items || []).map((item: any) => ({
        label: item.name,
        value: item.id,
        children: item.children.map((sub: any) => ({
          label: sub.name,
          value: sub.id,
        })),
      }))
    }
    catch (err) {
      return []
    }
  },
}
