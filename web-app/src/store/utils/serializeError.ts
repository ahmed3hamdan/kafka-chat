import { CommonError } from "../types";
import { ConnectionError, InternalServerError, ResponseError } from "@sdk/ApiSdk";

const serializeError: (error: Error) => CommonError = error => {
  if (error instanceof ResponseError) {
    return {
      key: error.key,
      message: error.message,
    };
  }
  if (error instanceof InternalServerError) {
    return {
      key: "internal-server-error",
      message: error.message,
    };
  }
  if (error instanceof ConnectionError) {
    return {
      key: "connection-error",
      message: error.message,
    };
  }
  return {
    key: "unknown-error",
    message: error.message,
  };
};

export default serializeError;
