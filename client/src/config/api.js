import axios from "axios";

// Create base URL API
const API = axios.create({
  baseURL: "https://hallo-corona-production-16c1.up.railway.app/api/v1",
  // baseURL: process.env.REACT_APP_BASE_URL,
});

// Set Authorization Token Header
const setAuthToken = (token) => {
  if (token) {
    API.defaults.headers.common["Authorization"] = `Bearer ${token}`;
  } else {
    delete API.defaults.headers.common["Authorization"];
  }
};

export {API, setAuthToken}