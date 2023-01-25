import { useCallback } from "react";
import { shallowEqual, useDispatch, useSelector } from "react-redux";
import { AppDispatch, RootState } from "@store/index";
import { clearSelfInfo, getSelfInfo } from "@store/slices/selfInfo";

const useSelfInfo = () => {
  const dispatch: AppDispatch = useDispatch();
  const { status, lastFetchedAt, data } = useSelector((state: RootState) => state.selfInfo, shallowEqual);
  const get = useCallback(() => dispatch(getSelfInfo()).unwrap(), [dispatch]);
  const clear = useCallback(() => dispatch(clearSelfInfo()), [dispatch]);
  return { status, lastFetchedAt, data, get, clear };
};

export default useSelfInfo;
