
## 目录

## 记录

接下里看自定义指令


## 集成

### axios
请求取消 

### router
路由表meta进行判断
切换时loading 和 取消请求

## 探索插件

**集成**

@vitejs/plugin-legacy - 旧版浏览器支持。

vite-plugin-ali-oss - 将项目捆绑的生产文件上传到阿里OSS，HTML除外。 / vite-plugin-tencent-oss - 将项目捆绑的生产文件上传到腾讯OSS，HTML除外。

vite-plugin-webfont-dl - 下载并注入 webfonts（谷歌字体）以提高网站的性能。



**装载？**

vite-plugin-rsw - 加载 rust-compiled (wasm-pack) WebAssembly 包。



**打包**

vite-plugin-imagemin - 压缩图像资产。

vite-plugin-compress - 压缩你的包 + 资产。/ vite-plugin-compression - 使用 gzip 或 brotli 压缩资源。/ vite-compression-

plugin - 使用 Node.js 流压缩文件到 gzip 或更多。

vite-plugin-banner - 在每个生成的块的顶部添加一个横幅。

>vite-plugin-vue-docs - 分析 vue 组件以自动生成可预览文档。
vite-plugin-importus - 模块化导入插件，兼容 antd、lodash、material-ui 等。
>vite-plugin-cdn-import CDN插件



**变行金刚？**

vite-plugin-unocss-to-uni - UnoCSSin uni-app。
vite-plugin-mock mock数据 提供本地和生产模拟服务。  https://github.com/vbenjs/vite-plugin-mock/blob/main/README.zh_CN.md



**好帮手**

>vite-plugin-theme - 动态改变主题颜色。
vite-plugin-mkcert - 提供证书 https 开发服务器。
vite-dts.d.ts -为库生成模块的超快插件。
vitawind - 自动安装和设置 Tailwind CSS。
@zougt/vite-plugin-theme-preprocessor - 轻松实现基于 LESS 或 SASS 的动态主题。
>vite-plugin-autogeneration-import-file - 自动生成导入文件。
>vite-plugin-md2vue - 将 markdown 模块转换为 Vue 组件。
>vite-plugin-svg-icons 用于生成 svg 雪碧图.  https://github.com/vbenjs/vite-plugin-svg-icons/blob/main/README.zh_CN.md



**其他**

>vite-plugin-md - Markdown 作为 Vue 组件 / Markdown 中的 Vue 组件。
vite-svg-loader - 将 SVG 文件加载为 Vue 组件。
v2 vite-plugin-vue2-svg - 将 SVG 文件加载为 Vue 组件。

SSG

vite-ssg - 服务器端生成。

vite-plugin-vue-i18n - Vue I18n 的集成。

vite-plugin-i18n-resources - 加载 i18n 翻译消息文件。



**用过的**

vite-plugin-remove-console - 在生产环境中删除 console.log 的 vite 插件。



## 使用的插件
```js
//vite/vue插件
import vue from "@vitejs/plugin-vue";

// 自动引入api插件 https://github.com/antfu/unplugin-auto-import
// 记得 ts.config.json文件引入声明文件: include中引入auto-imports.d.ts 要不然编辑器报错
import AutoImport from "unplugin-auto-import/vite";

// 自动引入组件插件 https://github.com/antfu/unplugin-vue-components
import Components from "unplugin-vue-components/vite";

// element-plus自动引入组件插件 element-plus
import { ElementPlusResolver } from "unplugin-vue-components/resolvers";

// 点哪里打开代码哪里 https://github.com/webfansplz/vite-plugin-vue-inspector 我的vscode没有shell 用不了操
import Inspector from "vite-plugin-vue-inspector";

// 图标库 https://github.com/antfu/unplugin-icons 图标地址 https://icon-sets.iconify.design https://github.com/iconify/iconify
import Icons from "unplugin-icons/vite";
// 图标库自动引入 这个结合 unplugin-vue-components 使用
import IconsResolver from "unplugin-icons/resolver";



// 原子化css https://github.com/unocss/unocss
import Unocss from "unocss/vite";
// 原子化css 第一个是工具类预设，第二个是属性化模式支持，第三个是icon支持 还有别的预设
import { presetUno, presetAttributify, presetIcons } from 'unocss'
import unocssRule from "./unocss";


// 使vue脚本设置语法支持name属性 https://github.com/chenxch/unplugin-vue-setup-extend-plus
import vueSetupExtend from 'unplugin-vue-setup-extend-plus/vite'


// 自动引入图片 https://github.com/sampullman/vite-plugin-vue-images
// 注意！如果变量（道具、数据）与图像名称冲突 它们将在模板中被破坏 在模板中使用图像而无需import通过data 直接写图像名称 目前不支持重复的图像名称。目前，或必须使用v-bind:src速记。:src
import ViteImages from "vite-plugin-vue-images";


// 打包时显示进度条 https://github.com/jeddygong/vite-plugin-progress
import progress from "vite-plugin-progress";

```

