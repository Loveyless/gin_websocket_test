<template>

  <div>
    <a href="https://vitejs.dev" target="_blank">
      <!-- /vite.svg 会默认去public下找 public下的东西好像不会被打包 -->
      <img src="/vite.svg" class="logo" alt="Vite logo" />
    </a>
    <a href="https://vuejs.org/" target="_blank">
      <!-- img的src自动可以用自己配的alias https://github.com/vitejs/vite/tree/main/packages/plugin-vue -->
      <img src="@/assets/logo/vue.svg" class="logo vue" alt="Vue logo" />
    </a>
  </div>

  <div>
    
  </div>

</template>

<script lang='ts' setup name="home">
</script> 

<style lang='less' scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
}

.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}

.logo.vue:hover {
  filter: drop-shadow(0 0 2em #008d47aa);
}

.hello_word_box {
  background-color: red;
}
</style>