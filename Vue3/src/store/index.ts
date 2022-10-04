import { defineStore, createPinia } from "pinia";
import piniaPluginPersistedstate from "pinia-plugin-persistedstate";

// 使用setup模式定义
export const GlobalStore = defineStore(
  "GlobalStore",
  () => {
    //websocket里面要用到

    //保存token websocket里面要用到
    let token = ref<string>("");
    function setToken(tokenStr: string) {
      token.value = tokenStr;
    }
    //保存用户信息
    let userInfo = ref<string>("");
    //注册邮箱倒计时
    let emailLow = ref<number>(0);
    function setEmailLow() {
      const timer = setInterval(() => {
        if (emailLow.value == 0) {
          clearInterval(timer);
          return;
        }
        emailLow.value -= 1;
      }, 1000);
    }

    let count = ref<number>(0);
    function countAdd(): void {
      count.value++;
    }
    function doubleCount(): number {
      return count.value * 2;
    }

    return { token, setToken, userInfo, emailLow, setEmailLow, count, countAdd, doubleCount };
  },
  {
    persist: true,
  }
);

const pinia = createPinia();
pinia.use(piniaPluginPersistedstate);
export default pinia;

// 使用options API模式定义
// export const Global = defineStore("Global", {
//   // 定义state
//   state: () => ({
//     count1: 1,
//   }),
//   // 定义action
//   actions: {
//     increment() {
//       this.count1++;
//     },
//   },
//   // 定义getters
//   getters: {
//     doubleCount(state) {
//       return state.count1 * 2;
//     },
//   },
// });
