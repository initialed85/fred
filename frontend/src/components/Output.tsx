import Button from "@mui/joy/Button";
import { useQuery } from "../api";
import { components } from "../api/api";
import { Status } from "./Status";

export interface OutputProps {
  outputId: string;
  setOutputId: (outputId: string) => void;
  setShowLogModal: (showModal: boolean) => void;
}

export function Output(props: OutputProps) {
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

  if (!data?.objects?.length) {
    return null;
  }

  const output: components["schemas"]["Output"] = data.objects[0];

  return (
    <tr>
      <td style={{ width: "25px" }}>{output!.task_id_object?.name}</td>
      <td style={{ width: "50px" }}>
        <Status status={output!.status!} />
      </td>
      <td style={{ width: "100px" }}>
        <Button
          size={"sm"}
          variant="soft"
          color={"primary"}
          sx={{
            fontSize: "var(--joy-fontSize-xs, 0.75rem)",
          }}
          onClick={() => {
            if (!output?.id) {
              return;
            }

            props.setOutputId(output.id);
            props.setShowLogModal(true);
          }}
        >
          Logs
        </Button>
      </td>
    </tr>
  );
}
