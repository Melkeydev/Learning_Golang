import React, { useState, Fragment, useCallback } from "react";
import axios from "axios";
import { useForm } from "react-hook-form";
import { Input } from "./common/input";
import { submitForm } from "../actions/formSubmit";

export const TestForm = () => {
  const [myForm, setMyForm] = useState<any>({});
  const { register, handleSubmit, formState, reset } = useForm();

  const onSubmit = async (data: any) => {
    const response = await submitForm(data);
  };

  const handleChange = useCallback((e) => {
    const { id, value } = e.target;

    setMyForm((state: any) => ({
      ...state,
      [id]: value,
    }));
  }, []);

  // TODO: Make forms components re usable
  return (
    <div className="w-full max-w-xs mx-auto">
      <form
        className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4"
        onSubmit={handleSubmit(onSubmit)}
      >
        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2">
            Test Field One
          </label>
          <input
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            id="first field"
            type="text"
            placeholder="One"
            value={myForm.One}
            {...register("one")}
            onChange={handleChange}
          />
        </div>
        <div className="mb-6">
          <label className="block text-gray-700 text-sm font-bold mb-2">
            Test Field Two
          </label>
          <input
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            id="two"
            type="text"
            placeholder="Two"
            value={myForm.Two}
            {...register("two")}
            onChange={handleChange}
          />
        </div>
        <div className="mb-6">
          <label className="block text-gray-700 text-sm font-bold mb-2">
            Test Field Three
          </label>
          <input
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            id="three"
            type="text"
            placeholder="three"
            value={myForm.Three}
            {...register("three")}
            onChange={handleChange}
          />
        </div>
        <div className="flex items-center justify-center">
          <button
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
            type="submit"
          >
            Sign In
          </button>
        </div>
      </form>
    </div>
  );
};
