<template>
    <h4 class="mt-4">Сообщения ректору</h4>
    <div class="table-responsive mt-4">

        <div v-if="!messes[0]" class="overlay">
            <img src="../assets/three-dots.svg" width="30">
        </div>
        <table v-if="messes[0]" class="table table-striped align-middle mb-0 bg-white mt-2">
        <thead class="bg-white">
            <tr>
            <th>Сообщение</th>
            <th>Действие</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="(arrays, index) in messes">
                <td style="max-width: 500px">
                    <div class="d-flex align-items-center">
                    <div class="ms-3" style="width: 100%">
                        <p class="fw-bold mb-1">{{arrays.title}}</p>
                        <p class="fw-bold mb-1">{{arrays.body}}</p>
                        <p class="text-muted mb-0">{{new Date(arrays.time * 1000).toLocaleDateString()}} {{new Date(arrays.time * 1000).toLocaleTimeString()}}</p>
                    </div>
                    </div>
                </td>
                <td>

                <button type="submit" class="btn text-white btn-floating bg-danger"  style=" margin-bottom: 5px" v-on:click="deleteMessage(arrays.id)"><i class="fas fa-trash"></i></button>
                </td>
            </tr>
        </tbody>
        </table>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: "AdminMess",
    data(){
        return {
            messes: []
        }
    },
    methods: {
        getMessages() {
            axios.get("https://universityweb.site/api/messages", {headers: {'Authorization': localStorage.getItem('accessToken')}}).then(respp => {
                if(respp.status == 200){
                    this.messes = respp.data.data;
                }
            }).catch(err => {
                this.$notify({
                    title: 'Ошибка',
                    type: 'error',
                    text: err.response.data.msg
                })
            })
        },
        deleteMessage(uid) {
            let req = {
                id: uid
            }

            axios.delete("https://universityweb.site/api/deletemessage", {headers: {'Authorization': localStorage.getItem('accessToken')}, data: req}).then(resp => {
                if(resp.status == 200){
                    this.getMessages()
                }
            }).catch(err => {
                console.log(err);
            })
        },
    },
    mounted() {
        this.getMessages()
    }

}
</script>