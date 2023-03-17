<template>
<li class="nav-item menu-open">
    <a href="#" class="nav-link active">
        <i class="nav-icon fas fa-tachometer-alt"></i>
        <p>
            {{ $t('left.sys_tool') }}
            <i class="right fas fa-angle-left" style="margin:-8px -120px;"> </i>
        </p>
    </a>
    
    <ul class="nav nav-treeview ">        
        <RouterLink to="/tsmccontentsysstatus">
        <li class="nav-item">
            <a href="#" class="nav-link" >
                <i class="far fa-circle nav-icon"></i>
                <p>
                    {{ $t('left.sys_status') }}
                </p>
            </a>
        </li>
        </RouterLink>
    </ul>

</li>
</template>

<script>
import { RouterLink } from 'vue-router'
export default {
    setup() {        
    },
    name:"TsmcSidebarMenu",
    data() {
        return {
            modalShow: true,
            form: {
                account: '',
            },
        }
    },
    methods: {
        scanip() {    
        },
    },
    created() {
    }  
}
</script>

