<template>
    <Header />
    <router-view/>
    <notifications />
</template>

<script>
import Header from './components/Header.vue'

export default {
    components: {
        Header
    }
}
</script>
