import AssignmentIcon from "@mui/icons-material/Assignment";
import Dropdown from "@mui/joy/Dropdown";
import Menu from "@mui/joy/Menu";
import MenuButton from "@mui/joy/MenuButton";
import MenuItem from "@mui/joy/MenuItem";
import { Dispatch, SetStateAction } from "react";
import { useQuery } from "../api";

export interface JobDropdownMenuProps {
  responsive: boolean;
  jobId: string | undefined;
  setJobId: Dispatch<SetStateAction<string | undefined>>;
  repositoryId: string | undefined;
  ruleId: string | undefined;
}

export default function JobDropdownMenu(props: JobDropdownMenuProps) {
  const { data } = useQuery("get", "/api/jobs", {
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
        <AssignmentIcon />
      </MenuButton>
      <Menu>
        {data?.objects?.map((job) => {
          return (
            <MenuItem
              key={`job-dropdown-menu-item-${job.id}`}
              selected={props.jobId === job.id}
              onClick={(event) => {
                props.setJobId(props.jobId !== job.id ? job.id : undefined);
              }}
            >
              {job.name}
            </MenuItem>
          );
        })}
      </Menu>
    </Dropdown>
  );
}
