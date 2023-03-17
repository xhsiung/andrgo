<template>
    <div>        
        <div class="container">
            <div class="row" style="height:220px;width: 820px;">
                <div class="col-12">                    
                    <p id="P61Gauge" style="font-size:20px;text-align:center;"> </p>
                </div>    
            </div>

            <div class="row">
                <div class="col-1">            
                </div>
                <div class="col-9">
                    <!-- <LABEL style="font-size:78px;text-align:center;">Taiwan Silicon Microelectroics Corp.</LABEL> -->
                    
            
                </div>
                <div class="col-2">            
                </div>
            </div>
    
        </div>
    </div>

</template>

<script>
export default {
    setup() {
        
    },
    name: "TsmcHome",
    mounted(){  
        window.emitter.on("mqtt", jobj => {    
            let linemsg = `id:${jobj.id},model:${jobj.modelname}, serial:${jobj.serialnumber},Press:${jobj.object[0].value[0]},Temp:${jobj.object[1].value[0]} <BR>`
            $( "#P61Gauge" ).prepend( linemsg );
            //console.log( jobj )
        })
    },
    created() {
      
    }
}
</script>


<style scoped>
.position-absolute {
   position: absolute !important;
}
</style>