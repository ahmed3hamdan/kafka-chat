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

const registerSchema = z.object({
  name: z.string().trim().max(60),
  username: z.string().max(20).regex(/^[a-z][a-z0-9_\-]*$/),
  password: z.string().max(72),
});

type RegisterValues = z.infer<typeof registerSchema>;

const initialValues: RegisterValues = {
  name: "",
  username: "",
  password: "",
};

const RegisterView = () => {
  const { api, setAuth } = useAppContext();
  const { mutateAsync, isError, error } = useMutation(
    ({ name, username, password }: RegisterValues) => api.register({ name, username, password }),
    {
      onSuccess: ({ userID, token }) => setAuth({ userID, token }),
    }
  );

  const handleSubmit = useCallback(
    (values: RegisterValues) => mutateAsync(values),
    [mutateAsync]
  );

  const responseError = useMemo(() => {
    if (error instanceof ResponseError) {
      switch (error.code) {
        case ResponseErrorCode.InvalidRequestBodyErrorCode:
          return "Invalid inputs";
        case ResponseErrorCode.UsernameRegisteredErrorCode:
          return "Username already registered";
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
        Register
      </Typography>
      <Formik initialValues={initialValues} onSubmit={handleSubmit} validationSchema={toFormikValidationSchema(registerSchema)}>
        {({ isSubmitting, values, touched, errors, handleChange }) => (
          <Form noValidate>
            <TextField
              margin="normal"
              required
              fullWidth
              label="Name"
              name="name"
              autoComplete="name"
              value={values.name}
              onChange={handleChange}
              error={touched.name && Boolean(errors.name)}
              helperText={touched.name && errors.name}
              autoFocus
            />
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
              Register
            </LoadingButton>
          </Form>
        )}
      </Formik>
      <Stack direction="row" sx={{ mt: 2 }}>
        <MuiLink component={Link} to="/login">
          Login
        </MuiLink>
      </Stack>
    </Container>
  );
};

export default RegisterView;
