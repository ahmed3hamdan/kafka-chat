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

const registerSchema = z.object({
  name: z.string().trim().max(60),
  username: z
    .string()
    .max(20)
    .regex(/^[a-z][a-z0-9_\-]*$/),
  password: z.string().max(72),
});

type RegisterValues = z.infer<typeof registerSchema>;

const initialValues: RegisterValues = {
  name: "",
  username: "",
  password: "",
};

const RegisterView = () => {
  const [error, setError] = useState<null | CommonError>(null);
  const { register } = useAuth();

  const handleSubmit = useCallback(
    (values: RegisterValues) => {
      setError(null);
      return register(values).catch(err => setError(err));
    },
    [register]
  );

  const responseError = useMemo(() => {
    if (error === null) return null;
    switch (error.key) {
      case ResponseErrorKeys.InvalidRequestBodyErrorKey:
        return "Invalid inputs";
      case ResponseErrorKeys.UsernameRegisteredErrorKey:
        return "Username already registered";
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
            {responseError && (
              <Alert severity="error" color="error">
                {responseError}
              </Alert>
            )}
            <LoadingButton loading={isSubmitting} type="submit" fullWidth variant="contained" sx={{ mt: 2 }}>
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
