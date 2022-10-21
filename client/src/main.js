import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import Notifications from '@kyvg/vue3-notification'
import VueCookies from 'vue-cookies';



createApp(App).use(store).use(Notifications).use(VueCookies).use(router).mount('#app')
