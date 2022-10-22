<template>
    <div class="container mt-5">
        <div class="row">
            <div class="d-flex justify-content-center">
                <div style="width: 300px">
                    <h2 class="text-center">Авторизация</h2>

                    <!-- Email input -->
                    <div class="form-outline mb-4">
                        <input type="email" id="form2Example1" class="form-control" v-model="user.email" autocomplete="off"/>
                        <label class="form-label" for="form2Example1">Электронная почта</label>
                    </div>
                
                    <!-- Password input -->
                    <div class="form-outline mb-4">
                        <input type="password" id="form2Example2" class="form-control" v-model="user.password"/>
                        <label class="form-label" for="form2Example2">Пароль</label>
                    </div>
                
                    <!-- 2 column grid layout for inline styling -->
                    <div class="row mb-4">
                        <div class="col">
                            <!-- Simple link -->
                            <router-link to="/registration">Нет аккаунта?</router-link>
                        </div>
                    </div>
                
                    <!-- Submit button -->
                    <button type="submit" class="btn btn-primary btn-block mb-4" v-on:click="login()">Авторизоваться</button>
                </div>
            </div>

        </div>
    </div>

    {{user}} <br>
    {{this.$store.getters.GETINFO}}
</template>

<script>
import axios from 'axios'
import VueJwtDecode from 'vue-jwt-decode'


export default {
    name: 'LoginPage',
    data: function() {
        return {
            user: {
                email: '',
                password: ''
            }
        }
    },
    methods: {
        login() {
            if(this.user.email && this.user.password){
                axios.post("https://universityweb.site/api/login", this.user).then(resp => {
                    console.log("OKAY 200")
                    console.log(resp.data.access);

                    axios.get("https://universityweb.site/api/verification", {headers: {'Authorization': resp.data.access}}).then(resp => {
                        if(resp.status == 200){
                            localStorage.setItem('role', resp.headers['role']);
                            console.log("OKAY")
                        }
                    }).catch(err => {
                        console.log(err.response.data.msg);
                    })

                    //this.$router.go()
                }).catch(err => {
                    console.log(err.response.data.msg);
                    this.$notify({
                        title: 'Ошибка',
                        type: 'error',
                        text: err.response.data.msg
                    })
                })
            }else {
                this.$notify({
                    title: 'Ошибка',
                    type: 'error',
                    text: 'Введите данные'
                })
            }
        }
    }
}
</script>