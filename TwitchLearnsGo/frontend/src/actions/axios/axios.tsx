import axios from "axios";

export const axiosApiInstance = axios.create({
  baseURL: "http://localhost:4000",
  timeout: 10000,
});

// Response interceptor for API calls
axiosApiInstance.interceptors.response.use(
  (response) => {
    return response;
  },
  async function (error) {
    const originalRequest = error.config;
    console.log("Re-try axios logic");
    if (error.response?.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;
      return axiosApiInstance(originalRequest);
    }
    return Promise.reject(error);
  }
);