## unocss
可以安装一下unocss插件
```js
安装unocss和三个预设，第一个是工具类预设，第二个是属性化模式支持，第三个是icon支持 还有别的预设
pnpm i -D unocss @unocss/preset-uno @unocss/preset-attributify @unocss/preset-icons

配置
// vite.config.ts
import Unocss from 'unocss/vite'
import { presetUno, presetAttributify, presetIcons } from 'unocss'

export default {
  plugins: [
    Unocss({ // 使用Unocss
      presets: [
        presetUno(),
        presetAttributify(),
        presetIcons()],
    }),
  ],
}



注册
// main.ts
import 'uno.css'




图标功能 // https://github.com/unocss/unocss/tree/main/packages/preset-icons https://icones.js.org/

下载图标库
pnpm i -D @iconify-json/[the-collection-you-want]
或者一次性安装
npm i -D @iconify/json

然后引入在配置
  presets: [
    presetIcons({
      cdn: 'https://esm.sh/' //您可以指定cdn选项 since v0.32.10。我们推荐esm.sh作为 CDN 提供程序。
    })
  ],
然后直接写在类名里就好




增加预设css配置
Unocss({
  presets: [
    presetUno(),
    presetAttributify(),
    presetIcons()
  ],
  rules: [ // 在这个可以增加预设规则, 也可以使用正则表达式
    [
      'p-c', // 使用时只需要写 p-c 即可应用该组样式
      {
        position: 'absolute',
        top: '50%',
        left: '50%',
        transform: `translate(-50%, -50%)`
      }
    ],
    [/^m-(\d+)$/, ([, d]) => ({ margin: `${d / 4}rem` })],
  ]
})


```
unocss默认无任何预设
但是可以选择默认预设 
例如：ml-3(Tailwind)、ms-2(Bootstrap)、ma4(Tachyons) 和mt-10px(Windi CSS) 都是有效的。
```css
.ma4 { margin: 1rem; }
.ml-3 { margin-left: 0.75rem; }
.ms-2 { margin-inline-start: 0.5rem; }
.mt-10px { margin-top: 10px; }
```
设置默认预设来这里
https://github.com/unocss/unocss/tree/main/packages/preset-uno



