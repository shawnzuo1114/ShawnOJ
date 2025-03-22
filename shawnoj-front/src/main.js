import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import App from './App.vue'
import axios from 'axios'
import router from './router'

const app = createApp(App)

app.use(ElementPlus)
app.use(axios)
app.use(router)
app.mount('#app')