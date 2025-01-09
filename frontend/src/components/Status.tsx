import Chip from "@mui/joy/Chip";

export function Status(props: { status: string }) {
  const statusText = (props.status || "").slice(0, 1).toUpperCase() + (props.status || "").slice(1);

  let status = <Chip color={"danger"}>{statusText}</Chip>;

  if (props.status === "running") {
    status = <Chip color={"warning"}>{statusText}</Chip>;
  } else if (props.status === "pending") {
    status = <Chip color={"primary"}>{statusText}</Chip>;
  } else if (props.status === "succeeded") {
    status = <Chip color={"success"}>{statusText}</Chip>;
  } else if (props.status === "skipped") {
    status = <Chip color={"warning"}>{statusText}</Chip>;
  }

  return status;
}
