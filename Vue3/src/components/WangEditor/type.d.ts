
// 图片的类型
type ImageElement = SlateElement & {
  src: string;
  alt: string;
  url: string;
  href: string;
};

//视频类型
type VideoElement = SlateElement & {
  src: string;
  poster?: string;
};
