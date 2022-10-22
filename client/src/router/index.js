import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import RegistrationPage from '../views/RegistrationPage.vue'
import success from '../views/success.vue'
import Activate from '../views/Activate.vue'
import LoginPage from '../views/LoginPage.vue'
import AdminPanel from '../views/AdminPanel.vue'
import PageNotFound from '../views/PageNotFound.vue'
import ProfilePage from '../views/ProfilePage.vue'
import store from '../store'

let authToken = localStorage.getItem('accessToken');
let authGuard = function(to, from, next) {
    if(!authToken) {
        next({name: 'Registration'})
    }else{
        next()
    }
}
let managerRegLog = function(to, from, next) {
    if(authToken) {
        next({name: 'Home'})
    }else{
        next()
    }
}
let managerAdminGuard = function(to, from, next) {
    if(!authToken) {
        next({name: 'LoginPage'})
    }else if(localStorage.getItem('role') != 'admin'){
        next( { name: 'Home' })
    }else {
        next()
    }
}


const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home,
        beforeEnter: authGuard
    },
    {
        path: '/registration',
        name: 'Registration',
        component: RegistrationPage,
        beforeEnter: managerRegLog
    },
    {
        path: '/login',
        name: 'LoginPage',
        component: LoginPage,
        beforeEnter: managerRegLog
    },
    {
        path: '/success',
        name: 'success',
        component: success
    },
    {
        path: '/activate/:uuid',
        name: 'Activate',
        component: Activate
    },
    {
        path: '/admin',
        name: 'AdminPanel',
        component: AdminPanel,
        beforeEnter: managerAdminGuard
    },
    {
        path: '/profile/:uid',
        name: 'ProfilePage',
        component: ProfilePage,
        beforeEnter: authGuard
    },
    { 
        path: "/:pathMatch(.*)*", 
        component: PageNotFound 
    }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
