<template>
    <div class="user-panel mt-3 pb-3 mb-3 d-flex" >            
            <div class="image">
                <img src="/static/images/logo.png" class="img-circle elevation-2" alt="User Image">
            </div>
            <div class="info" style="margin-top: 1px;">
                <a href="#" class="d-block">{{ username }}</a>
            </div>                                
    </div>
</template>

<script>
export default {
    setup() {        
    },
    name: "TsmcSidebarUser",
    data() {
        return {                      
            username: "demo",                            
        }
    },
    methods: {
    },
    mounted(){
    },
    created() {
        //get user session
        let self = this     

        // var cancelnterval = setInterval( ()=>{
        //     if (window.tsmcglb.login_user ){
        //         self.username = window.tsmcglb.login_user["username"]
        //         clearInterval(cancelnterval);
        //         console.log("end")
        //     }            
        // }, 600) 
    }
}
</script>