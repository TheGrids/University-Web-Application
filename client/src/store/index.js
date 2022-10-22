import { createStore } from 'vuex'

export default createStore({
    state: {
        isLogged: false,
        user: {
            email: null,
            userid: null,
            role: null
        }
    },
    getters: {
        GETINFO: state => {
            return state
        },
        GETSTATUS: state => {
            return state.isLogged
        },
        GETROLE: state => {
            return state.user.role;
        }
    },
    mutations: {
        loginSuccess(state, user) {
            state.isLogged = true;
            state.user.email = user.email;
            state.user.userid = user.userid;
            state.user.role = user.role;
        },
    },
    actions: {

    },
    modules: {
    }
})
