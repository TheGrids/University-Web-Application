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
                Все
            </button>
            <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton">
                <li><button type="submit" class="dropdown-item">Все</button></li>
                <li><button type="submit" class="dropdown-item">Социальная жизнь</button></li>
                <li><button type="submit" class="dropdown-item">Учебные новости</button></li>
                <li><button type="submit" class="dropdown-item">Жизнь ВУЗа</button></li>
            </ul>
        </div>
    </div>

    <div class="card mt-4">
        <h5 class="card-header d-flex justify-content-between">
            <button type="button" class="btn btn-outline-info btn-rounded" data-mdb-ripple-color="dark">
                <span class="sss">Социальная жизнь</span>
            </button> 
            <button type="submit" class="btn btn-danger btn-floating"  style=" margin-bottom: 5px">
                <i class="fas fa-trash"></i>
            </button>
        </h5>
        <div class="card-body">
            <h5 class="card-title">Животным денег не давать.</h5>
            <p class="card-text">О событиях на прошлый вечер: на площадке форума "Хакатон autumn 2022" найдены котики и они мурчат. Сначала один котик требовал деньги на прокорм, но после строгого выговора он перестал промышлять подобными делами. Затем котиков стало больше, они развлекали участников, разрешали делать фотографии с ними (в этот раз бесплатно). Затем один из плюшевых порадовал участников форума своей игрой на гитаре.</p>
        </div>
        <div class="card-footer text-muted">
            <i class="far fa-calendar-plus"></i>  23.10.2022 14:08:23 | Автор: Елизавета Пшеничная
        </div>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: 'NewsBlock',
    data() {
        return {

        }
    },
    mounted() {
        axios.get("https://universityweb.site/api/news", {headers: {'Authorization': localStorage.getItem('accessToken')}}).then(respp => {
            console.log(respp.data.data);
        }).catch(err => {
            console.log(err);
        })
    }
}
</script>