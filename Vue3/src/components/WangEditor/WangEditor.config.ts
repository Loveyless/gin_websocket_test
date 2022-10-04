import type { IToolbarConfig, IEditorConfig, SlateElement } from "@wangeditor/editor";
import type { AttachmentElement } from "@wangeditor/plugin-upload-attachment";

// 编辑器配置

// 链接

// 自定义校验链接
function customCheckLinkFn(text: string, url: string): string | boolean | undefined {
  if (!url) {
    return "请正确输入";
  }
  if (url.indexOf("http") !== 0) {
    return "链接必须以 http/https 开头";
  }
  return true;
  // 返回值有三种选择：
  // 1. 返回 true ，说明检查通过，编辑器将正常插入链接
  // 2. 返回一个字符串，说明检查未通过，编辑器会阻止插入。会 alert 出错误信息（即返回的字符串）
  // 3. 返回 undefined（即没有任何返回），说明检查未通过，编辑器会阻止插入。但不会提示任何信息
}
// 自定义转换链接 url
function customParseLinkUrl(url: string): string {
  if (url.indexOf("http") !== 0) {
    return `http://${url}`;
  }
  return url;
}

// 网络图片链接

// 自定义校验图片
function customCheckImageFn(src: string, alt: string, url: string): boolean | undefined | string {
  if (!src) {
    return;
  }
  // if (src.indexOf("http") !== 0) {
  //   return "图片网址必须以 http/https 开头";
  // }
  return true;

  // 返回值有三种选择：
  // 1. 返回 true ，说明检查通过，编辑器将正常插入图片
  // 2. 返回一个字符串，说明检查未通过，编辑器会阻止插入。会 alert 出错误信息（即返回的字符串）
  // 3. 返回 undefined（即没有任何返回），说明检查未通过，编辑器会阻止插入。但不会提示任何信息
}

// 转换网络图片链接
function customParseImageSrc(src: string): string {
  // if (src.indexOf("http") !== 0) {
  //   return `http://${src}`;
  // }
  return src;
}

//上传图片
const uploadImage: any = {
  server: "/api/upload",

  // form-data fieldName ，默认值 'wangeditor-uploaded-image'
  fieldName: "your-custom-name",

  // 单个文件的最大体积限制，默认为 2M
  maxFileSize: 50 * 1024 * 1024, // 50M

  // 小于该值就插入 base64 格式（而不上传），默认为 0
  base64LimitSize: 5 * 1024, // 5kb

  // 最多可上传几个文件，默认为 100
  maxNumberOfFiles: 50,

  // 选择文件时的类型限制，默认为 ['image/*'] 。如不想限制，则设置为 []
  allowedFileTypes: [],

  // 自定义上传参数，例如传递验证的 token 等。参数会被添加到 formData 中，一起上传到服务端。
  meta: {
    token: "xxx",
    otherKey: "yyy",
  },

  // 将 meta 拼接到 url 参数中，默认 false
  metaWithUrl: false,

  // 自定义增加 http  header
  headers: {
    Accept: "text/x-json",
    otherKey: "xxx",
  },

  // 跨域是否传递 cookie ，默认为 false
  withCredentials: false,

  // 超时时间，默认为 10 秒
  timeout: 5 * 1000, // 5 秒

  // 上传之前触发
  onBeforeUpload(file: File) {
    // file 选中的文件，格式如 { key: file }
    return file;

    // 可以 return
    // 1. return file 或者 new 一个 file ，接下来将上传
    // 2. return false ，不上传这个 file
  },

  // 上传进度的回调函数
  onProgress(progress: number) {
    // progress 是 0-100 的数字
    console.log("progress", progress);
  },

  // 单个文件上传成功之后
  onSuccess(file: File, res: any) {
    console.log(`${file.name} 上传成功`, res);
  },

  // 单个文件上传失败
  onFailed(file: File, res: any) {
    console.log(`${file.name} 上传失败`, res);
  },

  // 上传错误，或者触发 timeout 超时
  onError(file: File, err: any, res: any) {
    console.log(`${file.name} 上传出错`, err, res);
  },
};

// 自定义校验网络视频
function customCheckVideoFn(src: string): boolean | string | undefined {
  if (!src) {
    return;
  }
  if (src.indexOf("http") !== 0) {
    return "视频地址必须以 http/https 开头";
  }
  return true;

  // 返回值有三种选择：
  // 1. 返回 true ，说明检查通过，编辑器将正常插入视频
  // 2. 返回一个字符串，说明检查未通过，编辑器会阻止插入。会 alert 出错误信息（即返回的字符串）
  // 3. 返回 undefined（即没有任何返回），说明检查未通过，编辑器会阻止插入。但不会提示任何信息
}

