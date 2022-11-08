<template>
    <div class="container">
        <h2 class="mt-4 mb-4">Админ панель</h2>

        <!-- Tabs navs -->
        <ul class="nav nav-tabs mb-3" id="ex-with-icons" role="tablist">
            <li class="nav-item" role="presentation">
                <a class="nav-link active" id="ex-with-icons-tab-1" data-mdb-toggle="tab" href="#ex-with-icons-tabs-1" role="tab"
                aria-controls="ex-with-icons-tabs-1" aria-selected="true"><i class="fas fa-users fa-fw me-2"></i>Пользователи</a>
            </li>
            <li class="nav-item" role="presentation">
                <a class="nav-link" id="ex-with-icons-tab-2" data-mdb-toggle="tab" href="#ex-with-icons-tabs-2" role="tab"
                aria-controls="ex-with-icons-tabs-2" aria-selected="false"><i class="far fa-newspaper fa-fw me-2"></i>Новости</a>
            </li>
            <li class="nav-item" role="presentation">
                <a class="nav-link" id="ex-with-icons-tab-3" data-mdb-toggle="tab" href="#ex-with-icons-tabs-3" role="tab"
                aria-controls="ex-with-icons-tabs-3" aria-selected="false"><i class="far fa-comment fa-fw me-2"></i>Сообщения ректору</a>
            </li>
        </ul>
        <!-- Tabs navs -->

        <!-- Tabs content -->
        <div class="tab-content" id="ex-with-icons-content">
            <div class="tab-pane fade show active" id="ex-with-icons-tabs-1" role="tabpanel" aria-labelledby="ex-with-icons-tab-1">
                <AdminUsers />
            </div>
            <div class="tab-pane fade" id="ex-with-icons-tabs-2" role="tabpanel" aria-labelledby="ex-with-icons-tab-2">
                <AdminNews />
            </div>
            <div class="tab-pane fade" id="ex-with-icons-tabs-3" role="tabpanel" aria-labelledby="ex-with-icons-tab-3">
                <AdminMess />
            </div>
        </div>
        <!-- Tabs content -->


    </div>
</template>

<script>
import axios from 'axios'
import VueJwtDecode from 'vue-jwt-decode'
import AdminUsers from '../components/AdminUsers.vue'
import AdminNews from '../components/AdminNews.vue'
import AdminMess from '../components/AdminMess.vue'


export default {
  components: { AdminUsers, AdminNews, AdminMess },
    name: 'AdminPanel',
    data() {
        return {
            datas: [],
            change: {
                faculty: '',
                group: '',
                form_of_education: '',
                level: '',
                lastName: '',
                firstName: ''
            },
            uid: null,
            news: [],
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
        getUsers() {
            axios.get("https://universityweb.site/api/admin/profiles", {headers: {'Authorization': localStorage.getItem('accessToken')}}).then(respp => {
                if(respp.status == 200){
                    this.datas = respp.data.data;
                }
            }).catch(err => {
                console.log(err.response.data.msg);
            })
        },
        deletee(uid) {
            let req = {
                id: uid
            }

            axios.delete("https://universityweb.site/api/admin/person/delete", {headers: {'Authorization': localStorage.getItem('accessToken')}, data: req}).then(resp => {
                if(resp.status == 200){
                    this.getUsers()
                }
            }).catch(err => {
                console.log(err);
            })
        },
        ss(uid) {
            let req = {
                id: uid,
                faculty: this.change.faculty,
                group: this.change.group,
                form_of_education: this.change.form_of_education,
                level: this.change.level,
                last_name: this.change.lastName,
                first_name: this.change.firstName
            }

            axios.put("https://universityweb.site/api/admin/changedata", req,{headers: {'Authorization': localStorage.getItem('accessToken')}}).then(resp => {
                if(resp.status == 200){
                    this.change.faculty = '';
                    this.change.group = '';
                    this.change.form_of_education = '';
                    this.change.level = '';
                    this.change.lastName = '';
                    this.change.firstName = '';
                    this.getUsers()
                }
            }).catch(err => {
                console.log(err);
            })
        },
        changeRole(uid, rolee){ 
            let req = {
                id: uid,
                role: rolee,
            }

            axios.put("https://universityweb.site/api/admin/changerole", req, {headers: {'Authorization': localStorage.getItem('accessToken')}}).then(resp => {
                if(resp.status == 200){
                    this.getUsers()
                }
            }).catch(err => {
                console.log(err);
            })
        }
    },
    mounted() {
        let res = VueJwtDecode.decode(localStorage.getItem('accessToken'))
        this.uid = res.userid
        this.getUsers()
        this.getNews()
        this.getMessages()
    }
}
</script>