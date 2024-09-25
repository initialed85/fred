import AltRouteIcon from "@mui/icons-material/AltRoute";
import Dropdown from "@mui/joy/Dropdown";
import Menu from "@mui/joy/Menu";
import MenuButton from "@mui/joy/MenuButton";
import MenuItem from "@mui/joy/MenuItem";
import { Dispatch, SetStateAction } from "react";
import { useQuery } from "../api";

export interface RuleDropdownMenuProps {
  responsive: boolean;
  ruleId: string | undefined;
  setRuleId: Dispatch<SetStateAction<string | undefined>>;
  repositoryId: string | undefined;
}

export default function RuleDropdownMenu(props: RuleDropdownMenuProps) {
  const { data } = useQuery("get", "/api/rules", {
    params: {
      query: {
        branch_name__asc: "",
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
        <AltRouteIcon />
      </MenuButton>
      <Menu>
        {data?.objects?.map((rule) => {
          return (
            <MenuItem
              key={`rule-dropdown-menu-item-${rule.id}`}
              selected={props.ruleId === rule.id}
              onClick={(event) => {
                props.setRuleId(props.ruleId !== rule.id ? rule.id : undefined);
              }}
            >
              {rule.branch_name}
            </MenuItem>
          );
        })}
      </Menu>
    </Dropdown>
  );
}
