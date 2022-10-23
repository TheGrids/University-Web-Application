<template>
    <div class="container">
        <div class="card bg-dark text-white mt-4">
            <img src="../assets/5.png" class="card-img aa" alt="Stony Beach"/>
            <div class="card-img-overlay text-center align-middle">
                <span class="card-titlee"><b>Связь с ректором</b></span>
                <p class="card-text">
                    <button type="button" class="btn c-q mt-3   " data-mdb-toggle="modal" data-mdb-target="#exampleModal">Написать</button>
                </p>
            </div>
        </div>

        <!-- Modal -->
        <div class="modal fade" id="exampleModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
            <div class="modal-dialog">
                <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title" id="exampleModalLabel">Связь с ректором</h5>
                    <button type="button" class="btn-close" data-mdb-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    <div class="col-12">
                        
                        <div class="i-title">Заголовок</div>
                        <div class="form-outline mb-4">
                            <input type="text" id="form2Example1" v-model="why" class="form-control" autocomplete="off"/>
                        </div>

                        <div class="i-title">Сообщение</div>
                        <div class="form-outline mb-4">
                            <textarea class="form-control" style="min-height: 100px" v-model="message"></textarea>
                        </div>

                    </div>

                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-secondary" data-mdb-dismiss="modal">Отмена</button>
                    <button type="submit" class="btn btn-primary" id="OHSUPERID" v-on:click="sendMess()">Отправить</button>
                </div>
                </div>
            </div>
        </div>

        <div class="d-flex justify-content-between">
            <h2 class="mt-4">Новости</h2>
            <div class="d-flex align-self-center">
                <button
                    class="btn c-q dropdown-toggle"
                    style="height: 50%"
                    type="button"
                    id="dropdownMenuButton"
                    data-mdb-toggle="dropdown"
                    aria-expanded="false"
                >
                    {{filter}}
                </button>
                <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton">
                    <li><button type="submit" class="dropdown-item" v-on:click="doA('Все')">Все</button></li>
                    <li><button type="submit" class="dropdown-item" v-on:click="doA('Социальная жизнь')">Социальная жизнь</button></li>
                    <li><button type="submit" class="dropdown-item" v-on:click="doA('Учебные новости')">Учебные новости</button></li>
                    <li><button type="submit" class="dropdown-item" v-on:click="doA('Жизнь ВУЗа')">Жизнь ВУЗа</button></li>
                </ul>
            </div>
        </div>

        <div class="card mt-4" v-for="(arrays, index) in list">
            <h5 class="card-header d-flex justify-content-between">
                <button type="button" class="btn btn-outline-info btn-rounded" data-mdb-ripple-color="dark">
                    <span class="sss">{{arrays.tag}}</span>
                </button> 
                <button v-if="canDel == arrays.authorId || role == 'admin'" type="submit" class="btn c-q btn-floating bg-danger"  style=" margin-bottom: 5px" v-on:click="deleteMessage(arrays.id)">
                    <i class="fas fa-trash"></i>
                </button>
            </h5>
            <div class="card-body">
                <h5 class="card-title">{{arrays.title}}</h5>
                <p class="card-text">{{arrays.body}}</p>
            </div>
            <div class="card-footer text-muted">
                <i class="far fa-calendar-plus"></i> {{new Date(arrays.time * 1000).toLocaleDateString()}} {{new Date(arrays.time * 1000).toLocaleTimeString()}} | Автор: <router-link class="text-muted" :to="`/profile/`+arrays.authorId">{{arrays.authorFirstName}} {{arrays.authorLastName}}</router-link>
            </div>
        </div>

    </div>
</template>

<style>
.btn-rounded {
    padding: 5px 10px 5px 10px;
}
.btn:hover {
    color: #54B4D3
}
.sss {
    font-size: 12px;
}
.card-img {
    object-position: center center;
    background-size: cover;  
}
.aa {
    min-height: 150px;
}
@media (min-width: 758px) {
    .card-titlee {
        font-size: 62px!important;
    }
}
.card-titlee {
    font-size: 28px;
}
.card-title {
    font-size: 28px;
}
</style>

<script>
import axios from 'axios'

export default {
  name: 'Home',
  data: function(){
    return {
        list: [],
        filter: "Все",
        message: '',
        why: '',
        canDel: localStorage.getItem('uid'),
        role: localStorage.getItem('role'),
    }
  },
  components: {
    
  },
  methods: {
    sendMess(){
        let res = {
            title: this.why,
            body: this.message
        }

        axios.post("https://universityweb.site/api/addmessage", res, {headers: {'Authorization': localStorage.getItem('accessToken')}}).then(respp => {
            if(respp.status == 200){
                this.$notify({
                    title: 'Успех',
                    type: 'success',
                    text: 'Сообщение успешно отправлено'
                })
                setTimeout(() => {
                    this.$router.go()
                }, 1000);
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

        axios.delete("https://universityweb.site/api/deletenews", {headers: {'Authorization': localStorage.getItem('accessToken')}, data: req}).then(resp => {
            if(resp.status == 200){
                this.$router.go()
            }
        }).catch(err => {
            console.log(err);
        })
    },
    doA(name){
        this.filter = name; 
        if(this.filter == 'Все') this.$router.go()
        let ress = {
            tag: this.filter
        }
        axios.post("https://universityweb.site/api/newssort", ress, {headers: {'Authorization': localStorage.getItem('accessToken')}}).then(respp => {
            if(respp.status == 200){
                this.list = respp.data.data
            }
        }).catch(err => {
            this.$notify({
                title: 'Ошибка',
                type: 'error',
                text: err
            })
        })
    }
  },
  mounted() {
    if(this.filter == 'Все'){
        axios.get("https://universityweb.site/api/news", {headers: {'Authorization': localStorage.getItem('accessToken')}}).then(respp => {
            if(respp.status == 200){
                this.list = respp.data.data;
            }
        }).catch(err => {
            this.$notify({
                title: 'Ошибка',
                type: 'error',
                text: err.response.data.msg
            })
        })
    }else {
        let res = {
            tag: this.filter
        }
        axios.get("https://universityweb.site/api/newssort", res, {headers: {'Authorization': localStorage.getItem('accessToken')}}).then(respp => {
            if(respp.status == 200){
                console.log(respp)
            }
        }).catch(err => {
            this.$notify({
                title: 'Ошибка',
                type: 'error',
                text: err.response.data.msg
            })
        })
    }
  }
}
</script>
