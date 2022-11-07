<template>
    <div class="container">
        <h2 class="mt-4 mb-4">Админ панель</h2>

        <h4 class="mt-4">Пользователи</h4>
        <div class="table-responsive mt-4">

            <table class="table table-striped align-middle mb-0 bg-white mt-2">
            <thead class="bg-white">
                <tr>
                <th>Пользователь</th>
                <th>Роль</th>
                <th>Действие</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(arrays, index) in datas">
                    <td>
                        <div class="d-flex align-items-center">
                        <router-link :to="`/profile/`+arrays.id">
                            <svg style="border-radius: 50%; width: 45px" viewBox="0 0 183 183" fill="none" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
                                <rect width="183" height="183" fill="url(#pattern0)"/>
                                <defs>
                                <pattern id="pattern0" patternContentUnits="objectBoundingBox" width="1" height="1">
                                <use xlink:href="#image0_9_14" transform="translate(-0.40893) scale(0.0030197 0.00259067)"/>
                                </pattern>
                                <image id="image0_9_14" width="602" height="386" xlink:href="data:image/png;base64,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"/>
                                </defs>
                            </svg>
                        </router-link>
                        <div class="ms-3">
                            <p class="fw-bold mb-1">{{arrays.first_name}} {{arrays.last_name}}</p>
                            <p class="text-muted mb-0">{{arrays.email}}</p>
                        </div>
                        </div>
                    </td>
                    <td >
                        <button
                            v-if="arrays.id != uid"
                            class="btn btn-info dropdown-toggle"
                            type="button"
                            id="dropdownMenuButton"
                            data-mdb-toggle="dropdown"
                            aria-expanded="false"
                        >
                            {{arrays.role}}
                        </button>
                        <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton">
                            <li><button class="dropdown-item" v-on:click="changeRole(arrays.id, 'student')">Студент</button></li>
                            <li><button class="dropdown-item" v-on:click="changeRole(arrays.id, 'teacher')">Преподаватель</button></li>
                            <li><button class="dropdown-item" v-on:click="changeRole(arrays.id, 'admin')">Администратор</button></li>
                        </ul>
                    </td>
                    <td>
                    <button type="submit" class="btn btn-info btn-floating"  data-mdb-toggle="modal" :data-mdb-target="`#exampleModal`+arrays.id" style="margin-bottom: 5px; margin-right: 5px;" v-on:click=""><i class="fas fa-edit"></i></button>

                    <button v-if="arrays.id != uid" type="submit" class="btn text-white btn-floating bg-danger"  style=" margin-bottom: 5px" v-on:click="deletee(arrays.id)"><i class="fas fa-trash"></i></button>
                    </td>

                    <!-- Modal -->
                    <div class="modal fade" v-for="(arr, index) in datas" :id="`exampleModal`+arr.id" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
                        <div class="modal-dialog">
                            <div class="modal-content">
                            <div class="modal-header">
                                <h5 class="modal-title" id="exampleModalLabel">Редактирование профиля</h5>
                                <button type="button" class="btn-close" data-mdb-dismiss="modal" aria-label="Close"></button>
                            </div>
                            <div class="modal-body">
                                <div class="row">
                                    <h5>Личная информация</h5>
                                    <div class="col-12 col-md-6">

                                        <div class="i-title">Имя</div>
                                        <div class="form-outline mb-4">
                                            <input type="text" id="form2Example1" v-model="change.firstName" class="form-control" autocomplete="off"/>
                                        </div>

                                    </div>

                                    <div class="col-12 col-md-6">
                                        <div class="i-title">Фамилия</div>
                                        <div class="form-outline mb-4">
                                            <input type="text" id="form2Example1" v-model="change.lastName" class="form-control" autocomplete="off"/>
                                        </div>

                                    </div>

                                    <h5>Образование</h5>
                                    <div class="col-12 col-md-6">
                                        <div class="i-title">Факультет</div>
                                        <div class="form-outline mb-4">
                                            <input type="text" id="form2Example1" v-model="change.faculty" class="form-control" autocomplete="off"/>
                                        </div>

                                        <div class="i-title">Группа</div>
                                        <div class="form-outline mb-4">
                                            <input type="text" id="form2Example1"  v-model="change.group" class="form-control" autocomplete="off"/>
                                        </div>
                                    </div>
                                    <div class="col-12 col-md-6">
                                        <div class="i-title">Форма обучения</div>
                                        <div class="form-outline mb-4">
                                            <input type="text" id="form2Example1" v-model="change.form_of_education" class="form-control" autocomplete="off"/>
                                        </div>

                                        <div class="i-title">Уровень</div>
                                        <div class="form-outline mb-4">
                                            <input type="text" id="form2Example1" v-model="change.level"  class="form-control" autocomplete="off"/>
                                        </div>
                                    </div>
                                </div>
                                <span>Оставьте поле пустым, чтобы не изменять его.</span>
                            </div>
                            <div class="modal-footer">
                                <button type="button" class="btn btn-secondary" data-mdb-dismiss="modal">Закрыть</button>
                                <button type="submit" class="btn btn-primary" v-on:click="ss(arr.id)">Сохранить</button>
                            </div>
                            </div>
                        </div>
                    </div>
                </tr>
            </tbody>
            </table>
        </div>

        <h4 class="mt-4">Новости</h4>
        <div class="table-responsive mt-4">

            <table class="table table-striped align-middle mb-0 bg-white mt-2">
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

        <h4 class="mt-4">Сообщения ректору</h4>
        <div class="table-responsive mt-4">

            <table class="table table-striped align-middle mb-0 bg-white mt-2">
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


    </div>
</template>

<script>
import axios from 'axios'
import VueJwtDecode from 'vue-jwt-decode'

export default {
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