import React from "react";
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogProps,
  DialogTitle,
} from "@mui/material";
import ZoomTransition from "@components/transitions/ZoomTransition";

interface LogoutDialogProps extends Omit<DialogProps, "onClose"> {
  onLogout: () => void;
  onClose: () => void;
}

const LogoutDialog: React.FC<LogoutDialogProps> = ({
  onLogout,
  onClose,
  ...rest
}) => {
  return (
    <Dialog
      onClose={onClose}
      maxWidth="xs"
      TransitionComponent={ZoomTransition}
      {...rest}
    >
      <DialogTitle>Logout</DialogTitle>
      <DialogContent>
        <DialogContentText>Are you sure you want to logout?</DialogContentText>
      </DialogContent>
      <DialogActions>
        <Button type="button" onClick={onClose} autoFocus>
          Cancel
        </Button>
        <Button type="button" color="error" onClick={onLogout}>
          Logout
        </Button>
      </DialogActions>
    </Dialog>
  );
};

export default LogoutDialog;
