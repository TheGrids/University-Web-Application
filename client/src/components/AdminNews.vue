<template>
    <h4 class="mt-4">Новости</h4>
    <div class="table-responsive mt-4">

        <div v-if="!news[0]" class="overlay">
            <img src="../assets/three-dots.svg" width="30">
        </div>
        <table v-if="news[0]" class="table table-striped align-middle mb-0 bg-white mt-2">
        <thead class="bg-white">
            <tr>
            <th>Новость</th>
            <th>Действие</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="(arrays, index) in news">
                <td style="max-width: 500px">
                    <div class="d-flex align-items-center">
                    <div class="ms-3">
                        <p class="fw-bold mb-1">{{arrays.title}}</p>
                        <p class="text-muted mb-0">{{new Date(arrays.time * 1000).toLocaleDateString()}} {{new Date(arrays.time * 1000).toLocaleTimeString()}}</p>
                    </div>
                    </div>
                </td>
                <td>

                <button type="submit" class="btn text-white btn-floating bg-danger"  style=" margin-bottom: 5px" v-on:click="deleteNews(arrays.id)"><i class="fas fa-trash"></i></button>
                </td>
            </tr>
        </tbody>
        </table>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: "AdminNews",
    data() {
        return {
            news: []
        }
    },
    methods: {
        getNews() {
            axios.get("https://universityweb.site/api/news", {headers: {'Authorization': localStorage.getItem('accessToken')}}).then(respp => {
                if(respp.status == 200){
                    this.news = respp.data.data;
                }
            }).catch(err => {
                this.$notify({
                    title: 'Ошибка',
                    type: 'error',
                    text: err.response.data.msg
                })
            })
        },
        deleteNews(uid) {
            let req = {
                id: uid
            }

            axios.delete("https://universityweb.site/api/deletenews", {headers: {'Authorization': localStorage.getItem('accessToken')}, data: req}).then(resp => {
                if(resp.status == 200){
                    this.getNews()
                }
            }).catch(err => {
                console.log(err);
            })
        },
    },
    mounted(){
        this.getNews()
    }

}
</script>