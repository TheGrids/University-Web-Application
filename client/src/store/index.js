import axios from 'axios';
import { createStore } from 'vuex'
import router from '../router';

export default createStore({
    state: {
        status: false,
        user: null,
        role: null
    },
    mutations: {
        loginSuccess(state, user) {
            state.status = true;
            state.role = user.role;
            state.user = user.access;
        },
        loginFailure(state) {
            state.status = false;
            state.user = null;
            state.role = null;
        },
        logout(state) {
            state.status = false;
            state.user = null;
            state.role = null;
        },
        registerSuccess(state) {
            state.status = false;
            state.role = null;
        },
        registerFailure(state) {
            state.status = false;
            state.role = null;
        }
    },
    actions: {
        async login({ commit }, user) {
            await axios.post("https://universityweb.site/api/login", {email: user.email, password: user.password}).then(resp => {
                localStorage.setItem('accessToken', resp.data.access);
                axios.get("https://universityweb.site/api/verification", {headers: {'Authorization': localStorage.getItem('accessToken')}}).then(r =>{
                    let res = {
                        access: resp.data.access,
                        role: r.headers['role']
                    }
                    commit('loginSuccess', res);
                    return Promise.resolve(user)
                }).catch(er => {
                    console.log(er);
                })
            }).catch(err => {
                commit('loginFailure');
                return Promise.reject(err.response.data.msg)
            })
        },
        logout({ commit }) {
            localStorage.removeItem('accessToken');
            commit('logout');
            router.push('/login')
        },
        async register({ commit }, user) {
            await axios.post("https://universityweb.site/api/register", user).then(resp => {
                commit('registerSuccess');
                return Promise.resolve(resp)
            }).catch(err => {
                commit('registerFailure');
                return Promise.reject(err.response.data.msg)
            })
        }
    }
})
