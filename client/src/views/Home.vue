<template>
    <div class="container">
        <Rectors />
		
        <NewsBlock />

    </div>
</template>

<style>
.btn-rounded {
    padding: 5px 10px 5px 10px;
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
import NewsBlock from '../components/NewsBlock.vue'
import Rectors from '../components/Rectors.vue'

export default {
	name: 'Home',
	components: {
		NewsBlock,
		Rectors
	}
}
</script>
