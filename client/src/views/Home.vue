<template>
    <div class="container">
        HOME PAGE<br>
        {{this.$store.getters.GETINFO.user.userid}}<br>
    </div>
</template>

<script>
export default {
  name: 'Home',
  components: {
    
  }
}
</script>
