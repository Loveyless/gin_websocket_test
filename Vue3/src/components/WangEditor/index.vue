<template>
  <div>
    <Toolbar style="border-bottom: 1px solid #ccc" :editor="editorRef" :defaultConfig="toolbarConfig" :mode="mode" />
    <Editor :style='{ height : props.height }' v-model="valueHtml" :defaultConfig="editorConfig" :mode="mode"
      @onCreated="handleCreated" @onChange="handleChange" @onDestroyed="handleDestroyed" @onFocus="handleFocus"
      @onBlur="handleBlur" @customAlert="customAlert" @customPaste="customPaste" />
  </div>
</template>

<script lang='ts' setup name="wangeditor">
import '@wangeditor/editor/dist/css/style.css' // 引入 css
import { Editor, Toolbar, } from '@wangeditor/editor-for-vue';
import type { IToolbarConfig, IEditorConfig, SlateElement } from '@wangeditor/editor'
import { Boot } from '@wangeditor/editor'
import attachmentModule from '@wangeditor/plugin-upload-attachment'
// import ctrlEnterModule from '@wangeditor/plugin-ctrl-enter'
import markdownModule from '@wangeditor/plugin-md'
//配置文件
import WConfig from "./WangEditor.config"
const global = GlobalStore()

//上传附件插件注册 https://www.wangeditor.com/v5/plugins.html
Boot.registerModule(attachmentModule)
//ctrl+enter 换行插件
// Boot.registerModule(ctrlEnterModule)
//markdown 插件
Boot.registerModule(markdownModule)

// 编辑器实例，必须用 shallowRef
const editorRef = shallowRef()

// 内容 HTML
const valueHtml = ref('')

// mode
const mode = "default"


// onMounted(() => {
//   setTimeout(() => {
//     valueHtml.value = '<p>模拟 Ajax 异步设置内容</p>'
//   }, 1500)
// })

// 工具栏配置 Partial 使得T中的所有属性都是可选的
const toolbarConfig: Partial<IToolbarConfig> = {
  // 插入哪些菜单
  insertKeys: {
    index: 10, // 自定义插入的位置
    keys: ['uploadAttachment'], // “上传附件”菜单
  },
}


// 编辑器配置
const editorConfig: Partial<IEditorConfig> = {
  // 配置编辑器 placeholder
  placeholder: '请输入内容...',
  // 配置编辑器是否只读，默认为 false 只读状态可通过 editor.enable() 和 editor.disable() 切换，详见 https://www.wangeditor.com/v5/API.html
  readOnly: false,
  // 配置编辑器默认是否 focus ，默认为 true
  autoFocus: true,
  // 配置编辑器是否支持滚动，默认为 true 。注意，此时不要固定 editor- container 的高度，设置一个 min-height 即可。
  scroll: true,
  // 在编辑器中，点击选中“附件”节点时，要弹出的菜单
  hoverbarKeys: {
    attachment: {
      menuKeys: ['downloadAttachment'], // “下载附件”菜单
    },
  },

  //菜单设置
  MENU_CONF: {
    //插入链接
    insertLink: {
      checkLink: WConfig.customCheckLinkFn, // 也支持 async 函数
      parseLinkUrl: WConfig.customParseLinkUrl, // 也支持 async 函数
    },
    //更新链接
    editLink: {
      checkLink: WConfig.customCheckLinkFn, // 也支持 async 函数
      parseLinkUrl: WConfig.customParseLinkUrl, // 也支持 async 函数
    },
    //插入网络图片
    insertImage: {
      onInsertedImage(imageNode: ImageElement | null) {  // TS 语法
        if (imageNode == null) return
        const { src, alt, url, href } = imageNode
        console.log('inserted image', src, alt, url, href)
        let imageList1: any = []
        imageList1.push(imageNode)
        console.log("imglist", imageList1);
      },
      checkImage: WConfig.customCheckImageFn, // 也支持 async 函数
      parseImageSrc: WConfig.customParseImageSrc, // 也支持 async 函数
    },
    editImage: {
      onUpdatedImage(imageNode: ImageElement | null) {  // TS 语法
        if (imageNode == null) return
        const { src, alt, url } = imageNode
        console.log('updated image', src, alt, url)
      },
      checkImage: WConfig.customCheckImageFn, // 也支持 async 函数
      parseImageSrc: WConfig.customParseImageSrc, // 也支持 async 函数
    },
    //上传图片 这里对返回值有特殊要求 具体可看 https://www.wangeditor.com/v5/menu-config.html#%E6%9C%8D%E5%8A%A1%E7%AB%AF%E5%9C%B0%E5%9D%80
    uploadImage: WConfig.uploadImage,
    // 网络视频
    insertVideo: {
      onInsertedVideo(videoNode: VideoElement | null) {  // TS 语法
        // onInsertedVideo(videoNode) {                    // JS 语法
        if (videoNode == null) return

        const { src } = videoNode
        console.log('inserted video', src)
      },
      checkVideo: WConfig.customCheckVideoFn, // 也支持 async 函数
      parseVideoSrc: WConfig.customParseVideoSrc, // 也支持 async 函数
    },
    //上传视频
    uploadVideo: WConfig.uploadVideo,
    // 代码高亮
    codeSelectLang: [
      { text: 'CSS', value: 'css' },
      { text: 'HTML', value: 'html' },
      { text: 'XML', value: 'xml' },
    ]
  }
}



