d
<template>
<div class="wrapper">
    <!-- Navbar tsmc_navbar -->    
    <TsmcNavbar></TsmcNavbar>

    <!-- Main Sidebar Container -->
    <aside class="main-sidebar sidebar-dark-primary elevation-4" style="margin-top: 50px">
        <!-- tsmc_sidebar_timer -->
		<TsmcSidebarTimer></TsmcSidebarTimer>

        <!-- Brand Logo tsmc_sidebar_brand_logo-->
		<TsmcSidebarBrandLogo></TsmcSidebarBrandLogo>
       
        <!-- Sidebar -->
		<div class="sidebar">
			<!-- Sidebar user panel tsmc_sidebar_user (optional) -->
			<TsmcSidebarUser></TsmcSidebarUser>

			<!-- SidebarSearch Form  tsmc_sidebar_search-->
			<TsmcSidebarSearch></TsmcSidebarSearch>
            
             <!-- Sidebar Menu -->
             <nav class="mt-2">
                <ul	class="nav nav-pills nav-sidebar flex-column"
                    data-widget="treeview"
                    role="menu"
                    data-accordion="false">  


                    <!-- coding here-->
                    <!-- system menu tsmc_sidebar_menu_sys-->
                    <TsmcSidebarMenuSys></TsmcSidebarMenuSys>


                </ul>        
            </nav>            
        </div>
        <!-- /.sidebar -->
    </aside>



    <!-- Content Wrapper. Contains page content -->
    <div class="content-wrapper">
		<!-- Content Header (Page header) -->
		<!-- /.content-header -->

		<!-- Main content -->
		<KeepAlive>
			<RouterView></RouterView>
		</KeepAlive>

		<!-- /.content -->
    </div>
    <!-- /.content-wrapper -->
    
    <!-- Control Sidebar -->
    <aside class="control-sidebar control-sidebar-dark">
      <!-- Control sidebar content goes here -->
      	<div class="p-3">
			<h5>Title</h5>
			<p>Sidebar content</p>
		</div>
    </aside>
    <!-- /.control-sidebar -->

    <TsmcFooter></TsmcFooter>
</div>
</template>


<script>
import { RouterLink, RouterView ,KeepAlive} from 'vue-router'

//import { RouterLink, RouterView } from 'vue-router'
//import HelloWorld from './components/HelloWorld.vue'
import TsmcNavbar from "./components/os/TsmcNavbar.vue";
import TsmcSidebarTimer from "./components/os/TsmcSidebarTimer.vue";
import TsmcSidebarBrandLogo from "./components/os/TsmcSidebarBrandLogo.vue";
import TsmcSidebarUser from "./components/os/TsmcSidebarUser.vue";
import TsmcSidebarSearch from "./components/os/TsmcSidebarSearch.vue";
import TsmcSidebarMenuSys from "./components/sys/TsmcSidebarMenu.vue";

import TsmcFooter from "./components/os/TsmcFooter.vue";

export default {
    name: "App",
    components: {        
        TsmcNavbar,     
        TsmcSidebarTimer,   
        TsmcSidebarBrandLogo,
        TsmcSidebarUser,
        TsmcSidebarSearch,
        TsmcSidebarMenuSys,
        TsmcFooter
    }
}
</script>


<style></style>
