import axios, { AxiosRequestConfig, Canceler } from "axios";
import { isFunction } from "@/utils/is";
import qs from "qs";

// 当请求方式、请求 URL 地址和请求参数都一样时，我们就可以认为请求是一样的。
// 因此在每次发起请求时，我们就可以根据当前请求的请求方式、请求 URL 地址和请求参数来生成一个唯一的 key 
// 同时为每个请求创建一个专属的 CancelToken，然后把 key 和 cancel 函数以键值对的形式保存到 Map 对象中
// 使用 Map 的好处是可以快速的判断是否有重复的请求


// * 声明一个 Map 用于存储每个请求的标识 和 取消函数
let pendingMap = new Map<string, Canceler>();

// * 序列化参数
export const getPendingUrl = (config: AxiosRequestConfig) => [config.method, config.url, qs.stringify(config.data), qs.stringify(config.params)].join("&");

export class AxiosCanceler {
  /**
   * @description: 添加请求
   * @param {Object} config
   * @return void
   */
  addPending(config: AxiosRequestConfig) {
    // * 在请求开始前，对之前的请求做检查取消操作
    this.removePending(config);
    // 获取序列化
    const url = getPendingUrl(config);
    console.log("序列化参数",url);
    // 如果没有就赋值一个cancelToken
    config.cancelToken = config.cancelToken || new axios.CancelToken((cancel) => {
        //斌且如果map中没有则添加进去
        if (!pendingMap.has(url)) {
          pendingMap.set(url, cancel);
        }
      });
  }

  /**
   * @description: 移除请求
   * @param {Object} config
   */
  removePending(config: AxiosRequestConfig) {
    const url = getPendingUrl(config);

    if (pendingMap.has(url)) {
      // 如果在 pending 中存在当前请求标识，需要取消当前请求，并且移除
      const cancel = pendingMap.get(url);
      cancel && isFunction(cancel) && cancel();
      pendingMap.delete(url);
    }
  }


  /**
   * @description: 清空所有pending
   */
  removeAllPending() {
    pendingMap.forEach((cancel) => {
      cancel && isFunction(cancel) && cancel();
    });
    pendingMap.clear();
  }

  /**
   * @description: 重置
   */
  reset(): void {
    pendingMap = new Map<string, Canceler>();
  }
}


//函数式


// //用于根据当前请求的信息，生成请求 Key；
// function generateReqKey(config:AxiosRequestConfig) {
//   const { method, url, params, data } = config;
//   console.log("生产的请求key", [method, url, qs.stringify(params), qs.stringify(data)].join("&"));
//   return [method, url, qs.stringify(params), qs.stringify(data)].join("&");
// }

// //用于把当前请求信息添加到pendingRequest对象中；
// const pendingList = new Map();
// function addPendingRequest(config:AxiosRequestConfig) {
//   //创建一个key
//   const requestKey = generateReqKey(config);
//   //如果有的话添加 没有的话 创建一个赋值
//   config.cancelToken = config.cancelToken || new axios.CancelToken((cancel:any) => {
//     //如果没有的话赋值 有的话无事发生
//     if (!pendingList.has(requestKey)) {
//       pendingList.set(requestKey, cancel);
//     }
//   });
// }

// // 检查是否存在重复请求，若存在则取消已发的请求。
// function removePendingRequest(config:AxiosRequestConfig) {
//   // 拿到key
//   const requestKey = generateReqKey(config);
//   // 如果有这个key了已经
//   if (pendingList.has(requestKey)) {
//     // 拿到这个值
//     const cancelToken = pendingList.get(requestKey);
//     // 取消请求
//     cancelToken(requestKey);
//     // 删除对应key
//     pendingList.delete(requestKey);
//   }
// }
