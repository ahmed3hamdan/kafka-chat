import { useQuery } from "@tanstack/react-query";
import { useAppContext } from "@contexts/appContext";
import { SelfInfoResponse } from "@sdk/ApiSdk";
import { UseQueryOptions } from "@tanstack/react-query/src/types";

type UseGetSelfInfoQueryOptions = Omit<UseQueryOptions<SelfInfoResponse>, "queryKey" | "queryFn" | "initialData"> & { initialData?: () => undefined };

const useGetSelfInfoQuery = (options?: UseGetSelfInfoQueryOptions) => {
  const { api } = useAppContext();
  return useQuery<SelfInfoResponse>(["getSelfInfo"], api.getSelfInfo, options);
};

export default useGetSelfInfoQuery;
