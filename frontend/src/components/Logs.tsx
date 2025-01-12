import { LazyLog, ScrollFollow } from "@melloware/react-logviewer";
import { useEffect, useState } from "react";
import { useQuery } from "../api";

export interface LogsProps {
  outputId: string | undefined;
}

export const Logs = (props: LogsProps) => {
  const [isDone, setIsDone] = useState(false);

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

  useEffect(() => {
    const thisIsDone = ["succeeded", "failed", "errored"].includes(
      data?.objects?.[0].status || "",
    );
    if (thisIsDone !== isDone) {
      setIsDone(thisIsDone);
    }
  }, [data?.objects, isDone]);

  const text = data?.objects?.[0].log_id_object?.buffer
    ? atob(data.objects[0].log_id_object.buffer || "")
    : "(no data yet)";

  return (
    <ScrollFollow
      startFollowing={!isDone}
      render={({ follow, onScroll }) => (
        <LazyLog
          text={text}
          enableHotKeys={true}
          enableSearch={true}
          enableSearchNavigation={true}
          selectableLines={true}
          follow={!isDone ? follow : undefined}
          onScroll={!isDone ? onScroll : undefined}
        />
      )}
    />
  );
};
