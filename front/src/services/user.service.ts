import axios from "axios";
import authHeader from "./auth-header";
import {gamesPlayedByPlayerByQuery} from "./graphql-query/games-query";
import makeGraphQLCall from "./requestGraphql";
import {loginUserByQuery} from "./graphql-query/users-query";
import IUser from "../types/user.type";

const API_URL = "https://random-word-api.herokuapp.com/";

export const getPublicContent = () => {
  return axios.get(API_URL + "word");
};

export const getUserBoard = () => {
  makeGraphQLCall(gamesPlayedByPlayerByQuery())
      .then((response) => {
        if (response.status === 200) {
          localStorage.setItem("games", response.data.data.games);
        } else {
          throw new Error(response.statusText);
        }
      });
  return axios.get(API_URL + "word");
};

export const getModeratorBoard = () => {
  return axios.get(API_URL + "word", { headers: authHeader() });
};

export const getAdminBoard = () => {
  return axios.get(API_URL + "word", { headers: authHeader() });
};
