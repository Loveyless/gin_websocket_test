

// 官网的组合式函数例子 Composable
// 这个函数将添加和清除 DOM 事件监听器的逻辑也封装进一个组合式函数中

// import { onMounted,onUnmounted } from "vue";

export function useEventListener(target: Window & typeof globalThis, event: string, callbakc:any) {
  onMounted(() => {
    target.addEventListener(event, callbakc);
  });
  onUnmounted(() => {
    target.removeEventListener(event, callbakc);
    // alert("卸载了");
  });
}