import { axiosApiInstance } from "./axios/axios";

export const registerForm = async ({
  username,
  password,
}: {
  username: string;
  password: string;
}) => {
  const body = JSON.stringify({ username, password });

  try {
    const response = await axiosApiInstance.post("/register/user", body);
    const { data } = response;
    return data;
  } catch (error) {
    return false;
  }
};
