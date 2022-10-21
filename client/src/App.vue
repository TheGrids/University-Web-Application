<template>
    <Header />
    <router-view/>
    <notifications />
</template>

<script>
import Header from './components/Header.vue'
import VueJwtDecode from 'vue-jwt-decode'
import axios from 'axios'

export default {
    components: {
        Header
    },
    mounted() {
        console.log(localStorage.getItem('accessToken'))
        if(!this.$store.getters.GETTOKEN.isLogged){
            if(localStorage.getItem('accessToken')){

                let user = VueJwtDecode.decode(localStorage.getItem('accessToken'))

                this.$store.commit('loginSuccess', user);
            }else if(this.$cookies.get('refresh_token')){
                axios.post("https://universityweb.site/api/refresh/").then(resp => {
                    console.log("OKAY 200")
                    console.log(resp);
                }).catch(err => {
                    console.log("ERR SMTH")
                    console.log(err);
                })
            }
        }
    }
}
</script>
