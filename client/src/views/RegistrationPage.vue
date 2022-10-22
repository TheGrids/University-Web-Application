<template>
    <div class="container mt-5">
        <div class="row">
            <div class="d-flex justify-content-center">
                <div style="width: 300px">
                    <h2 class="text-center">Регистрация</h2>

                    <!-- Email input -->
                    <div class="form-outline mb-4">
                        <input type="email" id="form2Example1" class="form-control" autocomplete="off" v-model="user.email"/>
                        <label class="form-label" for="form2Example1">Электронная почта</label>
                    </div>
                
                    <!-- Password input -->
                    <div class="form-outline mb-4">
                        <input type="password" id="form2Example6" class="form-control" v-model="user.password"/>
                        <label class="form-label" for="form2Example6">Пароль</label>
                    </div>

                    <!-- Confirm Password input -->
                    <div class="form-outline mb-4">
                        <input type="password" id="form2Example8" class="form-control" v-model="user.password2"/>
                        <label class="form-label" for="form2Example8">Подтвердите пароль</label>
                    </div>

                    <!-- First Name input -->
                    <div class="form-outline mb-4">
                        <input type="text" id="form2Example3" class="form-control" autocomplete="off" v-model="user.first_name"/>
                        <label class="form-label" for="form2Example3">Имя</label>
                    </div>

                    <!-- Last Name input -->
                    <div class="form-outline mb-4">
                        <input type="text" id="form2Example4" class="form-control" autocomplete="off" v-model="user.last_name"/>
                        <label class="form-label" for="form2Example4">Фамилия</label>
                    </div>

                    <div class="form-check">
                        <input class="form-check-input" name="sm" type="radio" value="student" id="form1Example3" v-model="user.role"/>
                        <label class="form-check-label" for="form1Example3">Я студент</label>
                    </div>

                    <div class="form-check">
                        <input class="form-check-input" name="sm" type="radio" value="teacher" id="form1Example3"  v-model="user.role"/>
                        <label class="form-check-label" for="form1Example3">Я преподаватель</label>
                    </div>
                    <div class="row mt-2">
                        <div class="col">
                            <!-- Simple link -->
                            <router-link to="/login">Есть аккаунт?</router-link>
                        </div>
                    </div>
                    <!-- Submit button -->
                    <button type="submit" class="btn btn-primary btn-block mt-2" v-on:click="register()">Зарегистрироваться</button>
                </div>
            </div>

        </div>
    </div>

    {{user}}<br>
    {{this.$store.getters.GETINFO}}
</template>

<style>
.form-outline {
    background-color: #F6F6F6;
    border-radius: 0.25rem;
}
</style>

<script>
import axios from 'axios'
import store from '../store'

export default {
    name: 'RegistrationPage',
    data: function() {
        return {
            user: {
                email: '',
                password: '',
                password2: '',
                first_name: '',
                last_name: '',
                role: 'student'
            }
        }
    },
    methods: {
        register() {
            if(this.user.email && this.user.password && this.user.first_name && this.user.last_name){
                if(this.user.password === this.user.password2){
                    axios.post("https://universityweb.site/api/register", this.user).then(resp => {
                        console.log("OKAY 200")
                        console.log(resp);

                        this.$notify({
                            title: 'Успех',
                            type: 'success',
                            text: resp.data.msg
                        })
                        this.$router.push('/success')
                    }).catch(err => {
                        console.log(err.response.data.msg);
                        this.$notify({
                            title: 'Ошибка',
                            type: 'error',
                            text: err.response.data.msg
                        })
                    })
                }else{
                    this.$notify({
                        title: 'Ошибка',
                        type: 'error',
                        text: 'Пароли отличаются'
                    })
                }
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

<style>

</style>