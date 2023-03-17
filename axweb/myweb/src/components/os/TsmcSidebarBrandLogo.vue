<template>    
    <a href="/index" class="brand-link" style="margin-top:60px">
        <img src="/static/dist/img/AdminLTELogo.png" alt="TSMCCOMP Logo"
            class="brand-image img-circle elevation-3" style="opacity: .8">
        <span class="brand-text font-weight-light"> {{ $t("logotext") }} </span>
    </a>    
</template>

<script>
export default {
    setup() {             
    },
    name: "TsmcSidebarBrandLogo", 
    data() {
        return {                     
        }
    },
    methods: {

    },
    created() {
    }
}
</script>

<style scoped>
</style>