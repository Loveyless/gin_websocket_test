<template>
  <!-- 这里可以不用props.src -->
  <img v-bind="$attrs" :src="src" @error="onErr" :alt="alt">
</template>

<script lang='ts' setup>
//利用图片的error事件替换图片
const props = defineProps({
  src: {
    type: String,
    required: true
  },
  errSrc: {
    type: String,
    required: true
  },
  alt:{
    type: String,
    required: false,
    default(){
      return "img"
    }
  }
})

const onErr = (e: any) => {
  e.target.src = props.errSrc
}


//其实alt没什么用 因为errSrc如果是错的会递归执行onErr
//所以会无限出错 所以要保证errSrc必须有
//不过可以找找有没有解决放法

// 使用时
// <MyImg class="xxx" src="@/assets/imgs/pay.png" errSrc="@/assets/imgs/title.jpg"></MyImg>
</script> 
