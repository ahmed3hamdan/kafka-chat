import React, { HTMLAttributes, useState } from "react";
import { Link, Navigate, Outlet, useMatch } from "react-router-dom";
import {
  Avatar,
  Box,
  css,
  Divider,
  Drawer,
  IconButton,
  List,
  ListItem,
  ListItemAvatar,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Menu,
  MenuItem,
  MenuList,
  styled,
  Typography,
} from "@mui/material";
import { MoreVert as MoreVertIcon, Add as AddIcon, Logout as LogoutIcon } from "@mui/icons-material";
import LogoutDialog from "@components/overlay/LogoutDialog";
import NewConversationDialog, { NewConversationValues } from "@components/overlay/NewConversationDialog";
import { Components, Virtuoso } from "react-virtuoso";
import { SingleConversation } from "@sdk/types";
import useAuth from "@hooks/useAuth";

const StyledVirtuoso = styled(Virtuoso)`
  flex-grow: 1;
`;

interface ConversationListItemProps {
  conversation: SingleConversation;
}

const ConversationListItem: React.FC<ConversationListItemProps> = ({ conversation }) => {
  const match = useMatch({ path: "/:userID" });
  return (
    <ListItemButton selected={match !== null && match.params.userID === conversation.userID} component={Link} to={`/${conversation.userID}`} dense>
      <ListItemAvatar>
        <Avatar alt={conversation.name}>{conversation.name.charAt(0)}</Avatar>
      </ListItemAvatar>
      <ListItemText
        primary={<Typography noWrap>{conversation.name}</Typography>}
        secondary={
          <Typography variant="body2" color="text.secondary" noWrap>
            {conversation.lastMessage.content}
          </Typography>
        }
      />
    </ListItemButton>
  );
};

type DialogId = "logout" | "new-conversation";

const virtuosoComponents: Components = {
  List: React.forwardRef<HTMLDivElement, HTMLAttributes<HTMLDivElement>>(({ style, children }, listRef) => {
    return (
      <List style={{ padding: 0, ...style, margin: 0 }} component="div" ref={listRef}>
        {children}
      </List>
    );
  }),
  Item: ({ children, ...props }: HTMLAttributes<HTMLDivElement>) => {
    return (
      <>
        <ListItem component="div" {...props} style={{ margin: 0 }} disablePadding>
          {children}
        </ListItem>
        <Divider />
      </>
    );
  },
};

const layoutCss = css`
  height: 100%;
  padding-left: 320px;
`;

const MainView = () => {
  const { logout } = useAuth();
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const [openDialog, setOpenDialog] = useState<DialogId | null>(null);

  const handleProfileMoreClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleProfileMenuClose = () => {
    setAnchorEl(null);
  };

  const handleProfileNewConversationClick = () => {
    setAnchorEl(null);
    setOpenDialog("new-conversation");
  };

  const handleProfileLogoutClick = () => {
    setAnchorEl(null);
    setOpenDialog("logout");
  };

  const handleDialogClose = () => {
    setOpenDialog(null);
  };

  const handleLogoutConfirm = () => {
    setOpenDialog(null);
    logout();
  };

  const handleNewConversationSubmit = ({ email }: NewConversationValues) => {
    console.log({ email });
    return new Promise<void>(resolve => {
      setTimeout(() => {
        setOpenDialog(null);
        resolve();
      }, 3000);
    });
  };

  return (
    <>
      <Drawer variant="permanent" open>
        <Box
          sx={{
            width: "320px",
            height: "100%",
            display: "flex",
            flexDirection: "column",
          }}
        >
          <ListItem
            secondaryAction={
              <IconButton edge="end" onClick={handleProfileMoreClick}>
                <MoreVertIcon />
              </IconButton>
            }
          >
            <ListItemAvatar>
              <Avatar alt={"selfInfoQuery.data.name"}>{"selfInfoQuery.data.name".charAt(0)}</Avatar>
            </ListItemAvatar>
            <ListItemText primary={<Typography noWrap>{"selfInfoQuery.data.name"}</Typography>} secondary={"selfInfoQuery.data.username"} />
          </ListItem>
          <Divider />
          <StyledVirtuoso
            totalCount={1000}
            fixedItemHeight={65}
            components={virtuosoComponents}
            itemContent={index => (
              <ConversationListItem
                conversation={{
                  userID: `user-${index}`,
                  username: `username-${index}`,
                  name: `Conversation ${index}`,
                  lastMessage: {
                    userID: `user-${index}`,
                    content: "Hello World",
                  },
                }}
              />
            )}
          />
          <List disablePadding></List>
        </Box>
      </Drawer>
      <Menu
        anchorEl={anchorEl}
        open={Boolean(anchorEl)}
        onClose={handleProfileMenuClose}
        anchorOrigin={{ horizontal: "right", vertical: "bottom" }}
        transformOrigin={{ horizontal: "right", vertical: "top" }}
      >
        <MenuList disablePadding>
          <MenuItem onClick={handleProfileNewConversationClick}>
            <ListItemIcon>
              <AddIcon fontSize="small" />
            </ListItemIcon>
            <ListItemText>New conversation</ListItemText>
          </MenuItem>
          <MenuItem onClick={handleProfileLogoutClick}>
            <ListItemIcon>
              <LogoutIcon fontSize="small" />
            </ListItemIcon>
            <ListItemText>Logout</ListItemText>
          </MenuItem>
        </MenuList>
      </Menu>
      <LogoutDialog open={openDialog == "logout"} onClose={handleDialogClose} onLogout={handleLogoutConfirm} />
      <NewConversationDialog open={openDialog == "new-conversation"} onClose={handleDialogClose} onSubmit={handleNewConversationSubmit} />
      <div css={layoutCss}>
        <Outlet />
      </div>
    </>
  );
};

const MainViewWrapper = () => {
  const { loggedIn } = useAuth();
  if (!loggedIn) {
    return <Navigate to="/login" replace />;
  }
  return <MainView />;
};

export default MainViewWrapper;
