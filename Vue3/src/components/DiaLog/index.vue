<template>

  <teleport to="body">
    <div class="content" v-show="isShow">
      <div class="box">
        <div class="close" @click="emit('close')">
          <svg width="1em" height="1em" viewBox="0 0 24 24">
            <path fill="currentColor"
              d="m12 10.586l4.95-4.95l1.414 1.414l-4.95 4.95l4.95 4.95l-1.414 1.414l-4.95-4.95l-4.95 4.95l-1.414-1.414l4.95-4.95l-4.95-4.95L7.05 5.636z">
            </path>
          </svg>
        </div>
        <div class="title">
          <slot name="title">标题</slot>
        </div>
        <div class="main">
          <slot>内容</slot>
        </div>
        <div class="bottom">
          <slot name="bottom">底部按钮</slot>
        </div>
      </div>
    </div>
  </teleport>

</template>

<script lang='ts' setup>
//三个插槽 title 默认 bottom

const props = withDefaults(defineProps<{
  //显示与否
  isShow?: boolean,
  //滚动条
  isScroll?: boolean
}>(), {
  isShow: false,
  isScroll: true
})

//右上角icon事件
const emit = defineEmits(["close"])

watch(() => props.isShow, () => {
  //滚动条
  if (props.isScroll && props.isShow) {
    document.body.style.overflowY = "hidden";
  } else {
    document.body.style.overflowY = "auto";
  }
})


</script> 

<style lang='less' scoped>
@import "./index.less";
</style>