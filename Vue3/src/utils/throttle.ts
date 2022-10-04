export function throttle(fn: Function, delayTime = 1000) {
  //节流阀
  let swi = true;
  return function (this: any) {
    //如果关闭状态就退出
    if (!swi) return;

    //没关闭的话先关闭
    swi = false;

    //再开定时器
    setTimeout(() => {
      fn.apply(this, arguments);
      //开启节流阀
      swi = true;
    }, delayTime);
  };
}
