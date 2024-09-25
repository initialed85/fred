import { LazyLog, ScrollFollow } from "@melloware/react-logviewer";
import { useQuery } from "../api";

export interface LogsProps {
  outputId: string | undefined;
}

export const Logs = (props: LogsProps) => {
  const { isLoading, error, data } = useQuery(
    "get",
    "/api/outputs/{primaryKey}",
    {
      params: {
        path: {
          primaryKey: props.outputId || "",
        },
      },
    },
  );

  let text;

  if (isLoading) {
    text = "loading...";
  } else if (error) {
    text = `Error ${error.status}: ${error.error.join("; ")}`;
  } else {
    text = (data?.objects?.[0] as any)?.buffer;
  }

  return (
    <ScrollFollow
      startFollowing={true}
      render={({ follow, onScroll }) => (
        <LazyLog
          text={text}
          selectableLines={true}
          follow={follow}
          onScroll={onScroll}
        />
      )}
    />
  );
};
