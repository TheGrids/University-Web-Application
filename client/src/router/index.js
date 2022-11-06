import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Profile from '../views/Profile.vue'
import Login from '../views/Login.vue'
import Registration from '../views/Registration.vue'
import success from '../views/success.vue'
import CreateNews from '../views/CreateNews.vue'
import AdminPanel from '../views/AdminPanel.vue'
import store from '../store'
let authToken = localStorage.getItem('accessToken');

let managerAdminGuard = function(to, from, next) {
    if(!authToken) {
        next('/login')
    }else if(store.state.role != 'admin'){
        next('/')
    }else {
        next()
    }
}

let managerNewsGuard = function(to, from, next) {
    if(!authToken) {
        next('/login')
    }else if(store.state.role != 'admin' && store.state.role != 'teacher'){
        next('/')
    }else {
        next()
    }
}

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home
    },
    {
        path: '/profile',
        name: 'Profile',
        component: Profile
    },
    {
        path: '/registration',
        name: 'Registration',
        component: Registration
    },
    {
        path: '/login',
        name: 'Login',
        component: Login
    },
    {
        path: '/success',
        name: 'success',
        component: success
    },
    {
        path: '/createnews',
        name: 'CreateNews',
        component: CreateNews,
        beforeEnter: managerNewsGuard
    },
    {
        path: '/admin',
        name: 'AdminPanel',
        component: AdminPanel,
        beforeEnter: managerAdminGuard
    }
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

router.beforeEach((to, from, next) => {
    const publicPages = ['/login', '/registration', '/'];
    const authRequired = !publicPages.includes(to.path);
    const loggedIn = localStorage.getItem('accessToken');

    if (authRequired && !loggedIn) {
        next('/login');
    } else {
        next();
    }
})

export default router
