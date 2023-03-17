<template>    
    <a href="#" class="brand-link">
         <span class="brand-text font-weight-light" style="margin-left:14px; color:white ; font-size:22px"> {{ timer }}</span>
    </a>   
</template>

<script>
export default {
    setup() {        
    },
    name: "TsmcSidebarTimer",
    data() {
        return {
            timer:"0000-00-00 00:00:00"
        }
    },
    methods: {

    },
    mounted(){
        let self = this
        setInterval(()=>{
            let now = new Date();
            let mm = (now.getMonth()+1).toString()
            let dd = now.getDate().toString()
            let h = now.getHours().toString()
            let m = now.getMinutes().toString()
            let s = now.getSeconds().toString()
            let o = {
                "yy": now.getFullYear(),
                "mm": (mm.length == 1)? ('0'+mm) : (''+mm),
                "dd": (dd.length == 1)? ('0'+dd) : (''+dd),
                "h":  (h.length == 1) ? ('0'+h) : (''+h),
                "m":  (m.length == 1) ? ('0'+m) : (''+m),
                "s":  (s.length == 1) ? ('0'+s) : (''+s)
            }
            self.timer = `${o['yy']}-${o['mm']}-${o['dd']} ${o['h']}:${o['m']}:${o['s']}`
        }, 1000)
    },
    created() {
    }

}
</script>

