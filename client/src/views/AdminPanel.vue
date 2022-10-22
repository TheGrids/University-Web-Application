<template>
    <div class="container">
        <h2 class="mt-4">Админ панель</h2>
        <div class="table-responsive">

            <table class="table table-striped align-middle mb-0 bg-white mt-2">
            <thead class="bg-white">
                <tr>
                <th>Пользователь</th>
                <th>Роль</th>
                <th>Действия</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                <td>
                    <div class="d-flex align-items-center">
                    <img
                        src="https://mdbootstrap.com/img/new/avatars/8.jpg"
                        alt=""
                        style="width: 45px; height: 45px"
                        class="rounded-circle"
                        />
                    <div class="ms-3">
                        <p class="fw-bold mb-1">Имя Фамилия</p>
                        <p class="text-muted mb-0">email@gmail.com</p>
                    </div>
                    </div>
                </td>
                <td >
                    <button
                        class="btn btn-primary dropdown-toggle"
                        type="button"
                        id="dropdownMenuButton"
                        data-mdb-toggle="dropdown"
                        aria-expanded="false"
                    >
                        Студент
                    </button>
                    <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton">
                        <li><a class="dropdown-item" href="#">Студент</a></li>
                        <li><a class="dropdown-item" href="#">Преподаватель</a></li>
                        <li><a class="dropdown-item" href="#">Администратор</a></li>
                    </ul>
                </td>
                <td>
                    <button type="submit" class="btn c-q btn-floating" v-on:click=""><i class="fas fa-trash"></i></button>
                    <button type="submit" class="btn c-q btn-floating" v-on:click=""><i class="fas fa-trash"></i></button>
                </td>
                </tr>
            </tbody>
            </table>
        </div>
    </div>
</template>

<script>
export default {
    name: 'AdminPanel',
}
</script>