// 生命周期

// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
  const editor = editorRef.value
  if (editor == null) return
  editor.destroy()
})
//创建富文本时保存实例
const handleCreated = (editor: any) => {
  editorRef.value = editor // 记录 editor 实例，重要！
}
const handleChange = (editor: any) => {
  console.log('change:', editor.children)
  editor.handleTab = () => editor.insertText('    ')
}
const handleDestroyed = (editor: any) => { console.log('destroyed', editor) }
const handleFocus = (editor: any) => { console.log('focus', editor) }
const handleBlur = (editor: any) => { console.log('blur', editor) }
const customAlert = (info: any, type: any) => { alert(`【自定义提示】${type} - ${info}`) }
const customPaste = (editor: any, event: any, callback: any) => {
  // const html = event.clipboardData.getData('text/html') // 获取粘贴的 html
  // const text = event.clipboardData.getData('text/plain') // 获取粘贴的纯文本
  // const rtf = event.clipboardData.getData('text/rtf') // 获取 rtf 数据（如从 word wsp 复制粘贴）


  //这个功能好像富文本自带。。。 就等上传图片的接口了
  //复制图片上传功能 这里直接复制如果小于 base64配置项 就直接base64存入html模板
  // const imgData: any = event.clipboardData.items  //复制过来的值里面的clipboardData
  // let file: any = null
  // // console.log('ClipboardEvent 粘贴事件对象', imgData)
  // for (let i = 0; i < imgData.length; i++) {      //有可能图片文字一起复制过来的 遍历每一项
  //   console.log(imgData[i]);                      //DataTransferItem {kind: 'file', type: 'image/png'}
  //   if (imgData[i].type.includes("image")) {      //取到img的哪项
  //     file = imgData[i].getAsFile();              //上面有这个方法返回一个file对象
  //     // console.log(file,typeof file);
  //     break;
  //   }
  // }
  // if(file){
  //   let render = new FileReader()
  //   render.onload = function(event:any) {
  //     let img = document.createElement("img")
  //     img.src = event.target.result //这就是图片的base64地址？不太懂
  //   }
  //   render.readAsDataURL(file)
  // }


  // 自定义插入内容
  // editor.insertText('xxx')

  // 返回 false ，阻止默认粘贴行为
  // event.preventDefault()
  // callback(false) // 返回值（注意，vue 事件的返回值，不能用 return）

  // 返回 true ，继续默认的粘贴行为
  callback(true)
}







/**
 * 以下为组件 接受项 导出项
 */

const props = withDefaults(defineProps<{ height?: string }>(), { height: '700px' })

//更改编辑器html 字符串
function setEditorHtml(html: string) {
  valueHtml.value = html
}

//导出 模板内容 和 更改函数 获取时无需.value 直接 editorRef.value.valueHtml
defineExpose({
  editorRef,  //实例
  valueHtml,  //内容
  setEditorHtml, //更改内容
})
</script> 

<style lang="less" scoped>
@import "./index.less";
</style>