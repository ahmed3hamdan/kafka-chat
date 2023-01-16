import {
  Container,
  TextField,
  Typography,
  Link as MuiLink,
  Stack,
  Alert,
} from "@mui/material";
import { z } from "zod";
import { Form, Formik } from "formik";
import { useCallback, useMemo } from "react";
import { useAppContext } from "@contexts/appContext";
import { useMutation } from "@tanstack/react-query";
import { LoadingButton } from "@mui/lab";
import { Link } from "react-router-dom";
import {
  ConnectionError,
  InternalServerError,
  ResponseError,
  ResponseErrorCode,
} from "@sdk/ApiSdk";
import { toFormikValidationSchema } from "zod-formik-adapter";

const loginSchema = z.object({
  username: z.string().max(20).regex(/^[a-z][a-z0-9_\-]*$/),
  password: z.string().max(72),
});

type LoginValues = z.infer<typeof loginSchema>;

const initialValues: LoginValues = {
  username: "",
  password: "",
};

const LoginView = () => {
  const { api, setAuth } = useAppContext();
  const { mutateAsync, isError, error } = useMutation(
    ({ username, password }: LoginValues) => api.login({ username, password }),
    {
      onSuccess: ({ userID, token }) => setAuth({ userID, token }),
    }
  );

  const handleSubmit = useCallback(
    (values: LoginValues) => mutateAsync(values),
    [mutateAsync]
  );

  const responseError = useMemo(() => {
    if (error instanceof ResponseError) {
      switch (error.code) {
        case ResponseErrorCode.InvalidRequestBodyErrorCode:
          return "Invalid inputs";
        case ResponseErrorCode.NotFoundErrorCode:
          return "User not registered";
        case ResponseErrorCode.PasswordMismatchErrorCode:
          return "Incorrect password";
      }
    }
    if (error instanceof InternalServerError) {
      return "Internal server error";
    }
    if (error instanceof ConnectionError) {
      return "Failed to connect to server";
    }
    return "Unknown error";
  }, [error]);

  return (
    <Container component="main" maxWidth="xs" sx={{ py: 5 }}>
      <Typography component="h1" variant="h5" textAlign="center">
        Login
      </Typography>
      <Formik initialValues={initialValues} onSubmit={handleSubmit} validationSchema={toFormikValidationSchema(loginSchema)}>
        {({ isSubmitting, values, touched, errors, handleChange }) => (
          <Form noValidate>
            <TextField
              margin="normal"
              required
              fullWidth
              label="Username"
              name="username"
              autoComplete="username"
              value={values.username}
              onChange={handleChange}
              error={touched.username && Boolean(errors.username)}
              helperText={touched.username && errors.username}
              autoFocus
            />
            <TextField
              margin="normal"
              required
              fullWidth
              label="Password"
              name="password"
              type="password"
              autoComplete="current-password"
              value={values.password}
              onChange={handleChange}
              error={touched.password && Boolean(errors.password)}
              helperText={touched.password && errors.password}
            />
            {isError && (
              <Alert severity="error" color="error">
                {responseError}
              </Alert>
            )}
            <LoadingButton
              loading={isSubmitting}
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 2 }}
            >
              Log in
            </LoadingButton>
          </Form>
        )}
      </Formik>
      <Stack direction="row" justifyContent="flex-end" sx={{ mt: 2 }}>
        <MuiLink component={Link} to="/register">
          Register
        </MuiLink>
      </Stack>
    </Container>
  );
};

export default LoginView;
