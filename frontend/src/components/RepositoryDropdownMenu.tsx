import InventoryIcon from "@mui/icons-material/Inventory";
import Dropdown from "@mui/joy/Dropdown";
import Menu from "@mui/joy/Menu";
import MenuButton from "@mui/joy/MenuButton";
import MenuItem from "@mui/joy/MenuItem";
import { Dispatch, SetStateAction } from "react";
import { useQuery } from "../api";

export interface RepositoryDropdownMenuProps {
  responsive: boolean;
  repositoryId: string | undefined;
  setRepositoryId: Dispatch<SetStateAction<string | undefined>>;
}

export default function RepositoryDropdownMenu(
  props: RepositoryDropdownMenuProps,
) {
  const { data } = useQuery("get", "/api/repositories", {
    params: {
      query: {
        url__asc: "",
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
        <InventoryIcon />
      </MenuButton>
      <Menu>
        {data?.objects?.map((repository) => {
          return (
            <MenuItem
              key={`repository-dropdown-menu-item-${repository.id}`}
              selected={props.repositoryId === repository.id}
              onClick={(event) => {
                props.setRepositoryId(
                  props.repositoryId !== repository.id
                    ? repository.id
                    : undefined,
                );
              }}
            >
              {repository.url}
            </MenuItem>
          );
        })}
      </Menu>
    </Dropdown>
  );
}
