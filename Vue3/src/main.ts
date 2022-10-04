import { createApp } from "vue";
import "./assets/globalStyle.less";
import App from "./App.vue";

import router from "@/router";

import pinia from "@/store"

//vconsole
import Vconsole from "vconsole";

//unocss
import "uno.css";

// element css
import "element-plus/dist/index.css";
//element-plus-icon 如果您正在使用CDN引入，请删除下面一行。
import * as ElementPlusIconsVue from "@element-plus/icons-vue";


const app = createApp(App);

//element-plus-icon
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component);
}

//development
//production
if (import.meta.env.VITE_VCONSOLE == true) {
  let vConsole: any = new Vconsole();
  app.use(vConsole);
}

app.use(pinia).use(router).mount("#app");
