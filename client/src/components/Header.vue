<template>
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
        <div class="container">
            <button
                class="navbar-toggler"
                type="button"
                data-mdb-toggle="collapse"
                data-mdb-target="#navbarCenteredExample"
                aria-controls="navbarCenteredExample"
                aria-expanded="false"
                aria-label="Toggle navigation"
            >
            <i class="fas fa-bars"></i>
            </button>
            <div class="collapse navbar-collapse justify-content-between" id="navbarCenteredExample">
                <!-- Left links -->
                <ul class="navbar-nav mb-2 mb-lg-0">
                    <li class="nav-item">
                        <router-link class="nav-link active" to="/">Главная</router-link>
                    </li>
                    <li class="nav-item">
                        <router-link class="nav-link " to="/createnews">Опубликовать новость</router-link>
                    </li>
                    <li class="nav-item">
                        <router-link class="nav-link " to="/admin">Админ панель</router-link>
                    </li>
                </ul>

                <ul class="navbar-nav mb-2 mb-lg-0">
                    <li class="nav-item" style="margin-right: 10px">
                        <router-link to="/profile">
                            <button type="button" class="btn btn-info btn-floating">
                                <i class="fa fa-user-alt"></i>
                            </button>
                        </router-link>
                    </li>
                    <li class="nav-item">   
                        <button type="submit" class="btn btn-info btn-floating" v-on:click="Logout()"><i class="fas fa-sign-out-alt"></i></button>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
    {{this.$store.state}}
</template>

<script>
export default {
    name: 'Header',
    data() {
        return {

        }
    },
    methods: {
        Logout() {
            this.$store.dispatch('logout');
            this.$router.push('/registration');
        }
    }
}
</script>