// 自定义转换网络视频
function customParseVideoSrc(src: string): string {
  if (src.includes(".bilibili.com")) {
    // 转换 bilibili url 为 iframe （仅作为示例，不保证代码正确和完整）
    // const arr = location.pathname.split('/')
    // const vid = arr[arr.length - 1]
    // return `<iframe src="//player.bilibili.com/player.html?bvid=${vid}" scrolling="no" border="0" frameborder="no" framespacing="0" allowfullscreen="true"> </iframe>`
  }
  return src;
}

// 上传视频
const uploadVideo: any = {
  // 服务端地址
  server: "/api/upload",

  // form-data fieldName ，默认值 'wangeditor-uploaded-video'
  fieldName: "your-custom-name",

  // 单个文件的最大体积限制，默认为 10M
  maxFileSize: 200 * 1024 * 1024, // 200M

  // 最多可上传几个文件，默认为 5
  maxNumberOfFiles: 5,

  // 选择文件时的类型限制，默认为 ['video/*'] 。如不想限制，则设置为 []
  allowedFileTypes: [],

  // 自定义上传参数，例如传递验证的 token 等。参数会被添加到 formData 中，一起上传到服务端。
  meta: {
    token: "xxx",
    otherKey: "yyy",
  },

  // 将 meta 拼接到 url 参数中，默认 false
  metaWithUrl: false,

  // 自定义增加 http  header
  headers: {
    Accept: "text/x-json",
    otherKey: "xxx",
  },

  // 跨域是否传递 cookie ，默认为 false
  withCredentials: false,

  // 超时时间，默认为 30 秒
  timeout: 15 * 1000, // 15 秒

  // 视频不支持 base64 格式插入

  // 上传之前触发
  onBeforeUpload(file: File) {
    // file 选中的文件，格式如 { key: file }
    return file;

    // 可以 return
    // 1. return file 或者 new 一个 file ，接下来将上传
    // 2. return false ，不上传这个 file
  },

  // 上传进度的回调函数
  onProgress(progress: number) {
    // progress 是 0-100 的数字
    console.log("progress", progress);
  },

  // 单个文件上传成功之后
  onSuccess(file: File, res: any) {
    console.log(`${file.name} 上传成功`, res);
  },

  // 单个文件上传失败
  onFailed(file: File, res: any) {
    console.log(`${file.name} 上传失败`, res);
  },

  // 上传错误，或者触发 timeout 超时
  onError(file: File, err: any, res: any) {
    console.log(`${file.name} 上传出错`, err, res);
  },
};

// 上传附件
const uploadAttachment: any = {
  server: "/api/upload", // 服务端地址
  timeout: 5 * 1000, // 5s

  fieldName: "custom-fileName",
  meta: { token: "xxx", a: 100 }, // 请求时附加的数据
  metaWithUrl: true, // meta 拼接到 url 上
  headers: { Accept: "text/x-json" },

  maxFileSize: 10 * 1024 * 1024, // 10M

  onBeforeUpload(file: File) {
    console.log("onBeforeUpload", file);
    return file; // 上传 file 文件
    // return false // 会阻止上传
  },
  onProgress(progress: number) {
    console.log("onProgress", progress);
  },
  onSuccess(file: File, res: any) {
    console.log("onSuccess", file, res);
  },
  onFailed(file: File, res: any) {
    alert(res.message);
    console.log("onFailed", file, res);
  },
  onError(file: File, err: Error, res: any) {
    alert(err.message);
    console.error("onError", file, err, res);
  },

  // // 上传成功后，用户自定义插入文件
  // customInsert(res: any, file: File, insertFn: Function) {
  //   console.log('customInsert', res)
  //   const { url } = res.data || {}
  //   if (!url) throw new Error(`url is empty`)

  //   // 插入附件到编辑器
  //   insertFn(`customInsert-${file.name}`, url)
  // },

  // // 用户自定义上传
  // customUpload(file: File, insertFn: Function) {
  //   console.log('customUpload', file)

  //   return new Promise(resolve => {
  //     // 插入一个文件，模拟异步
  //     setTimeout(() => {
  //       const src = `https://www.w3school.com.cn/i/movie.ogg`
  //       insertFn(`customUpload-${file.name}`, src)
  //       resolve('ok')
  //     }, 500)
  //   })
  // },

  // // 自定义选择
  // customBrowseAndUpload(insertFn: Function) {
  //   alert('自定义选择文件，如弹出图床')
  //   // 自己上传文件
  //   // 上传之后用 insertFn(fileName, link) 插入到编辑器
  // },

  // 插入到编辑器后的回调
  onInsertedAttachment(elem: AttachmentElement) {
    console.log("inserted attachment", elem);
  },

  // 其他...
};
export default {
  customCheckLinkFn,
  customParseLinkUrl,
  customCheckImageFn,
  customParseImageSrc,
  uploadImage,
  customCheckVideoFn,
  customParseVideoSrc,
  uploadVideo,
  uploadAttachment,
};
