import { createRouter, createWebHistory } from 'vue-router'
import LoginPage from '../components/LoginPage.vue'
import RegisterPage from '../components/RegisterPage.vue'
import HomePage from '../components/HomePage.vue'

const routes = [
    {
        path: '/',
        redirect: '/login', // 添加重定向到登录页
    },
    {
        path: '/login',
        name: 'Login',      // 添加路由名称
        component: LoginPage
    },
    {
        path: '/register',
        name: 'Register',   // 添加路由名称
        component: RegisterPage
    },
    {
        path: '/home',
        name: 'Home',       // 添加路由名称
        component: HomePage
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router
