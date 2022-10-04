
// 官网的组合式函数例子 Composable

// import { onMounted,onUnmounted } from "vue";

//结合了另一个组合式函数 useEventLister
import { useEventListener } from "./useEventLister";

export function useMouse() {
  const x = ref<number>(0);
  const y = ref<number>(0);

  const update = (e: MouseEvent) => {
    x.value = e.pageX;
    y.value = e.pageY;
  };

  useEventListener(window, "mousemove", update);

  return { x, y };
}
