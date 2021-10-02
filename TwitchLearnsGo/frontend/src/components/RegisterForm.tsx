import React, { useState, Fragment, useCallback } from "react";
import axios from "axios";
import { useForm } from "react-hook-form";
import { Input } from "./common/input";
import { submitForm } from "../actions/formSubmit";
import { registerForm } from "../actions/authActions";

export const RegisterForm = () => {
  interface IregisterForm {
    username?: string;
    password?: string;
  }
  const [registerForm, setRegisterForm] = useState<IregisterForm>({});
  const { register, handleSubmit, formState, reset } = useForm();

  const onSubmit = (data: IregisterForm) => {
    console.log(data);
  };

  const handleChange = useCallback((e) => {
    const { id, value } = e.target;

    setRegisterForm((state: any) => ({
      ...state,
      [id]: value,
    }));
  }, []);

  return (
    <div className="w-full max-w-xs">
      <form
        className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4"
        onSubmit={handleSubmit(onSubmit)}
      >
        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2">
            Create a Username
          </label>
          <input
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            id="username"
            type="text"
            value={registerForm.username}
            placeholder="Username"
            {...register("username")}
            onChange={handleChange}
          />
        </div>
        <div className="mb-6">
          <label className="block text-gray-700 text-sm font-bold mb-2">
            Choose a password
          </label>
          <input
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            id="password"
            type="password"
            {...register("password")}
            onChange={handleChange}
          />
        </div>
        <div className="flex items-center justify-between">
          <button
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
            type="submit"
          >
            Register
          </button>
        </div>
      </form>
    </div>
  );
};
