import Grid from "@mui/joy/Grid";
import Sheet from "@mui/joy/Sheet";
import Typography from "@mui/joy/Typography";
import { useEffect, useState } from "react";
import JobDropdownMenu from "./components/JobDropdownMenu";
import ModeToggle from "./components/ModeToggle";
import RepositoryDropdownMenu from "./components/RepositoryDropdownMenu";
import RuleDropdownMenu from "./components/RuleDropdownMenu";
import TaskDropdownMenu from "./components/TaskDropdownMenu";
import { responsiveWidth } from "./config";
import { BuildTable } from "./components/BuildTable";

function App() {
  const [responsive, setResponsive] = useState(
    window.innerWidth < responsiveWidth,
  );
  const [portrait, setPortrait] = useState(
    window.innerWidth < window.innerHeight,
  );
  const [windowWidth, setWindowWidth] = useState(window.innerWidth);
  const [windowHeight, setWindowHeight] = useState(window.innerHeight);

  const [repositoryId, setRepositoryId] = useState<string | undefined>(
    undefined,
  );
  const [ruleId, setRuleId] = useState<string | undefined>(undefined);
  const [jobId, setJobId] = useState<string | undefined>(undefined);
  const [taskId, setTaskId] = useState<string | undefined>(undefined);

  useEffect(() => {
    const handleResize = () => {
      const desiredResponsive = window.innerWidth < responsiveWidth;
      if (desiredResponsive !== responsive) {
        setResponsive(window.innerWidth < responsiveWidth);
      }

      const desiredPortrait = window.innerWidth < window.innerHeight;
      if (desiredPortrait !== portrait) {
        setPortrait(desiredPortrait);
      }

      const desiredWindowWidth = window.innerWidth;
      if (desiredWindowWidth !== windowWidth) {
        setWindowWidth(desiredWindowWidth);
      }

      const desiredWindowHeight = window.innerHeight;
      if (desiredWindowHeight !== windowHeight) {
        setWindowHeight(desiredWindowHeight);
      }
    };

    const eventListener = () => {
      handleResize();
    };

    window.addEventListener("resize", eventListener);

    return () => {
      window.removeEventListener("reisze", eventListener);
    };
  }, [portrait, responsive, windowHeight, windowWidth]);

  return (
    <Sheet
      variant="soft"
      sx={{
        mx: 0,
        my: 0,
        py: 1,
        px: 1,
        display: "flex",
        flexDirection: "column",
        borderRadius: "none",
        boxShadow: "md",
        height: "100%",
      }}
    >
      <Grid container sx={{ pb: 1 }}>
        <Grid
          xs={11}
          sx={{
            display: "flex",
            alignItems: "center",
            justifyContent: "flex-start",
          }}
        >
          <Typography
            level="h4"
            component="h4"
            sx={{ pt: 0.1, pl: 0.75, pr: 1, textAlign: "center" }}
            color="neutral"
          >
            {responsive ? "F" : "Fred"}
          </Typography>
          <RepositoryDropdownMenu
            responsive={responsive}
            repositoryId={repositoryId}
            setRepositoryId={setRepositoryId}
          />
          <RuleDropdownMenu
            responsive={responsive}
            ruleId={ruleId}
            setRuleId={setRuleId}
            repositoryId={repositoryId}
          />
          <JobDropdownMenu
            responsive={responsive}
            jobId={jobId}
            setJobId={setJobId}
            repositoryId={repositoryId}
            ruleId={ruleId}
          />
          <TaskDropdownMenu
            responsive={responsive}
            taskId={taskId}
            setTaskId={setTaskId}
            repositoryId={repositoryId}
            ruleId={ruleId}
            jobId={jobId}
          />
        </Grid>

        <Grid xs={1} sx={{ display: "flex", justifyContent: "end", pr: 0.5 }}>
          <ModeToggle responsive={responsive} />
        </Grid>
      </Grid>

      <Grid
        xs={11}
        sx={{
          display: "flex",
          alignItems: "flex-start",
          justifyContent: "flex-start",
          height: "100vh",
        }}
      >
        <BuildTable
          responsive={responsive}
          portrait={portrait}
          windowWidth={windowWidth}
          windowHeight={windowHeight}
          repositoryId={repositoryId}
          ruleId={ruleId}
          jobId={jobId}
          taskId={taskId}
        />
      </Grid>
    </Sheet>
  );
}

export default App;
