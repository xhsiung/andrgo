<template>
    <footer class="main-footer">
        <!-- To the right -->
        <div class="float-right d-none d-sm-inline">      
        </div>
        
        <!-- Default to the left -->
        <strong>Copyright &copy; {{time1}}-{{time2}} <a href="https://tsmccomp.com/">tsmccomp</a>.</strong> All rights reserved. ver:1.0
    </footer>
</template>

<script>
export default {
    setup() {        
    },
    name: "TsmcFooter",
    data(){
        return{
            time1:"",
            time2:""
        }
    },
    methods: {

    },
    mounted(){
        this.time1 = new Date().getFullYear();
        this.time2 = this.time1 + 5
    },
    created(){
    }
}
</script>

<style scoped>
</style>