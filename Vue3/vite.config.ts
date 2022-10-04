import { defineConfig } from "vite";

import { fileURLToPath, URL } from "url";

//如果这样引入 需要在tsconfig.node.json 里配置"include": ["config/*.ts"]
import build from "./config/build";
import plugins from "./config/plugins";

// https://vitejs.dev/config/
export default defineConfig({
  base: "/",
  publicDir: "public",
  // 设为 false 可以避免 Vite 清屏而错过在终端中打印某些关键信息
  clearScreen: false,
  // 以 envPrefix 开头的环境变量会通过 import.meta.env 暴露在你的客户端源码中。 envPrefix 不应被设置为空字符串 ''
  envPrefix: "VITE_",
  resolve: {
    alias: {
      //    转为字符串      生成新的连接    后续连接         当前连接
      //https://developer.mozilla.org/zh-CN/docs/Web/API/URL/URL
      "@": fileURLToPath(new URL("./src", import.meta.url)),
      "@v": fileURLToPath(new URL("./src/view", import.meta.url)),
      "@c": fileURLToPath(new URL("./src/components", import.meta.url)),
    },

    //另一种写法
    // alias: [
    //   { find: "@/", replacement: fileURLToPath(new URL("./src", import.meta.url)) },
    //   { find: "@c/", replacement: fileURLToPath(new URL("./src/components", import.meta.url)) }
    // ]
  },
  build,
  plugins,
  server: {},
});
