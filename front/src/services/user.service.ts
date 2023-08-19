import axios from "axios";
import authHeader from "./auth-header";

const API_URL = "https://random-word-api.herokuapp.com/";

export const getPublicContent = () => {
  return axios.get(API_URL + "word");
};

export const getUserBoard = () => {
  return axios.get(API_URL + "word", { headers: authHeader() });
};

export const getModeratorBoard = () => {
  return axios.get(API_URL + "word", { headers: authHeader() });
};

export const getAdminBoard = () => {
  return axios.get(API_URL + "word", { headers: authHeader() });
};
