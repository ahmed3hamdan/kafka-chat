import { Container, TextField, Typography, Link as MuiLink, Stack, Alert } from "@mui/material";
import { z } from "zod";
import { Form, Formik } from "formik";
import { useCallback, useMemo, useState } from "react";
import { LoadingButton } from "@mui/lab";
import { Link } from "react-router-dom";
import { ResponseErrorKeys } from "@sdk/ApiSdk";
import { toFormikValidationSchema } from "zod-formik-adapter";
import useAuth from "@hooks/useAuth";
import { CommonError } from "@store/types";

const loginSchema = z.object({
  username: z
    .string()
    .max(20)
    .regex(/^[a-z][a-z0-9_\-]*$/),
  password: z.string().max(72),
});

type LoginValues = z.infer<typeof loginSchema>;

const initialValues: LoginValues = {
  username: "",
  password: "",
};

const LoginView = () => {
  const [error, setError] = useState<null | CommonError>(null);
  const { login } = useAuth();

  const handleSubmit = useCallback(
    (values: LoginValues) => {
      setError(null);
      return login(values).catch(err => setError(err));
    },
    [login]
  );

  const responseError = useMemo(() => {
    if (error === null) return null;
    switch (error.key) {
      case ResponseErrorKeys.InvalidRequestBodyErrorKey:
        return "Invalid inputs";
      case ResponseErrorKeys.UserNotFoundErrorKey:
        return "User not registered";
      case ResponseErrorKeys.PasswordMismatchErrorKey:
        return "Incorrect password";
      case "internal-server-error":
        return "Internal server error";
      case "connection-error":
        return "Failed to connect to server";
      default:
        return "Unknown error";
    }
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
            {responseError && (
              <Alert severity="error" color="error">
                {responseError}
              </Alert>
            )}
            <LoadingButton loading={isSubmitting} type="submit" fullWidth variant="contained" sx={{ mt: 2 }}>
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
