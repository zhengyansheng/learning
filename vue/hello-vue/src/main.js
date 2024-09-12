import { createApp } from 'vue'
import App from './App.vue'

import {createRouter, createWebHashHistory} from "vue-router";
import axios from 'axios';
import VueAxios from "vue-axios";
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import mainRouter from "./routers/router.js"

// 定义路由
const router = createRouter({
    history: createWebHashHistory(),
    routes: mainRouter,
})

const app = createApp(App)

app.use(router)
app.use(ElementPlus)
app.use(axios,VueAxios)
app.mount('#app')
