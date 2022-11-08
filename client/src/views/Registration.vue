<template>
    <div class="d-flex justify-content-center mt-5">
        <form style="width: 300px" v-on:submit.prevent="Register">
            <h2 class="text-center">Регистрация</h2>

            <!-- Email input -->
            <div class="i-title">Email</div>
            <div class="form-outline mb-4">
                <input type="email" id="form2Example1" class="form-control" v-model="user.email"/>
            </div>
        
            <!-- Password input -->
            <div class="i-title">Пароль</div>
            <div class="form-outline mb-4">
                <input type="password" id="form2Example6" class="form-control" v-model="user.password"/>
            </div>

            <!-- Confirm Password input -->
            <div class="i-title">Подтвердите пароль</div>
            <div class="form-outline mb-4">
                <input type="password" id="form2Example8" class="form-control" v-model="user.password2"/>
            </div>

            <!-- First Name input -->
            <div class="i-title">Имя</div>
            <div class="form-outline mb-4">
                <input type="text" id="form2Example3" class="form-control" v-model="user.first_name"/>
            </div>

            <!-- Last Name input -->
            <div class="i-title">Фамилия</div>
            <div class="form-outline mb-4">
                <input type="text" id="form2Example4" class="form-control" v-model="user.last_name"/>
            </div>

            <div class="form-check">
                <input class="form-check-input" name="sm" type="radio" value="student" v-model="user.role"/>
                <label class="form-check-label" for="form1212">Я студент</label>
            </div>
            
            <div class="form-check">
                <input class="form-check-input" name="sm" type="radio" value="teacher" v-model="user.role"/>
                <label class="form-check-label" for="form1212">Я преподаватель</label>
            </div>
            <div class="row mt-2">
                <div class="col">
                    <!-- Simple link -->
                    <router-link to="/login">Есть аккаунт?</router-link>
                </div>
            </div>
            <!-- Submit button -->
            <button type="submit" class="btn btn-info btn-block mt-2">Зарегистрироваться</button>
        </form>
    </div>
</template>

<script>
export default {
    name: 'Registration',
    data() {
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
        Register(){
            if(this.user.email && this.user.password && this.user.first_name && this.user.last_name){
                if(this.user.password != this.user.password2){
                    console.log('Пароли лол');
                }else {
                    this.$store.dispatch('register', this.user).then(() => {
                        this.$router.push('/success')
                    }, err => {
                        console.log(err);
                    })
                }
            }
        }
    }
}
</script>