import { createStore } from 'vuex'

export default createStore({
    state: {
        isLogged: false,
        user: {
            email: null,
            userid: null
        }
    },
    getters: {
        GETINFO: state => {
            return state
        },
        GETSTATUS: state => {
            return state.isLogged
        }
    },
    mutations: {
        loginSuccess(state, us) {
            state.isLogged = true;
            state.user.email = us.email;
            state.user.userid = us.userid;
        },
    },
    actions: {

    },
    modules: {
    }
})
