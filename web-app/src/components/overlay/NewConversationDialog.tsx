import React, { useRef } from "react";
import * as Yup from "yup";
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogProps,
  DialogTitle,
  TextField,
} from "@mui/material";
import ZoomTransition from "@components/transitions/ZoomTransition";
import { Field, FieldProps, Form, Formik, FormikProps } from "formik";
import { LoadingButton } from "@mui/lab";

export interface NewConversationValues {
  email: string;
}

const initialValues: NewConversationValues = {
  email: "",
};

const validationSchema = Yup.object().shape({
  email: Yup.string().trim().required().email(),
});

interface NewConversationDialogProps
  extends Omit<DialogProps, "onSubmit" | "onClose"> {
  onSubmit: (values: NewConversationValues) => Promise<void>;
  onClose: () => void;
}

const NewConversationDialog: React.FC<NewConversationDialogProps> = ({
  onSubmit,
  onClose,
  ...rest
}) => {
  const formikRef = useRef<FormikProps<NewConversationValues>>(null);
  const handleClose = () => {
    if (formikRef.current == null || formikRef.current?.isSubmitting) {
      return;
    }
    return onClose();
  };
  return (
    <Dialog
      onClose={handleClose}
      maxWidth="lg"
      TransitionComponent={ZoomTransition}
      {...rest}
    >
      <Formik
        initialValues={initialValues}
        validationSchema={validationSchema}
        onSubmit={onSubmit}
        innerRef={formikRef}
      >
        {({ isSubmitting }: FormikProps<NewConversationValues>) => (
          <Form noValidate>
            <DialogTitle>New Conversation</DialogTitle>
            <DialogContent>
              <DialogContentText>
                Please enter user email to start a new conversation
              </DialogContentText>
              <Field name="email">
                {({ field, meta: { error, touched } }: FieldProps) => (
                  <TextField
                    {...field}
                    error={touched && Boolean(error)}
                    helperText={touched && error}
                    label="Email"
                    margin="dense"
                    variant="standard"
                    autoFocus
                    required
                    fullWidth
                    disabled={isSubmitting}
                  />
                )}
              </Field>
            </DialogContent>
            <DialogActions>
              <Button
                type="button"
                onClick={handleClose}
                disabled={isSubmitting}
              >
                Cancel
              </Button>
              <LoadingButton loading={isSubmitting} type="submit">
                Create
              </LoadingButton>
            </DialogActions>
          </Form>
        )}
      </Formik>
    </Dialog>
  );
};

export default NewConversationDialog;
