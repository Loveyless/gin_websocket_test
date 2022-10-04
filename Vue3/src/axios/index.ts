import type { AxiosInstance, AxiosError, AxiosRequestConfig, AxiosResponse } from "axios";
import instance from "@/axios/config";
import { openLoading, closeLoading } from "@/hooks/loading";
import NProgress from "@/hooks/nprogress";
import { Loading } from "@element-plus/icons-vue";
import { AxiosCanceler } from "@/axios/cancel";
import { ElMessage } from "element-plus";

//实例化取消请求
const axiosCanceler = new AxiosCanceler();

//拦截
instance.interceptors.request.use(
  (config: AxiosRequestConfig) => {
    const globalStore = GlobalStore();
    // 将当前请求添加到 pending 中
    axiosCanceler.addPending(config);

    //全局loading 和 进度条 请求头中 { loading : 0 , loadingdark : 0 , nprogress : 0 }  0为不显示 1为显示
    //默认需要进度条
    config.headers!.nprogress != 0 && NProgress.start();
    if (config.headers!.loading != 0) {
      //默认需要loading
      if (config.headers!.loadingdark == 0) {
        //默认黑色loading
        openLoading();
      } else {
        openLoading(true); //白色loading
      }
    }

    // 打印请求
    console.log("-------------------");
    console.log("地址", config.method, config.url);
    console.log("请求头", config.headers);
    console.log("请求参数", config.data);

    return { ...config };
  },
  (err: AxiosError) => {
    closeLoading();
    NProgress.done();

    //前置请求错误
    console.log("前置请求错误", err.message);
    return Promise.reject(err);
  }
);

//响应
instance.interceptors.response.use(
  (response: AxiosResponse) => {
    const { data, config } = response;
    // 在请求结束后，移除本次请求
    axiosCanceler.removePending(config);

    closeLoading();
    NProgress.done();

    // 打印返回值
    console.log("请求结果", data);
    // 根据请求弹窗 可删除
    if (data.status == 200) {
      ElMessage({
        message: data.message,
        type: "success",
      });
    } else {
      ElMessage({
        message: data.message,
        type: "error",
      });
    }
    return response;
  },
  (err: AxiosError) => {
    closeLoading();
    NProgress.done();

    //后置请求错误
    console.log("后置请求错误", err);
    console.log("后置错误详情", err.response);
    return Promise.reject(err);
  }
);

export const http = instance;
