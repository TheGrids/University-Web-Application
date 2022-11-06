<template>
    <div class="container">
        <h2 class="mt-4 mb-4">Создать новость</h2>

        <div class="row">
            <div>
                <div class="i-title">Название</div>
                <div class="form-outline mb-4">
                    <input type="text" class="form-control" v-model="title"/>
                </div>

                <div class="i-title">Текст</div>
                <div class="form-outline mb-4">
                    <textarea class="form-control" style="min-height: 150px" v-model="body"></textarea>
                </div>

                <div class="i-title">Выбор тега</div>
                <button
                    class="btn btn-info dropdown-toggle"
                    type="button"
                    id="dropdownMenuButton"
                    data-mdb-toggle="dropdown"
                    aria-expanded="false"
                >
                {{tag}}
                </button>
                <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton">
                    <li><button class="dropdown-item" v-on:click="this.tag = 'Социальная жизнь'">Социальная жизнь</button></li>
                    <li><button class="dropdown-item" v-on:click="this.tag = 'Учебные новости'">Учебные новости</button></li>
                    <li><button class="dropdown-item" v-on:click="this.tag = 'Жизнь ВУЗа'">Жизнь ВУЗа</button></li>
                </ul>
                
                <button type="submit" class="btn btn-info btn-block mt-4" v-on:click="create()">Создать</button>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'CreateNews',
    data() {
        return {
            tag: 'Все',
            title: '',
            body: ''
        }
    },
    methods: {
        create() {
            let res = {
                title: this.title,
                body: this.body,
                tag: this.tag
            }

            axios.post("https://universityweb.site/api/addnews", res, {headers: {'Authorization': localStorage.getItem('accessToken')}}).then(respp => {
                if(respp.status == 200){
                    this.$router.push('/')
                }
            }).catch(err => {
                console.log(err);
            })
        }
    }
}
</script>