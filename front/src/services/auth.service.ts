import axios from "axios";
import makeGraphQLCall from "./requestGraphql";
import {loginUserByQuery} from "./graphql-query/users-query";
import IUser from "../types/user.type";

const API_URL = "http://localhost:8080/api/auth/";

export const register = (username: string, email: string, password: string) => {
  return axios.post(API_URL + "signup", {
    username,
    email,
    password,
  });
};

export const login = (email: string, password: string) => {
    const dataload: IUser = {
        accessToken: "",
        email: email,
        password: password,
    };


  return makeGraphQLCall(loginUserByQuery(dataload))
    .then((response) => {
      if (response.data.data.loginUser.token) {

          const loggedinUser: IUser = {
              accessToken: response.data.data.loginUser.token,
              email: response.data.data.loginUser.userdata.email,
              password: "",
              _key: response.data.data.loginUser.userdata._key ,
              _id:  response.data.data.loginUser.userdata._id,
              rev:  response.data.data.loginUser.userdata.rev,
              firstname:  response.data.data.loginUser.userdata.firstname ,
          };

        localStorage.setItem("user", JSON.stringify(loggedinUser));
      }

      return response.data;
    });
};

export const logout = () => {
  localStorage.removeItem("user");
};

export const getCurrentUser = () => {
  const userStr = localStorage.getItem("user");
  if (userStr) return JSON.parse(userStr);

  return null;
};
