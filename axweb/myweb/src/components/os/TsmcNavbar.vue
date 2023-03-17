<template>
    <nav class="main-header navbar navbar-expand navbar-dark">
        <!-- Left navbar links -->
        <ul class="navbar-nav">
            <li class="nav-item">
                <a class="nav-link" @click="resize()" data-widget="pushmenu" href="#" role="button"><i
                        class="fas fa-bars"></i></a>
            </li>
            <li class="nav-item d-none d-sm-inline-block">
                <a href="#" class="nav-link" style="width: 64px">
                    {{ $t("top.home") }}</a>
            </li>
            <li class="nav-item d-none d-sm-inline-block">
                <a href="#" class="nav-link" style="width: 96px">
                    {{ $t("top.contact") }}
                </a>
            </li>
        </ul>


        <!-- &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        <marquee direction="left" scrollamount="6">
            <div style="color: #ff6443; font-size: 26px">{{ msg }}</div>
        </marquee> -->

        <!-- Right navbar links -->
        <ul class="navbar-nav ml-auto">
            <!-- Starter-->
            <li class="nav-item">
                <a class="nav-link" data-widget="fullscreen" href="#" role="button">
                    <i class="fas fa-expand-arrows-alt"></i>                    
                </a>
            </li>

            <li class="nav-item">
                <a class="nav-link" data-widget="control-sidebar" data-slide="true" href="#" role="button">
                    <i class="fas fa-th-large"></i>                    
                </a>
            </li>

            <li class="nav-item">
                <a class="nav-link" data-widget="" data-slide="" href="#" @click="logout" role="button">                                                                            
                    <i class="fas fa-sign-out-alt"></i>
                </a>
            </li>
        </ul>
    </nav>
</template>

<script>
export default {
    setup() { },
    name: "TsmcBavbar",
    data() {
        return {
            msg: "",
            eventhings: [],
            cleartime: 300
        }
    },
    methods: {
        resize() {
            //alert("test")
            setTimeout(() => {
                window.dispatchEvent(new Event('resize'));
            }, 600)
        },
        logout() {
            let self = this
            console.log("logout")
            sessionStorage.removeItem("tsmc_info")
            location.href = "/login"
        },
        filterEventThings(data) {
            let self = this
            if (data) {
                let action = data["action"].toLowerCase()
                let key = ""
                if (action.indexOf("press") != -1) {
                    key = "pressurevalue"
                }

                if (action.indexOf("temp") != -1) {
                    key = "temperaturevalue"
                }

                if (action.indexOf("humidity") != -1) {
                    key = "humidityvalue"
                }

                let obj = {
                    "evaction": data["action"],
                    "evtype": data["evtype"],
                    "serialnumber": data["data"]["serialnumber"],
                    "modelname": data["data"]["modelname"],
                    "val": data["data"][key][0],
                    "unit": data["data"][key + 'unit'],
                    "createdatetime": data["data"]["createdatetime"],
                }

                let index = self.eventhings.findIndex(o => {
                    if (o["serialnumber"] == obj["serialnumber"]) {
                        if (o["evaction"] == obj["evaction"]) {
                            return true
                        } else {
                            return false
                        }
                        //return true
                    }
                })

                if (index == -1) {
                    self.eventhings.push(obj)
                } else {
                    self.eventhings.splice(index, 1, obj)
                }

            }
        },
        filterMsg() {
            let self = this
            let msgarr = []
            for (let i = 0; i < self.eventhings.length; i++) {
                let elem = self.eventhings[i]
                let hsm = tsmcglb.dtfmt(elem["createdatetime"], "hms")
                let msg = `SN:${elem["serialnumber"]}(${elem["evaction"]} Val:${elem["val"]}${elem["unit"]} Time:${hsm})`
                msgarr.push(msg)
            }
            self.msg = msgarr.join("，")
        },
        clearMsg() {
            let self = this
            let clearms = (self.cleartime) * 1000
            let index = self.eventhings.findIndex(o => {
                let diff = Date.now() - o["createdatetime"]
                if (diff > clearms) {
                    return true
                }
            })

            if (index >= 0) {
                self.eventhings.splice(index, 1)
            }
        }
    },
    created() {
        let self = this
        console.log("Created")

        // tsmcev.listen("tsmc_navbar", function (data) {
        //     //console.log("tsmc_navbar", data )            
        //     self.filterEventThings(data)
        // })  

        setInterval(() => {
            self.filterMsg()
        }, 3000);

        setInterval(() => {
            self.clearMsg()
        }, 3000)

    }
};
</script>

<style></style>