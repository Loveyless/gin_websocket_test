<template>
  <div class="container flex flex-row-reverse">

    <div class="chat w-100% h-80%">
      <div v-for="item in messageList">{{ item.username + " " + item.time + " " + item.message }}</div>
    </div>


    <div class="w-100% h-20%">
      <van-field v-model="message" center placeholder="发送消息">
        <template #button>
          <van-button size="small" type="primary" @click="sendMessage">发送</van-button>
        </template>
      </van-field>
    </div>
  </div>
</template>

<script lang='ts' setup>

const globalStore = GlobalStore()


const ws = useWebSocket(
  `ws://${import.meta.env.VITE_BASE_WSURL as string}/websocket/all?token=${globalStore.token}`,
  (data) => {
    console.log("接收的值", JSON.parse(data.data))
    messageList.value.unshift(JSON.parse(data.data))
  }
)


const message = ref('')
const messageList = ref<any>([])

// 发送消息
const sendMessage = async () => {
  if (message.value == '') {
    return
  }
  ws.send(JSON.stringify({ message: message.value }))
  message.value = ''
}


onUnmounted(() => {
  ws.close()
})

</script> 

<style lang='less' scoped>
@import 'index.less';
</style>