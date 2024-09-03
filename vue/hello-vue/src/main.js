import { createApp } from 'vue'
import App from './App.vue'

import {createRouter, createWebHashHistory} from "vue-router";
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'


const routers = [
    {path: "/", redirect: '/home'},
    {path: "/home", component: () => import('./components/HomePage.vue')},
    {path: "/about",  component: () => import('./components/AboutPage.vue')},
]


const router = createRouter({
    history: createWebHashHistory(),
    routes: routers,
})

const app = createApp(App)

app.use(router)
app.use(ElementPlus)
app.mount('#app')
