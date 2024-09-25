import TaskIcon from "@mui/icons-material/Task";
import Dropdown from "@mui/joy/Dropdown";
import Menu from "@mui/joy/Menu";
import MenuButton from "@mui/joy/MenuButton";
import MenuItem from "@mui/joy/MenuItem";
import { Dispatch, SetStateAction } from "react";
import { useQuery } from "../api";

export interface TaskDropdownMenuProps {
  responsive: boolean;
  taskId: string | undefined;
  setTaskId: Dispatch<SetStateAction<string | undefined>>;
  repositoryId: string | undefined;
  ruleId: string | undefined;
  jobId: string | undefined;
}

export default function TaskDropdownMenu(props: TaskDropdownMenuProps) {
  const { data } = useQuery("get", "/api/tasks", {
    params: {
      query: {
        name__asc: "",
      },
    },
  });

  return (
    <Dropdown>
      <MenuButton
        variant="soft"
        color="danger"
        size={"sm"}
        sx={{ marginLeft: 0.5, marginRight: 1, width: 33 }}
      >
        <TaskIcon />
      </MenuButton>
      <Menu>
        {data?.objects?.map((task) => {
          return (
            <MenuItem
              key={`task-dropdown-menu-item-${task.id}`}
              selected={props.taskId === task.id}
              onClick={(event) => {
                props.setTaskId(props.taskId !== task.id ? task.id : undefined);
              }}
            >
              {task.name}
            </MenuItem>
          );
        })}
      </Menu>
    </Dropdown>
  );
}
