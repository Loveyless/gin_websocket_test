export function debouncer(fn: Function, delayTime: number = 1000) {
  let timer: any = null;
  return function (this: any, ...arr: any) {
    // 这个函数this指向没问题

    console.log("this", this);
    console.log("arr", arr);

    //如果有计时器就清除
    if (timer) clearTimeout(timer);

    //再开定时器
    timer = setTimeout(() => {
      // fn(...arguments); //箭头函数所以可以arguments 要不然这个arguments取的是settimeout函数的arguments
      // 但是要设置this指向 所以直接用apply
      fn.apply(this, arr); //因为用了箭头函数 this直接拿了外面的this
    }, delayTime);
  };
}
