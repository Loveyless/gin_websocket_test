
//过滤空 字符串 null undefined
export function filterObject(obj: any) {
  let body: any = {};

  for (let key in obj) {
    if (obj[key] != "" && null && undefined) {
      body[key] = obj[key];
    }
  }

  return body;
}
