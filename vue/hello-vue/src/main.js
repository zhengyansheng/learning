import { createApp } from 'vue'
import App from './App.vue'

import {createRouter, createWebHashHistory} from "vue-router";
import axios from 'axios';
import VueAxios from "vue-axios";
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'


const routers = [
    {
        path: "/",
        redirect: '/home',
        name: "root"
    },
    {
        path: "/home",
        component: () => import('./components/HomePage.vue'),
        name: "home"
    },
    {
        path: "/about",
        component: () => import('./components/AboutPage.vue'),
        name: "about"
    },
    {
        path: "/menu",
        component: () => import('./components/MenuPage.vue'),
        name: "menu"
    },
    {
        path: "/axios",
        component: () => import('./components/AxiosPage.vue'),
        name: "axios"
    },
]


const router = createRouter({
    history: createWebHashHistory(),
    routes: routers,
})

const app = createApp(App)

app.use(router)
app.use(ElementPlus)
app.use(axios,VueAxios)
app.mount('#app')
