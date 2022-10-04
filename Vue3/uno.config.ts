//这里的(-?\d+)可以把负数也匹配上！牛比
//[/^p-(-?\d+)(px|%|vw|vh|rem|em)?$/, (d: any) => ({ padding: `${d / 4}px` })],
//函数的第一个参数是匹配结果，您可以对其进行解构以获取匹配的组。

const width_height: any = [
  [
    /^w-(-?\d+)(px|%|vw|vh|rem|em)?$/,
    (d: any) => ({
      width: `${d[1]}${d[2] || "px"}`,
    }),
  ],
  [
    /^h-(-?\d+)(px|%|vw|vh|rem|em)?$/,
    (d: any) => ({
      "height": `${d[1]}${d[2] || "px"}`,
    }),
  ]
];

const margin: any = [
  [
    /^m-(-?\d+)(px|%|vw|vh|rem|em)?$/,
    (d: any) => ({
      margin: `${d[1]}${d[2] || "px"}`,
    }),
  ],
  [
    /^mt-(-?\d+)(px|%|vw|vh|rem|em)?$/,
    (d: any) => ({
      "margin-top": `${d[1]}${d[2] || "px"}`,
    }),
  ],
  [
    /^mr-(-?\d+)(px|%|vw|vh|rem|em)?$/,
    (d: any) => ({
      "margin-right": `${d[1]}${d[2] || "px"}`,
    }),
  ],
  [
    /^mb-(-?\d+)(px|%|vw|vh|rem|em)?$/,
    (d: any) => ({
      "margin-bottom": `${d[1]}${d[2] || "px"}`,
    }),
  ],
  [
    /^ml-(-?\d+)(px|%|vw|vh|rem|em)?$/,
    (d: any) => ({
      "margin-left": `${d[1]}${d[2] || "px"}`,
    }),
  ],
  [
    /^mx-(-?\d+)(px|%|vw|vh|rem|em)?$/,
    (d) => ({
      "margin-right": `${d[1]}${d[2] || "px"}`,
      "margin-left": `${d[1]}${d[2] || "px"}`,
    }),
  ],
  [
    /^my-(-?\d+)(px|%|vw|vh|rem|em)?$/,
    (d) => ({
      "margin-top": `${d[1]}${d[2] || "px"}`,
      "margin-bottom": `${d[1]}${d[2] || "px"}`,
    }),
  ],
];

const padding: any = [
  [
    /^p-(-?\d+)(px|%|vw|vh|rem|em)?$/,
    (d: any) => ({
      padding: `${d[1]}${d[2] || "px"}`,
    }),
  ],
  [
    /^pt-(-?\d+)(px|%|vw|vh|rem|em)?$/,
    (d: any) => ({
      "padding-top": `${d[1]}${d[2] || "px"}`,
    }),
  ],
  [
    /^pr-(-?\d+)(px|%|vw|vh|rem|em)?$/,
    (d: any) => ({
      "padding-right": `${d[1]}${d[2] || "px"}`,
    }),
  ],
  [
    /^pb-(-?\d+)(px|%|vw|vh|rem|em)?$/,
    (d: any) => ({
      "padding-bottom": `${d[1]}${d[2] || "px"}`,
    }),
  ],
  [
    /^pl-(-?\d+)(px|%|vw|vh|rem|em)?$/,
    (d: any) => ({
      "padding-left": `${d[1]}${d[2] || "px"}`,
    }),
  ],
  [
    /^px-(-?\d+)(px|%|vw|vh|rem|em)?$/,
    (d) => ({
      "padding-right": `${d[1]}${d[2] || "px"}`,
      "padding-left": `${d[1]}${d[2] || "px"}`,
    }),
  ],
  [
    /^py-(-?\d+)(px|%|vw|vh|rem|em)?$/,
    (d) => ({
      "padding-top": `${d[1]}${d[2] || "px"}`,
      "padding-bottom": `${d[1]}${d[2] || "px"}`,
    }),
  ],
];

const flex: any = [
  [
    "flex", // 使用时只需要写 p-c 即可应用该组样式
    {
      display: "flex",
      "flex-flow": "row wrap", //简写属性
      "justify-content": "space-evenly",
      "align-items": "center",
      "justify-items": "stretch", //https://developer.mozilla.org/zh-CN/docs/Web/CSS/justify-items
      "align-content": "space-evenly", //行与行之间的
    },
  ],
];

export default [
  // body 函数的第一个参数是匹配结果，您可以对其进行解构以获取匹配的组。

  ...width_height,
  ...margin,
  ...padding,
  ...flex,

  // 在这个可以增加预设规则, 也可以使用正则表达式
  // [
  //   "xxxxxxx-xxxxx", // 使用时只需要写 p-c 即可应用该组样式
  //   {
  //     position: "absolute",
  //     top: "50%",
  //     left: "50%",
  //     transform: `translate(-50%, -50%)`,
  //   },
  // ],
  // [/^m-(\d+)$/, ([, d]: any) => ({ margin: `${d / 4}rem` })],
];
