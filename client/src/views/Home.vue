<template>
    <div class="container">
        <!-- <div class="d-flex justify-content-between">
            <h2 class="mt-4">Связь с ректором</h2>
        </div> -->

        <div class="d-flex justify-content-between">
            <h2 class="mt-4">Новости</h2>
            
            <button
                class="btn c-q dropdown-toggle"
                style="height: 50%; margin-top: auto;margin-bottom: auto"
                type="button"
                id="dropdownMenuButton"
                data-mdb-toggle="dropdown"
                aria-expanded="false"
            >
                {{filter}}
            </button>
            <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton">
                <li><button type="submit" class="dropdown-item" v-on:click="doA('Нет')">Нет</button></li>
                <li><button type="submit" class="dropdown-item" v-on:click="doA('Социальная жизнь')">Социальная жизнь</button></li>
                <li><button type="submit" class="dropdown-item" v-on:click="doA('Учебные новости')">Учебные новости</button></li>
                <li><button type="submit" class="dropdown-item" v-on:click="doA('Жизнь ВУЗа')">Жизнь ВУЗа</button></li>
            </ul>
        </div>

        <div class="card mt-4" v-for="(arrays, index) in list">
            <h5 class="card-header"><button type="button" class="btn btn-outline-info btn-rounded" data-mdb-ripple-color="dark"><span class="sss">{{arrays.tag}}</span></button></h5>
            <div class="card-body">
                <h5 class="card-title">{{arrays.title}}</h5>
                <p class="card-text">{{arrays.body}}</p>
            </div>
            <div class="card-footer text-muted">
                <i class="far fa-calendar-plus"></i> {{new Date(arrays.time * 1000).toLocaleDateString()}} {{new Date(arrays.time * 1000).toLocaleTimeString()}}
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
</style>

<script>
import axios from 'axios'

export default {
  name: 'Home',
  data: function(){
    return {
        list: [],
        filter: "Нет"
    }
  },
  components: {
    
  },
  methods: {
    doA(name){
        this.filter = name; 
        if(this.filter == 'Нет') this.$router.go()
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
    if(this.filter == 'Нет'){
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
