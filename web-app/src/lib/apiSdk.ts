import ApiSdk from "@sdk/ApiSdk";
import { API_BASE_URL } from "@config/index";

const apiSdk = new ApiSdk({ baseURL: API_BASE_URL });

export default apiSdk;
