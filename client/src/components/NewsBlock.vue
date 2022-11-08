<template>
    <div class="d-flex justify-content-between">
        <h2 class="mt-4">Новости</h2>
        <div class="d-flex align-self-center">
            <button
                class="btn btn-info  dropdown-toggle"
                style="height: 50%"
                type="button"
                id="dropdownMenuButton"
                data-mdb-toggle="dropdown"
                aria-expanded="false"
            >
                {{this.tag}}
            </button>
            <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton">
                <li><button type="submit" class="dropdown-item" v-on:click="changeTag('Все')">Все</button></li>
                <li><button type="submit" class="dropdown-item" v-on:click="changeTag('Социальная жизнь')">Социальная жизнь</button></li>
                <li><button type="submit" class="dropdown-item" v-on:click="changeTag('Учебные новости')">Учебные новости</button></li>
                <li><button type="submit" class="dropdown-item" v-on:click="changeTag('Жизнь ВУЗа')">Жизнь ВУЗа</button></li>
            </ul>
        </div>
    </div>

    <div v-if="!list[0]" class="overlay">
        <img src="../assets/three-dots.svg" width="30">
    </div>
    <div v-if="list[0]" class="card mt-4" v-for="(arrays, index) in list">
        <h5 class="card-header d-flex justify-content-between">
            <button type="button" class="btn btn-outline-info btn-rounded" data-mdb-ripple-color="dark">
                <span class="sss">{{arrays.tag}}</span>
            </button> 
            <button v-if="this.$store.state.role == 'admin'"  v-on:click="deleteMessage(arrays.id)" type="submit" class="btn btn-danger btn-floating"  style=" margin-bottom: 5px">
                <i class="fas fa-trash"></i>
            </button>
        </h5>
        <div class="card-body">
            <h5 class="card-title">{{arrays.title}}</h5>
            <p class="card-text">{{arrays.body}}</p>
        </div>
        <div class="card-footer text-muted">
            <i class="far fa-calendar-plus"></i>  {{new Date(arrays.time * 1000).toLocaleDateString()}} {{new Date(arrays.time * 1000).toLocaleTimeString()}} | Автор: <router-link class="text-muted" :to="`/profile/`+arrays.authorId">{{arrays.authorFirstName}} {{arrays.authorLastName}}</router-link>
        </div>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'NewsBlock',
    data() {
        return {
            list: [],
            tag: 'Все'
        }
    },
    methods: {
        getAllNews(){
            axios.get("https://universityweb.site/api/news", {headers: {'Authorization': localStorage.getItem('accessToken')}}).then(respp => {
                this.list = respp.data.data;
            }).catch(err => {
                console.log(err);
            })
        },
        changeTag(tagg){
            this.tag = tagg;
            let ress = {
                tag: this.tag
            }
            if(this.tag == 'Все'){
                this.getAllNews()
            }else{
                axios.post("https://universityweb.site/api/newssort", ress, {headers: {'Authorization': localStorage.getItem('accessToken')}}).then(respp => {
                    if(respp.status == 200){
                        this.list = respp.data.data
                    }
                }).catch(err => {
                    console.log(err);
                })
            }
        },
        deleteMessage(uid) {
            let req = {
                id: uid
            }

            axios.delete("https://universityweb.site/api/deletenews", {headers: {'Authorization': localStorage.getItem('accessToken')}, data: req}).then(resp => {
                if(resp.status == 200){
                    axios.get("https://universityweb.site/api/news", {headers: {'Authorization': localStorage.getItem('accessToken')}}).then(respp => {
                        this.list = respp.data.data;
                    }).catch(err => {
                        console.log(err);
                    })
                }
            }).catch(err => {
                console.log(err);
            })
        },
    },
    mounted() {
        this.getAllNews();
    }
}
</script>