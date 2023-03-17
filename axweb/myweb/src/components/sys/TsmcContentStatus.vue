<template>
<div>
    <div class="content-header">
        <div class="container-fluid">
        <div class="row mb-2">
            <div class="col-sm-6">
            <h1 class="m-0"> {{ setting_title }} </h1>
            </div><!-- /.col -->
            <div class="col-sm-6">
            <ol class="breadcrumb float-sm-right">
                <li class="breadcrumb-item"><a href="#">Home</a></li>
                <li class="breadcrumb-item active">System</li>                
            </ol>
            </div><!-- /.col -->
        </div><!-- /.row -->
        </div><!-- /.container-fluid -->
    </div>


    <div class="content">
        <div class="container-fluid">

            <div class="row">
                <div class="col-12">
                    <div id="toolbar">
                        <div class="form-inline" role="form">
                            <div class="form-group">                                                                
                            </div>         

                            <div class="form-group">                    
                            </div>

                            <div class="form-group">                            
                            </div>                  

                            <div class="form-group">
                            </div>                  

                        </div>
                    </div>  

                    <table
                        id="sensorStatus"
                        data-toggle="table"                        
                        data-toolbar="#toolbar"
                        data-height="760"
                        data-search="true"
                        data-click-to-select="true"
                        data-pagination="true"> 
                    </table>    
                                        
                </div>
            </div>    
            

            <div class="row">
                <div class="col-12">
                    <div class="form-inline" role="form">                        
                    </div>
                </div>
            </div>

      </div>
    </div>
</div>
</template>

<script>
export default {
    setup() {        
    },
    name:"TsmcContentStatus",
        data() {
        return {
            setting_title: "Dashboard OnLine",        
            stopinerval:null
        }
    },
    methods: {
        getSensorStatus() {
            let appurl = tsmcglb.apiurl + "/settingstatus"                       
            $.ajax({
                type: "GET",
                url: appurl,
                async: true,
                contentType: 'application/json; charset=utf-8',                
                success: function (jobj) {
                    console.log("get Status OK")
                    console.log( jobj )                    
                    $('#sensorStatus').bootstrapTable('load', jobj.data )
                },
                error: function (xhr) {
                    console.log("err")
                }
            });
        },
        restLocation( serialnumber , location){
            let self = this
            let appurl = tsmcglb.apiurl +  "/settingsconf"
            let jobj = { "serialnumber": serialnumber, "location":location }

            $.ajax({
                type: "POST",
                url: appurl,
                async: true,
                contentType: 'application/json; charset=utf-8',
                data: JSON.stringify(jobj),
                success: function (jobj) {                    
                    //console.log( jobj )
                    if ( jobj["success"]){
                        self.getSensorStatus()
                    }
                    
                },
                error: function (xhr) {
                    toastr.info(`search err...`);
                }
            });
        },
        operateEvents() {
            let self = this;
            return {
                'click .editOperate': function (e, value, row, index) {
                    console.log("edit")
                    //console.log(  row )
                    let serialnumber = row["serialnumber"]
                    let location = row["location"]
                    let locationval = prompt("Edit Location", location )

                    self.restLocation( serialnumber, locationval.trim() )
                },
                'click .rollcallOperate': function (e, value, row, index) {
                    //console.log("rollcallOperate click")
                }
            }
        },
     
    },
    mounted() {
        //boostrap table初使化
        let self = this;        
        $('#sensorStatus').bootstrapTable({
            locale: tsmcglb.tblang,
            pageSize: 10,
            exportDataType: 'all',
            //exportTypes:[ 'csv', 'txt','json', 'sql', 'doc', 'excel', 'xlsx', 'pdf'],
            exportTypes: ['csv', 'txt'],
            showExport: true,
            columns: [
                {
                    field: 'serialnumber',
                    title: this.$t("boostable.serialnumber"),
                },  
                {
                    field: 'modelname',
                    title: this.$t("boostable.modelname"),
                },                        
                {
                    field: 'status',
                    title: this.$t("boostable.status"),
                    formatter: function(value, row, index){                       
                        if (value == 1){
                            return "On"
                        }else if (value == 0){
                            return "Off"
                        }
                    },
                },
                {
                    field: 'location',
                    title: this.$t("boostable.location"),
                },
                {
                    field: "updatedatetime",
                    title: this.$t("boostable.updatedatetime"),
                    formatter: function(value, row, index){                       
                        let v = parseInt(value)
                        return tsmcglb.dtfmt( v )
                    },
                },
                {
                    field: 'operate',
                    title: this.$t("boostable.operate"),
                    formatter: `<button class="btn bg-gradient-info editOperate" type="button"> ${this.$t('button.edit')}</>`,
                    events: self.operateEvents()
                }
            ],
            onClickRow: function (row, $element, field) {
            }
        })

        //test event
    },
    created() {
        //order channel sys
        console.log("tsmc_content_sys_status created")
        var self = this;
    },
    activated() {
        console.log("sys_status activeted")      
        this.getSensorStatus()
        this.stopinerval = setInterval(() => {
            this.getSensorStatus()
        }, 5000);
    },
    deactivated() {
        console.log("sys_status disactiveted")              
        clearInterval( this.stopinerval )
    },
}
</script>


<style>
</style>