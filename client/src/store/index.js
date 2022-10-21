import { createStore } from 'vuex'
import store from '../store'

export default createStore({
    state: {
        isLogged: false,
        token: null,
        user: null
    },
    getters: {
        GETTOKEN: state => {
            return state;
        },
    },
    mutations: {
        loginSuccess(state, user) {
            state.isLogged = true;
            state.user = user;
        },
    },
    actions: {
        LOGIN: () => {
            store.commit('loginSuccess');
        }
    },
    modules: {
    }
})
