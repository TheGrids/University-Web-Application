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
        // if(this.$store.getters.GETSTATUS){
        //     if(localStorage.getItem('accessToken')){

        //         let user = VueJwtDecode.decode(localStorage.getItem('accessToken'))

        //         axios.get("https://universityweb.site/api/verification", {headers: {Authorization: localStorage.getItem('accessToken')}}).then(resp => {
        //             console.log(resp)
        //         }).catch(err => {
        //             console.log(err.response.data.msg);
        //         })
        //     }else {
        //         console.log("ERR SMTH")
        //     }
        // }
    }
}
</script>
