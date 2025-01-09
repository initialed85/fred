import Table from "@mui/joy/Table";
import { useQuery } from "../api";
import { components } from "../api/api";
import { Output } from "./Output";
import { Status } from "./Status";

export interface ExecutionProps {
  executionId: string;
  setOutputId: (outputId: string) => void;
  setShowLogModal: (showModal: boolean) => void;
}

export function Execution(props: ExecutionProps) {
  const { data: dataForExecutions } = useQuery("get", "/api/executions/{primaryKey}", {
    params: {
      path: {
        primaryKey: props.executionId || "",
      },
      query: {
        depth: 2,
      },
    },
  });

  if (!dataForExecutions?.objects?.length) {
    return null;
  }

  const execution: components["schemas"]["Execution"] = dataForExecutions!.objects[0];

  return (
    <tr key={execution?.id}>
      <td style={{ width: "200px" }}>{execution?.job_id_object?.name}</td>
      <td style={{ width: "100px" }}>
        <Status status={execution?.status!} />
      </td>
      <td>
        <Table size="sm" sx={{ p: 0, m: 0 }} borderAxis="y">
          <tbody>
            {execution?.referenced_by_output_execution_id_objects
              ?.filter((output) => output.execution_id === execution?.id)
              .sort((a, b) => {
                if (a!.task_id_object?.index! < b!.task_id_object?.index!) {
                  return -1;
                } else if (a!.task_id_object?.index! > b!.task_id_object?.index!) {
                  return 1;
                } else {
                  return 0;
                }
              })
              .map(
                (output) =>
                  output?.id && (
                    <Output
                      outputId={output.id}
                      setOutputId={props.setOutputId}
                      setShowLogModal={props.setShowLogModal}
                      key={output.id}
                    />
                  )
              )}
          </tbody>
        </Table>
      </td>
    </tr>
  );
}
