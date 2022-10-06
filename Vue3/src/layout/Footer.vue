<template>
  <van-tabbar v-model="active" @change="onChange" class="fixed">
    <van-tabbar-item name="home" replace to="/layout/home" icon="home-o">首页</van-tabbar-item>
    <van-tabbar-item name="chat" replace to="/layout/chat" icon="friends-o">聊天</van-tabbar-item>
    <van-tabbar-item name="setting" replace to="/layout/setting" icon="setting-o">设置</van-tabbar-item>
  </van-tabbar>
</template>

<script lang='ts' setup>

const active = ref<string>("home");
const onChange = (index: string) => {

}

</script> 

<style lang='less' scoped>
</style>