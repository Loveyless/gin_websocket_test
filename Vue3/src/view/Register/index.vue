<template>

  <br>
  <h1>聊天系统</h1>
  <h2>注册</h2>
  <br>

  <van-form @submit="onSubmit">

    <van-cell-group inset>
      <van-field v-model="registerInfo.username" name="username" label="用户名" placeholder="用户名 3-20位"
        :rules="[{ required: true, message: '请填写用户名' }]" clearable />
      <van-field v-model="registerInfo.password" type="password" name="password" label="密码" placeholder="密码 3-20位"
        :rules="[{ required: true, message: '请填写密码' }]" clearable />
      <van-field v-model="registerInfo.email" type="email" name="email" label="邮箱" placeholder="邮箱"
        :rules="[{ required: true, message: '请填写邮箱' }]" clearable />
      <van-field v-model="registerInfo.code" type="code" name="code" label="验证码" placeholder="验证码"
        :rules="[{ required: true, message: '请填写验证码' }]" clearable>
        <template #button>
          <van-button v-if="globalStore.emailLow == 0" size="small" type="primary" @click="sendCode">{{ "发送验证码" }}
          </van-button>
          <van-button v-if="globalStore.emailLow != 0" size="small" type="primary" disabled>{{ globalStore.emailLow }}
          </van-button>
        </template>
      </van-field>
    </van-cell-group>

    <div style="margin: 16px;">
      <van-button round block type="primary" native-type="submit">
        注册
      </van-button>
      <div class="mt-3"></div>
      <van-button round block type="default" @click="$router.push('/login')">
        登录
      </van-button>
    </div>

  </van-form>
</template>

<script lang='ts' setup>
const globalStore = GlobalStore()
const router = useRouter()

//验证码倒计时
if (globalStore.emailLow != 0) {
  globalStore.setEmailLow()
}

const registerInfo = ref({
  username: '',
  password: '',
  email: '',
  code: ''
})
const sendCode = async () => {
  const { data } = await http.post("/send/code", { email: registerInfo.value.email })
  if (data.status == 200) {
    globalStore.emailLow = 30
    globalStore.setEmailLow()
  }
}
const onSubmit = async (values) => {
  const { data } = await http.post("/register", registerInfo.value)
  if (data.status == 200) {
    router.push("/login")
  }
};

</script> 

<style lang='less' scoped>
</style>