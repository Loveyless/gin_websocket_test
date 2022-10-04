import router from "@/router/routes";
import { openLoading, closeLoading } from "@/hooks/loading";
import NProgress from "@/hooks/nprogress";
import { GlobalStore } from "@/store";
import { AxiosCanceler } from "@/axios/cancel";

//实例化取消请求
const axiosCanceler = new AxiosCanceler();

//拦截
router.beforeEach((to, from, next) => {
  const globalStore = GlobalStore();
  //转跳路由之前 清除所有请求
  axiosCanceler.removeAllPending();

  // 判断当前路由是否需要loading切换
  if (!to.matched.some((item) => item.meta.noLoading)) {
    openLoading();
  }

  // 判断当前路由是否需要顶部进度条
  if (!to.matched.some((item) => item.meta.noNProgress)) {
    NProgress.start();
  }

  // 判断当前路由是否需要访问权限
  // if (to.matched.some((item) => item.meta.noRequireAuth)){
  //   //不要权限 直接放行
  //   return next();
  // }else if(globalStore.token){
  //   //有token 放行
  //   return next();
  // }else{
  //   next("/login")
  // }

  next();
});

//响应
router.afterEach((to, from, failure) => {
  closeLoading();
  NProgress.done();
});

export default router;
