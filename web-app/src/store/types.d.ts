import { ResponseErrorKey } from "@sdk/ApiSdk";

export interface CommonError {
  key: "unknown-error" | "connection-error" | "internal-server-error" | ResponseErrorKey;
  message: string;
}
