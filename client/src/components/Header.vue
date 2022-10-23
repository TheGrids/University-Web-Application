<template>
    <nav class="navbar navbar-expand-lg navbar-light bb">
        <div class="container">
            <button
                class="navbar-toggler text-white"
                type="button"
                data-mdb-toggle="collapse"
                data-mdb-target="#navbarCenteredExample"
                aria-controls="navbarCenteredExample"
                aria-expanded="false"
                aria-label="Toggle navigation"
            >
            <i class="fas fa-bars"></i>
            </button>
            <div class="collapse navbar-collapse justify-content-between" style="margin-left: 10px" id="navbarCenteredExample">
                <!-- Left links -->
                <ul class="navbar-nav mb-2 mb-lg-0">
                    <li class="nav-item active">
                        <router-link class="nav-link  text-white" to="/">Главная</router-link>
                    </li>
                    <li class="nav-item active" v-if="this.smth || this.hehe">
                        <router-link class="nav-link  text-white" to="/createnews">Опубликовать новость</router-link>
                    </li>
                    <li class="nav-item active" v-if="this.smth">
                        <router-link class="nav-link  text-white" to="/admin">Админ панель</router-link>
                    </li>
                </ul>

                <ul v-if="this.$store.getters.GETSTATUS" class="navbar-nav mb-2 mb-lg-0 d-flex flex-row ">
                    <li class="nav-item" style="margin-right: 8px">
                        <router-link style="color: white" :to="`/profile/`+this.$store.getters.GETINFO.user.userid">
                            <button type="button" class="btn c-q btn-floating">
                                <i class="fa fa-user-alt"></i>
                            </button>
                        </router-link>
                    </li>
                    <li class="nav-item">
                        <button type="submit" class="btn c-q btn-floating" v-on:click="logout()"><i class="fas fa-sign-out-alt"></i></button>
                    </li>
                </ul>
                <ul v-if="!this.$store.getters.GETSTATUS" class="navbar-nav mb-2 mb-lg-0">
                    <li class="nav-item">
                        <router-link class="btn c-q btn-floating" to="/login"><i class="fas fa-sign-in-alt"></i></router-link>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
</template>
<style>
    .c-q {
        background-color: #2496A0;
        color: white;
    }
    .bb {
        background-color: #3F3558;
    }
    .nav-item {
        margin-right: 10px;
    }
</style>
<script>
import axios from 'axios'

export default {
    name: 'Header',
    props: ['smth', 'hehe'],
    data: function( ){
        return {
            admin: false
        }
    },
    methods: {
        logout() {
            localStorage.removeItem('accessToken');
            localStorage.removeItem('uid');
            localStorage.removeItem('role');
            this.$store.state.isLogged = false;
            this.$store.state.user = {
                email: null,
                userid: null,
                role: null
            }
            this.$router.go()
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>

</style>
