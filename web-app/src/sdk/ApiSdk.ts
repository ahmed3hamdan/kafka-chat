import Axios, { AxiosError, AxiosInstance, AxiosRequestConfig, RawAxiosRequestHeaders } from "axios";

export const ResponseErrorKeys = {
  InvalidRequestBodyErrorKey: "invalid-request-body",
  UsernameRegisteredErrorKey: "username-registered",
  UserNotFoundErrorKey: "user-not-found",
  PasswordMismatchErrorKey: "password-mismatch",
  InvalidAuthTokenErrorKey: "invalid-auth-token",
} as const;

export type ResponseErrorKey = typeof ResponseErrorKeys[keyof typeof ResponseErrorKeys];

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

export class ResponseError<T = ResponseErrorKey> extends Error {
  readonly key: ResponseErrorKey;
  readonly message: string;
  readonly data: T;

  constructor(code: ResponseErrorKey, message: string, data: T) {
    super(message);
    this.key = code;
    this.message = message;
    this.data = data;
  }
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

export type LoginResponseErrors = ResponseError<typeof ResponseErrorKeys.InvalidRequestBodyErrorKey | typeof ResponseErrorKeys.UsernameRegisteredErrorKey>;

export type RegisterResponseErrors = ResponseError<
  typeof ResponseErrorKeys.InvalidRequestBodyErrorKey | typeof ResponseErrorKeys.UserNotFoundErrorKey | typeof ResponseErrorKeys.PasswordMismatchErrorKey
>;

class ApiSdk {
  private authToken: string | null = null;
  private readonly axios: AxiosInstance;

  constructor({ baseURL }: ApiSdkParams) {
    this.axios = Axios.create({
      baseURL,
    });
    this.axios.interceptors.request.use((config: AxiosRequestConfig) => {
      if (this.authToken !== null) {
        config.headers = (config.headers ?? {}) as RawAxiosRequestHeaders;
        config.headers["Authorization"] = `Bearer ${this.authToken}`;
      }
      return config;
    });
  }

  public setAuthorization = (authToken: string | null) => {
    this.authToken = authToken;
  };

  private static axiosErrorHandler = <T extends ResponseError = ResponseError>(error: unknown) => {
    const err = error as AxiosError<T>;

    if (err.response === undefined) {
      throw new ConnectionError();
    }
    if (err.response.status < 500) {
      const { key, message, data } = err.response.data;
      throw new ResponseError(key, message, data);
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
