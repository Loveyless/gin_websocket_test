/**
 * @description: 二进制形式读取txt文件 并判断其是否为utf-8
 * @param {type}
 * @return:
 */
function readFile(file: any) {
  return new Promise((resolve: any, reject: any) => {
    const reader: any = new FileReader();
    reader.onload = function (evt: any) {
      resolve(evt.target.result);
    };
    reader.readAsArrayBuffer(file);
  });
}

type isUtf8Type = (Fild: any) => Promise<boolean>;

export const isUtf8: isUtf8Type = async function (file: any) {
  const res: any = await readFile(file);
  const firstCode: any = new Uint8Array(res)[0];
  return firstCode >= 33 && firstCode <= 126;
};
