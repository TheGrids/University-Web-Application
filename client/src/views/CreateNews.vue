<template>
    <div class="container">
        <h2 class="mt-4 mb-4">Создать новость</h2>

        <div class="row">
            <div>
                <div class="i-title">Название</div>
                <div class="form-outline mb-4">
                    <input type="text" class="form-control" v-model="titlee"/>
                </div>

                <div class="i-title">Текст</div>
                <div class="form-outline mb-4">
                    <textarea class="form-control" style="min-height: 150px" v-model="bodyy"></textarea>
                </div>

                <div class="i-title">Выбор тега</div>
                <button
                    class="btn c-q dropdown-toggle"
                    type="button"
                    id="dropdownMenuButton"
                    data-mdb-toggle="dropdown"
                    aria-expanded="false"
                >
                {{tagg}}
                </button>
                <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton">
                    <li><button class="dropdown-item" v-on:click="this.tagg = 'Социальная жизнь'">Социальная жизнь</button></li>
                    <li><button class="dropdown-item" v-on:click="this.tagg = 'Учебные новости'">Учебные новости</button></li>
                    <li><button class="dropdown-item" v-on:click="this.tagg = 'Жизнь ВУЗа'">Жизнь ВУЗа</button></li>
                </ul>
                
                <button type="submit" class="btn c-q btn-block mt-4" v-on:click="create()">Создать</button>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    data: function() {
        return {
            titlee: '',
            bodyy: '',
            tagg: 'Социальная жизнь'
        }
    },
    methods: {
        create() {
            let res = {
                title: this.titlee,
                body: this.bodyy,
                tag: this.tagg
            }

            axios.post("https://universityweb.site/api/addnews", res, {headers: {'Authorization': localStorage.getItem('accessToken')}}).then(respp => {
                if(respp.status == 200){
                    this.$notify({
                        title: 'Успех',
                        type: 'success',
                        text: 'Новость успешно добавлена'
                    })
                    this.$router.push('/')
                }
            }).catch(err => {
                this.$notify({
                    title: 'Ошибка',
                    type: 'error',
                    text: "Какая-то ошибка..."
                })
            })
        }
    }
}
</script>