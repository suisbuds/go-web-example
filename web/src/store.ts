import { createPinia, defineStore } from 'pinia' 

// 创建 pinia 实例
const pinia = createPinia()
export default pinia

interface User {
  id: number
  email: string
  roles: []
  enabled: boolean
  activated: boolean
  timezone: string
  isSuper: boolean
}

// 定义用户 store, 包含 state, getters, actions
export const useUserStore = defineStore('user', {
  // 表示用户是否认证, user 存储用户详细信息
  state: () => ({
    isAuthenticated: false,
    user: {} as User,
  }),
  // 获取用户角色名称
  getters: {
    roles: state => state.user.roles.map((role: any) => role.name),
  },
  // 设置用户认证状态并存储用户信息
  actions: {
    signin(user: User) {
      this.isAuthenticated = true
      this.user = user
    },
    signout() {
      this.isAuthenticated = false
      this.user = {} as User
      localStorage.clear()
    },
  },
})
