import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import RegistrationPage from '../views/RegistrationPage.vue'
import success from '../views/success.vue'
import Activate from '../views/Activate.vue'
import LoginPage from '../views/LoginPage.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/registration',
    name: 'Registration',
    component: RegistrationPage
  },
  {
    path: '/login',
    name: 'LoginPage',
    component: LoginPage
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
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
