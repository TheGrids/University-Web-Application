<template>
    <div class="d-flex justify-content-center mt-5">
        <form style="width: 300px" v-on:submit.prevent="Login">
            <h2 class="text-center">Авторизация</h2>

            <div class="i-title">Email</div>
            <div class="form-outline mb-4">
                <input type="email" id="form2Example1" class="form-control" v-model="user.email"/>
            </div>
        
            <div class="i-title">Пароль</div>
            <div class="form-outline mb-4">
                <input type="password" id="form2Example2" class="form-control" v-model="user.password"/>
            </div>
        
            <div class="row mb-4">
                <div class="col">
                    <router-link to="/registration">Нет аккаунта?</router-link>
                </div>
            </div>
        
            <button type="submit" class="btn btn-info btn-block mb-4">Авторизоваться</button>
        </form>
    </div>
</template>

<script>
import VueJwtDecode from 'vue-jwt-decode'

export default {
    name: 'Login',
    data() {
        return {
            user: {
                email: '',
                password: ''
            },
            idd: null
        }
    },
    methods: {
        Login() {
            this.$store.dispatch('login', this.user).then(() => {
                if(localStorage.getItem('accessToken')){
                    this.idd = VueJwtDecode.decode(localStorage.getItem('accessToken')).userid
                }
                this.$router.push('/profile/'+this.idd).then(() => {
                    this.$router.go()
                })
            }, err => {
                console.log(err);
            })
        }
    },
    mounted() {
        if(localStorage.getItem('accessToken')){
            this.idd = VueJwtDecode.decode(localStorage.getItem('accessToken')).userid
        }
    }
}
</script>