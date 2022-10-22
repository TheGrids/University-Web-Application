<template>
    <Header  :smth="ok"/>
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
    data: function() {
        return {
            ok: false
        }
    },
    mounted() {
        if(!this.$store.getters.GETSTATUS){
            if(localStorage.getItem('accessToken')){
                axios.get("https://universityweb.site/api/verification", {headers: {'Authorization': localStorage.getItem('accessToken')}}).then(respp => {
                    if(respp.status == 200){
                        let us = VueJwtDecode.decode(localStorage.getItem('accessToken'))
                        localStorage.setItem('role', respp.headers['role']);
                        if(respp.headers['role'] == 'admin') this.ok = true
                        localStorage.setItem('uid', us.userid);
                        localStorage.setItem('accessToken', localStorage.getItem('accessToken'));
                        
                        this.$store.commit('loginSuccess', us)
                    }
                }).catch(err => {
                    console.log(err.response.data.msg);
                    localStorage.removeItem('accessToken');
                    localStorage.removeItem('uid');
                    localStorage.removeItem('role');
                })
            }else {
                console.log("NO TOKEN")
            }
        }else{
            console.log("NO ALL")
        }
    }
}
</script>
