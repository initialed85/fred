import Input from "@mui/joy/Input";

export interface BranchFilterTextBoxProps {
  setBranchFilter: (branchFilter: string) => void;
}

export default function BranchFilterTextBox(props: BranchFilterTextBoxProps) {
  return (
    <Input
      size="sm"
      sx={{ mr: 1.5, width: "100%", maxWidth: 300 }}
      onChange={(e) => {
        props.setBranchFilter((e.target.value || "").trim().toLowerCase());
      }}
    />
  );
}
