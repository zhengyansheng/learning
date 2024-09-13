
const router = [
    {
        path: "/",
        redirect: '/02vue_template',
        name: "root"
    },
    {
        path: "/home",
        component: () => import('../components/HomePage.vue'),
        name: "home"
    },
    {
        path: "/about",
        component: () => import('../components/AboutPage.vue'),
        name: "about"
    },
    {
        path: "/menu",
        component: () => import('../components/MenuPage.vue'),
        name: "menu"
    },
    {
        path: "/axios",
        component: () => import('../components/AxiosPage.vue'),
        name: "axios"
    },
    {
        path: "/02vue_template",
        component: () => import('../components/02VueTemplate.vue'),
        name: "vue_template"
    },
    {
        path: "/03method_cal",
        component: () => import('../components/03MethodCal.vue'),
        name: "method_cal"
    },
    {
        path: "/04v_model",
        component: () => import('../components/04v-model.vue'),
        name: "v_model"
    },
    {
        path: "/05user_interactive",
        component: () => import('../components/05user-interactive.vue'),
        name: "user_interactive"
    },
    {
        path: "/06component_api",
        component: () => import('../components/06component_api.vue'),
        name: "component_api"
    },
    {
        path: "/login",
        component: () => import('../components/LoginIndex.vue'),
        name: "login"
    },
    {
        path: "/register",
        component: () => import('../components/RegisterIndex.vue'),
        name: "register"
    },
    {
        path: "/dashboard",
        component: () => import('../components/home/DashboardIndex.vue'),
        name: "dashboard"
    },
]

export default router