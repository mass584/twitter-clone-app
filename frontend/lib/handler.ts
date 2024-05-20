import axios from "axios";

export const configuredAxios = axios.create({
  baseURL: process.env.API_SERVER_HOST,
  headers: {
    "Content-Type": "application/json",
  },
  withCredentials: true,
});
