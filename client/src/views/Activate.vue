<template>
    <div class="container">
        Activate page {{this.$route.params.uuid}}
    </div>
</template>

<script>
import axios from 'axios';

export default {
    name: 'Activate',
    data: function() {
        return {
            uuid: this.$route.params.uuid
        }
    },
    mounted() {
        axios.get("https://universityweb.site/api/activate/"+this.uuid).then(resp => {
            console.log("OKAY 200")
            console.log(resp);

            this.$notify({
                title: 'Успех',
                type: 'success',
                text: resp.data.msg
            })
        }).catch(err => {
            console.log(err.response.data.msg);
            this.$notify({
                title: 'Ошибка',
                type: 'error',
                text: err.response.data.msg
            })
        })
    }
}
</script>