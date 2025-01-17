import '@unocss/reset/tailwind.css'
import 'unocss-ui/style.css'
import 'uno.css'
import './style.scss'

import { createApp } from 'vue'
import App from './App.vue'
import store from './store'
import router from './router'

// 初始化 Vue 应用, 导入全局样式, 使用状态管理和路由, 并将应用挂载到 HTML 中的 #app 元素
const app = createApp(App)
app.use(store)
app.use(router)
app.mount('#app')
