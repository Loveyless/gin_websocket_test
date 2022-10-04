<template>
  <div class="common-layout">
    <el-container>

      <el-main style=" margin:0 ;padding:0;">
        <router-view v-slot="{ Component, route }">
          <Transition appear name="fade-transform" mode="out-in">
            <keep-alive :exclude="['/xxx']">
              <div :key="$route.path">
                <component :is="Component" :key="route.path" />
              </div>
            </keep-alive>
          </Transition>
        </router-view>
      </el-main>


      <!-- <el-footer> -->
      <Header />
      <!-- </el-footer> -->

    </el-container>
  </div>
</template>

<script lang='ts' setup name="layout">
// import Header from "./Header.vue";

</script> 

<style lang='less' scoped>
/* fade-transform */
.fade-transform-leave-active,
.fade-transform-enter-active {
  transition: all 0.2s;
}

.fade-transform-enter-from {
  opacity: 0;
  transform: translateX(-30px);
  transition: all 0.2s;
}

.fade-transform-leave-to {
  opacity: 0;
  transform: translateX(30px);
  transition: all 0.2s;
}
</style>