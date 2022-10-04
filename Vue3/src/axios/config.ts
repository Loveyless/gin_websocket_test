import axios from "axios";

let instance = axios.create({
  baseURL: import.meta.env.VITE_BASE_URL as string,
  // timeout: 1000,
  // headers: { "X-Custom-Header": "foobar" },
  // 跨域时候允许携带凭证
  withCredentials: true,

  // 允许在向服务器发送前，修改请求数据
  // axios默认是Request Payload格式，加了transformRequest会默认变成form Data格式
  //需要自己再转一下变回Request Payload 也就是下面的headers里的Content-Type
  // transformRequest: [
  //   function (data) {
  //     // 对 data 进行任意转换处理

  //     return data;
  //   },
  // ],
  // //再转一下变回Request Payload
  // headers: {
  //   "Content-Type": "application/json;charset=UTF-8",
  // },

  // 在传递给 then/catch 前，允许修改响应数据
  // transformResponse: [
  //   function (data) {
  //     // 对 data 进行任意转换处理

  //     return JSON.parse(data);
  //   },
  // ],

  // `onUploadProgress` 允许为上传处理进度事件
  onUploadProgress: function (progressEvent) {
    // 对原生进度事件的处理
  },

  // `onDownloadProgress` 允许为下载处理进度事件
  onDownloadProgress: function (progressEvent) {
    // 对原生进度事件的处理
  },

  // `cancelToken` 指定用于取消请求的 cancel token
  // cancelToken: new CancelToken(function (cancel) {}),
});

export default instance;
