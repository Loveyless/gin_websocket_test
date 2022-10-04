<template>

  <div class="flex">

    <div class="w-500px">
      <h1>登录</h1>
      <el-form :model="loginUserInfo" label-width="80px" :inline="false">
        <el-form-item label="账号">
          <el-input v-model="loginUserInfo.username"></el-input>
        </el-form-item>
        <el-form-item label="密码">
          <el-input v-model="loginUserInfo.password"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="login">登录</el-button>
          <el-button type="warning" @click="quit">退出登录</el-button>
        </el-form-item>
      </el-form>
    </div>

    <div class="w-500px">
      <h1>注册</h1>
      <el-form :model="registerUserInfo" label-width="80px" :inline="false">
        <el-form-item label="账号">
          <el-input v-model="registerUserInfo.username"></el-input>
        </el-form-item>
        <el-form-item label="密码">
          <el-input v-model="registerUserInfo.password"></el-input>
        </el-form-item>
        <el-form-item label="邮箱">
          <el-input v-model="registerUserInfo.email">
            <template #suffix>
              <el-button v-if="globalStore.emailLow == 0" type="primary" plain @click="sendCode">发送验证码</el-button>
              <el-button v-else disabled type="primary" plain>{{ globalStore.emailLow }}秒后重试
              </el-button>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item label="验证码">
          <el-input v-model="registerUserInfo.code"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="register">注册</el-button>
          <!-- <el-button>取消</el-button> -->
        </el-form-item>
      </el-form>
    </div>

  </div>

  <Chat></Chat>

</template>

<script lang='ts' setup>
// import Chat from "./Chat.vue"

//pinia
const globalStore = GlobalStore()

//登录
const loginUserInfo = ref({
  username: "",
  password: "",
})

async function login() {
  const { data } = await http.post('/login', loginUserInfo.value)
  if (data.status === 200) {
    globalStore.setToken(data.data.token)
  }
}

async function quit() {
  globalStore.setToken("")
}
console.log(globalStore.token)

//注册
const registerUserInfo = ref({
  username: "",
  password: "",
  email: "",
  code: ""
})

// 刷新网页时 如果倒计时不是0一样执行
if (globalStore.emailLow != 0) {
  globalStore.setEmailLow()
}
async function sendCode() {
  const { data } = await http.post('/send/code', {
    email: registerUserInfo.value.email
  })
  //设置倒计时并且每秒减少1
  globalStore.emailLow = 30
  globalStore.setEmailLow()
}

async function register() {
  const { data } = await http.post('/register', registerUserInfo.value)
}

</script> 

<style lang='less' scoped>
@import "./index.less";
</style>