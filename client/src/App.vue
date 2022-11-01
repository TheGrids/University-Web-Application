<template>
    <Header />
    <router-view/>
</template>

<style>

</style>

<script>
import Header from './components/Header.vue'
import axios from 'axios'
    
export default {
    components: {
        Header
    },
    mounted() {
        if(localStorage.getItem('accessToken')){
            axios.get("https://universityweb.site/api/verification", {headers: {'Authorization': localStorage.getItem('accessToken')}}).then(resp => {
                console.log(resp.headers['role']);
                this.$store.commit('loginSuccess', {access: localStorage.getItem('accessToken'), role: resp.headers['role']})
            }).catch(err => {
                this.$store.dispatch('logout')
            })
        }else{
            this.$store.dispatch('logout')
        }
    }
}
</script>
