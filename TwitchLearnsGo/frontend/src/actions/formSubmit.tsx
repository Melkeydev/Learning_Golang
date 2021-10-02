import { axiosApiInstance } from "./axios/axios";

export const submitForm = async ({
  one,
  two,
  three,
}: {
  one: string;
  two: string;
  three: string;
}) => {
  const body = JSON.stringify({ one, two, three });

  try {
    const response = await axiosApiInstance.post("/form", body);
    const { data } = response;
    return data;
  } catch (error) {
    return false;
  }
};
