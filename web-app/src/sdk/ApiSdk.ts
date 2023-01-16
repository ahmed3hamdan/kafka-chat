import Axios, { AxiosError, AxiosInstance, RawAxiosRequestHeaders } from "axios";

export const ResponseErrorCode = {
  InvalidRequestBodyErrorCode: 1001,
  UsernameRegisteredErrorCode: 1002,
  InvalidParamsErrorCode: 1003,
  NotFoundErrorCode: 1004,
  PasswordMismatchErrorCode: 1005,
  InvalidAuthTokenErrorCode: 1006,
} as const;

export type ResponseErrorCodeType = typeof ResponseErrorCode[keyof typeof ResponseErrorCode];

export class ConnectionError extends Error {
  constructor() {
    super("connection error");
  }
}

export class InternalServerError extends Error {
  constructor() {
    super("internal server error");
  }
}

export class ResponseError<T = null> extends Error {
  readonly code: ResponseErrorCodeType;
  readonly message: string;
  readonly data: T;

  constructor(code: ResponseErrorCodeType, message: string, data: T) {
    super(message);
    this.code = code;
    this.message = message;
    this.data = data;
  }
}

interface InvalidRequestBodyResponseError extends ResponseError {
  code: typeof ResponseErrorCode.InvalidAuthTokenErrorCode;
}

interface UsernameRegisteredResponseError extends ResponseError {
  code: typeof ResponseErrorCode.UsernameRegisteredErrorCode;
}

export interface ApiSdkParams {
  authToken?: string;
  baseURL: string;
}

export interface LoginParams {
  username: string;
  password: string;
}

export interface RegisterParams {
  name: string;
  username: string;
  password: string;
}

export interface AuthResponse {
  userID: number;
  token: string;
}

export interface SelfInfoResponse {
  userID: number;
  name: string;
  username: string;
}

export type LoginResponseErrors = InvalidRequestBodyResponseError;

export type RegisterResponseErrors = InvalidRequestBodyResponseError | UsernameRegisteredResponseError;

class ApiSdk {
  private readonly axios: AxiosInstance;

  constructor({ authToken, baseURL }: ApiSdkParams) {
    const headers: RawAxiosRequestHeaders = {};
    if (authToken !== undefined) {
      headers["Authorization"] = `Bearer ${authToken}`;
    }
    this.axios = Axios.create({
      baseURL,
      headers,
    });
  }

  private static axiosErrorHandler = <T extends ResponseError = ResponseError>(error: unknown) => {
    const err = error as AxiosError<T>;

    if (err.response === undefined) {
      throw new ConnectionError();
    }
    if (err.response.status < 500) {
      const { code, message, data } = err.response.data;
      throw new ResponseError(code, message, data);
    }
    throw new InternalServerError();
  };

  public register = (data: RegisterParams) =>
    this.axios
      .post<AuthResponse>("/api/auth/register", data)
      .then(({ data }) => data)
      .catch(ApiSdk.axiosErrorHandler<RegisterResponseErrors>);

  public login = (data: LoginParams) =>
    this.axios
      .post<AuthResponse>("/api/auth/login", data)
      .then(({ data }) => data)
      .catch(ApiSdk.axiosErrorHandler<LoginResponseErrors>);

  public getSelfInfo = () =>
    this.axios
      .post<SelfInfoResponse>("/api/auth/get-self-info")
      .then(({ data }) => data)
      .catch(ApiSdk.axiosErrorHandler);
}

export default ApiSdk;
