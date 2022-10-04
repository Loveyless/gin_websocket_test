<template>

  <br>
  <h1>聊天系统</h1>
  <h2>登录</h2>
  <br>

  <van-form @submit="onSubmit">

    <van-cell-group inset>
      <van-field v-model="loginInfo.username" name="username" label="用户名" placeholder="用户名 3-20位"
        :rules="[{ required: true, message: '请填写用户名' }]" clearable />
      <van-field v-model="loginInfo.password" type="password" name="password" label="密码" placeholder="密码 3-20位"
        :rules="[{ required: true, message: '请填写密码' }]" clearable />
    </van-cell-group>

    <div style="margin: 16px;">
      <van-button round block type="primary" native-type="submit">
        登录
      </van-button>
      <div class="mt-3"></div>
      <van-button round block type="default" @click="$router.push('/register')">
        注册
      </van-button>
    </div>

  </van-form>
</template>

<script lang='ts' setup>
const globalStore = GlobalStore()
const router = useRouter()

const loginInfo = ref({
  username: '',
  password: ''
})
const onSubmit = async (values) => {
  const { data } = await http.post("/login", loginInfo.value)
  if (data.status == 200) {
    globalStore.setToken(data.data.token)
    router.push("/layout")
  }
};

</script> 

<style lang='less' scoped>
</style>