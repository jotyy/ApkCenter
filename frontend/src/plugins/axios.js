import axios from "axios";

export default () => {
  return axios.create({
    timeout: 60000,
    baseURL: `http://192.168.100.132:8686/api/v1`,  // 注意IP地址必须修改
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
};
