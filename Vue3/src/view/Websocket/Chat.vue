<template>

  <div style="text-align: left;">

    <h2 v-if="globalStore.token == ''">{{ "登录后开启聊天" }}</h2>
    <el-button type="primary" @click="connectWebsocket">连接聊天</el-button>
    <el-button type="warning" @click="close">关闭聊天</el-button>

    <div class="w-500px h500px overflow-auto box_chat">
      <div class="box_chat_item" v-for="item in chatList">
        {{ item.username + " " + item.time + " " + item.message }}
      </div>
    </div>
    <el-input style="margin-left:10px;width:500px;" v-model="msgData.message">
      <template #suffix>
        <el-button type="primary" plain @click="sendMessage">发送</el-button>
      </template>
    </el-input>

  </div>


</template>

<script lang='ts' setup>
//pinia
const globalStore = GlobalStore()

var ws: WebSocket

interface chartListRule {
  username: string,
  message: string,
}
let chatList = ref<any>([])

//发送消息
let msgData = ref({
  message: "",
  room_identity: "980",
  token: globalStore.token,
})

//发送消息
function sendMessage() {
  ws.send(JSON.stringify(msgData.value));
  msgData.value.message = ""
}

//连接聊天
function connectWebsocket() {
  ws = useWebSocket(
    `ws://localhost:8080/websocket/message`,
    // 接收到消息
    (evt) => {
      console.log('收到来自服务端的消息：', evt.data);
      chatList.value.push(JSON.parse(evt.data))
      // chatList.value.push(evt.data)
    }
  )
};

//关闭
function close() {
  ws.close()
}



</script> 

<style lang='less' scoped>
@import "./index.less";
</style>