**可以自定义预设 不用他的**
一个老哥的自定义预设
```js
export default {
  presets: [], // 不用预设自定义引擎
  rules: [
    [/^m-(-?\d+)(px|%|vw|vh|rem|em)?$/, (match) => ({ margin: `${match[1]}${match[2] || 'px'}` })],
    [
      /^mx-(-?\d+)(px|%|vw|vh|rem|em)?$/,
      (match) => ({
        'margin-right': `${match[1]}${match[2] || 'px'}`,
        'margin-left': `${match[1]}${match[2] || 'px'}`
      })
    ],
    [
      /^my-(-?\d+)(px|%|vw|vh|rem|em)?$/,
      (match) => ({
        'margin-top': `${match[1]}${match[2] || 'px'}`,
        'margin-bottom': `${match[1]}${match[2] || 'px'}`
      })
    ],
    [
      /^mt-(-?\d+)(px|%|vw|vh|rem|em)?$/,
      (match) => ({ 'margin-top': `${match[1]}${match[2] || 'px'}` })
    ],
    [
      /^mr-(-?\d+)(px|%|vw|vh|rem|em)?$/,
      (match) => ({ 'margin-right': `${match[1]}${match[2] || 'px'}` })
    ],
    [
      /^ml-(-?\d+)(px|%|vw|vh|rem|em)?$/,
      (match) => ({ 'margin-left': `${match[1]}${match[2] || 'px'}` })
    ],
    [
      /^mb-(-?\d+)(px|%|vw|vh|rem|em)?$/,
      (match) => ({ 'margin-bottom': `${match[1]}${match[2] || 'px'}` })
    ],

    [/^p-(\d+)(px|%|vw|vh|rem|em)?$/, (match) => ({ padding: `${match[1]}${match[2] || 'px'}` })],
    [
      /^px-(\d+)(px|%|vw|vh|rem|em)?$/,
      (match) => ({
        'padding-right': `${match[1]}px`,
        'padding-left': `${match[1]}${match[2] || 'px'}`
      })
    ],
    [
      /^py-(\d+)(px|%|vw|vh|rem|em)?$/,
      (match) => ({
        'padding-top': `${match[1]}${match[2] || 'px'}`,
        'padding-bottom': `${match[1]}${match[2] || 'px'}`
      })
    ],
    [
      /^pt-(\d+)(px|%|vw|vh|rem|em)?$/,
      (match) => ({ 'padding-top': `${match[1]}${match[2] || 'px'}` })
    ],
    [
      /^pr-(\d+)(px|%|vw|vh|rem|em)?$/,
      (match) => ({ 'padding-right': `${match[1]}${match[2] || 'px'}` })
    ],
    [
      /^pl-(\d+)(px|%|vw|vh|rem|em)?$/,
      (match) => ({ 'padding-left': `${match[1]}${match[2] || 'px'}` })
    ],
    [
      /^pb-(\d+)(px|%|vw|vh|rem|em)?$/,
      (match) => ({ 'padding-bottom': `${match[1]}${match[2] || 'px'}` })
    ],

    [
      /^radius-(\d+)(px|%|vw|vh|rem|em)?$/,
      (match) => ({ 'border-radius': `${match[1]}${match[2] || 'px'}` })
    ],
    [
      /^radius1-(\d+)(px|%|vw|vh|rem|em)?$/,
      (match) => ({ 'border-top-left-radius': `${match[1]}${match[2] || 'px'}` })
    ],
    [
      /^radius2-(\d+)(px|%|vw|vh|rem|em)?$/,
      (match) => ({ 'border-top-right-radius': `${match[1]}${match[2] || 'px'}` })
    ],
    [
      /^radius3-(\d+)(px|%|vw|vh|rem|em)?$/,
      (match) => ({ 'border-bottom-right-radius': `${match[1]}${match[2] || 'px'}` })
    ],
    [
      /^radius4-(\d+)(px|%|vw|vh|rem|em)?$/,
      (match) => ({ 'border-bottom-left-radius': `${match[1]}${match[2] || 'px'}` })
    ],

    [
      /^translateX\((-?\d+)(px|%|vw|vh|rem|em)?\)$/,
      (match) => ({ transform: `translateX(${match[1]}${match[2] || 'px'})` })
    ],
    [
      /^translateY\((-?\d+)(px|%|vw|vh|rem|em)?\)$/,
      (match) => ({ transform: `translateY(${match[1]}${match[2] || 'px'})` })
    ],

    [/^b-(\d+)-{(.*)}$/, (match) => ({ border: `${match[1]}px solid ${match[2]}` })],
    [/^bt-(\d+)-{(.*)}$/, (match) => ({ 'border-top': `${match[1]}px solid ${match[2]}` })],
    [/^br-(\d+)-{(.*)}$/, (match) => ({ 'border-right': `${match[1]}px solid ${match[2]}` })],
    [/^bl-(\d+)-{(.*)}$/, (match) => ({ 'border-left': `${match[1]}px solid ${match[2]}` })],
    [/^bb-(\d+)-{(.*)}$/, (match) => ({ 'border-bottom': `${match[1]}px solid ${match[2]}` })],

    [/^(static|fixed|absolute|relative|sticky)$/, (match) => ({ position: `${match[1]}` })],

    [
      /^text-(\d+)(px|%|vw|vh|rem|em)?$/,
      (match) => ({ 'font-size': `${match[1]}${match[2] || 'px'}` })
    ],
    [/^weight-(\d+)$/, (match) => ({ 'font-weight': `${match[1]}` })],

    [/^w-(\d+)(px|%|vw|vh|rem|em)?$/, (match) => ({ width: `${match[1]}${match[2] || 'px'}` })],
    [/^h-(\d+)(px|%|vw|vh|rem|em)?$/, (match) => ({ height: `${match[1]}${match[2] || 'px'}` })],
    [
      /^max-w-(\d+)(px|%|vw|vh|rem|em)?$/,
      (match) => ({ 'max-width': `${match[1]}${match[2] || 'px'}` })
    ],
    [
      /^min-w-(\d+)(px|%|vw|vh|rem|em)?$/,
      (match) => ({ 'min-width': `${match[1]}${match[2] || 'px'}` })
    ],
    [
      /^max-h-(\d+)(px|%|vw|vh|rem|em)?$/,
      (match) => ({ 'max-height': `${match[1]}${match[2] || 'px'}` })
    ],
    [
      /^min-h-(\d+)(px|%|vw|vh|rem|em)?$/,
      (match) => ({ 'min-height': `${match[1]}${match[2] || 'px'}` })
    ],

    ['w-full', { width: '100%' }],
    ['h-full', { height: '100%' }],
    ['full', { height: '100%', width: '100%' }],

    ['flex', { display: 'flex' }],
    ['flex-1', { flex: 1 }],
    ['jc-center', { 'justify-content': 'center' }],
    ['jc-between', { 'justify-content': 'space-between' }],
    ['ai-center', { 'align-items': 'center' }],
    [/^fg-(\d+)$/, (match) => ({ 'flex-grow': `${match[1]}` })],
    [/^fs-(\d+)$/, (match) => ({ 'flex-shrink': `${match[1]}` })],
    [/^fd-(column|row|revert|row-reverse+)$/, (match) => ({ 'flex-direction': `${match[1]}` })],

    [/^overflow-y-(auto|hidden|scrol+)$/, (match) => ({ 'overflow-y': `${match[1]}` })],
    [/^overflow-x-(auto|hidden|scrol+)$/, (match) => ({ 'overflow-x': `${match[1]}` })],
    [/^overflow-(auto|hidden|scroll+)$/, (match) => ({ overflow: `${match[1]}` })],
    [/^shadow-(\d+)$/, (match) => ({ 'box-shadow': `0 0 ${match[1]}px rgba(0, 0, 0, 0.1)` })],
    // 特别注意.* 容易影响其他匹配
    [/^bgc-{(.*)}$/, (match) => ({ 'background-color': `${match[1]}` })],
    [/^color-{(.*)}$/, (match) => ({ color: `${match[1]}` })]
  ]
}

```

