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
                            <a href="#!">Забыли пароль?</a>
                        </div>
                    </div>
                
                    <!-- Submit button -->
                    <button type="submit" class="btn btn-primary btn-block mb-4" v-on:click="login()">Авторизоваться</button>
                </div>
            </div>

        </div>
    </div>

    {{user}} <br>
    {{this.$store.getters.GETTOKEN}}
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
                password: '',
                smth: this.$cookies.get('smth'),
                ss: VueJwtDecode.decode("")
            }
        }
    },
    methods: {
        login() {
            if(this.user.email && this.user.password){
                axios.post("https://universityweb.site/api/login", this.user).then(resp => {
                    console.log("OKAY 200")
                    console.log(resp.data.access);
                    localStorage.setItem('accessToken', resp.data.access);

                    this.$notify({
                        title: 'Успех',
                        type: 'success',
                        text: resp.data.msg
                    })
                    this.$store.commit('loginSuccess', this.user);
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