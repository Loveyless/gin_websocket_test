//element plus loading
import { ElLoading } from "element-plus";

//这句话还没怎么理解
let LoadingTimer: ReturnType<typeof ElLoading.service>;


export const openLoading = (dark: boolean = false) => {
  //开loading
  if (!dark) {
    LoadingTimer = ElLoading.service({ fullscreen: true, lock: true });
  } else {
    LoadingTimer = ElLoading.service({
      lock: true,
      text: "Loading",
      background: "rgba(0, 0, 0, 0.7)",
    });
  }
};


export const closeLoading = () => {
  //关loading
  LoadingTimer && LoadingTimer.close();
};
