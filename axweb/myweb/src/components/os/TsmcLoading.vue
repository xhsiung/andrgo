<template>
    <div id="loadingmodal" class="modal fade" tabindex="-1" role="dialog" data-backdrop="static" data-keyboard="false">
        <div class="modal-dialog modal-dialog-centered justify-content-center" role="document">
            <span class="fa fa-spinner fa-spin fa-8x"></span>
        </div>
    </div>
</template>

<script>
export default {
    setup(){
    },
    name:"TsmcLoading",
    data(){
        return {
            msg:"..."
        }
    },
    methods:{
        show(tt=5000){
            $('#loadingmodal').modal('show');
            setTimeout(function () {            
                $('#loadingmodal').modal('hide');
            }, tt);
        },
        hide(tt=0){            
            setTimeout(function () {            
                $('#loadingmodal').modal('hide');
            }, tt);
        }
    },
    created(){
        var self = this;
               
    }
}
</script>