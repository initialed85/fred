import { LazyLog, ScrollFollow } from "@melloware/react-logviewer";
import { useQuery } from "../api";

export interface LogsProps {
  outputId: string | undefined;
}

export const Logs = (props: LogsProps) => {
  const { data } = useQuery("get", "/api/outputs/{primaryKey}", {
    params: {
      path: {
        primaryKey: props.outputId || "",
      },
      query: {
        depth: 2,
      },
    },
  });

  const text = data?.objects?.[0].log_id_object?.buffer ? atob(data.objects[0].log_id_object.buffer || "") : "(no data yet)";

  return (
    <ScrollFollow
      startFollowing={true}
      render={({ follow, onScroll }) => (
        <LazyLog text={text} selectableLines={true} follow={follow} onScroll={onScroll} />
      )}
    />
  );
};
