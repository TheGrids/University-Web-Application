<template>
    <div class="card bg-dark text-white mt-4">
        <img src="../assets/5.png" class="card-img aa" alt="Stony Beach"/>
        <div class="card-img-overlay text-center align-middle">
            <span class="card-titlee"><b>Связь с ректором</b></span>
            <p class="card-text">
                <button type="button" class="btn btn-primary mt-3 " data-mdb-toggle="modal" data-mdb-target="#exampleModal">Написать</button>
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
                        <input type="text" id="form2Example1" class="form-control"/>
                    </div>

                    <div class="i-title">Сообщение</div>
                    <div class="form-outline mb-4">
                        <textarea class="form-control" style="min-height: 100px"></textarea>
                    </div>

                </div>

            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-secondary" data-mdb-dismiss="modal">Отмена</button>
                <button type="submit" class="btn btn-primary" >Отправить</button>
            </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Rectors'
}
</script>