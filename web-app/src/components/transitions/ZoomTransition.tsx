import React from "react";
import { TransitionProps } from "@mui/material/transitions";
import { Zoom } from "@mui/material";

const ZoomTransition = React.forwardRef(
  (
    props: TransitionProps & {
      children: React.ReactElement<any, any>;
    },
    ref: React.Ref<unknown>
  ) => <Zoom ref={ref} {...props} />
);

export default ZoomTransition;
