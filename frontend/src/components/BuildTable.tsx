import Button from "@mui/joy/Button";
import Chip from "@mui/joy/Chip";
import Link from "@mui/joy/Link";
import Modal from "@mui/joy/Modal";
import ModalDialog from "@mui/joy/ModalDialog";
import Table from "@mui/joy/Table";
import Tooltip from "@mui/joy/Tooltip";
import Typography from "@mui/joy/Typography";
import { useInfiniteQuery } from "@tanstack/react-query";
import { useEffect, useState } from "react";
import { useInView } from "react-intersection-observer";
import { clientForReactQuery } from "../api";
import { components } from "../api/api";
import { Logs } from "./Logs";

const defaultLimit = 10;

export interface BuildTableProps {
  responsive: boolean;
  portrait: boolean;
  windowWidth: number;
  windowHeight: number;
  repositoryId: string | undefined;
  ruleId: string | undefined;
  jobId: string | undefined;
  taskId: string | undefined;
}

export function BuildTable(props: BuildTableProps) {
  const [ref, inView] = useInView();

  const [outputId, setOutputId] = useState<string | undefined>(undefined);
  const [showLogModal, setShowLogModal] = useState(false);

  const queryHash = JSON.stringify(props);

  const relevantLimit = defaultLimit;

  const {
    data: infinitegExecutionsData,
    hasNextPage,
    fetchNextPage,
  } = useInfiniteQuery({
    queryKey: ["executions"],
    queryHash: queryHash,
    queryFn: async ({ pageParam = 0 }) => {
      const res = await clientForReactQuery.GET("/api/executions", {
        params: {
          query: {
            created_at__desc: "",
            depth: 4,
            limit: relevantLimit,
            offset: pageParam,
          },
        },
      });
      return res.data;
    },
    initialPageParam: 0,
    getNextPageParam: (lastPage, pages) => {
      /*
      TODO: this doesn't cater for the fact that we have new data coming in- really we should
      use something like timestamp for the cursor, and even then we should probably split out
      finished executions from processing executions
      */

      if (lastPage?.count === 0) {
        return lastPage?.offset;
      }

      return (lastPage?.offset || 0) + relevantLimit;
    },
  });

  const executionsData: {
    objects: components["schemas"]["Execution"][];
  } = {
    objects: [],
  };

  const ids = new Set<string>();

  infinitegExecutionsData?.pages.forEach((page) => {
    page?.objects?.forEach((object) => {
      if (!object?.id) {
        return;
      }

      if (ids.has(object.id)) {
        return;
      }

      ids.add(object.id);

      executionsData?.objects.push(object);
    });
  });

  const truncateStyleProps = props.responsive
    ? {
        maxWidth: "100%",
        overflow: "hidden",
        textOverflow: "ellipsis",
      }
    : {};

  useEffect(() => {
    if (inView && hasNextPage) {
      void fetchNextPage();
    }
  }, [fetchNextPage, hasNextPage, inView]);

  return (
    <>
      <Table
        size="sm"
        width="100%"
        sx={{
          width: "100vw",
          th: {
            textAlign: "center",
            p: 0,
            m: 0,
          },
          td: {
            textAlign: "center",
            p: 0,
            m: 0,
            pt: 0.66,
            ...truncateStyleProps,
          },
        }}
        stickyHeader={true}
      >
        <thead>
          <tr>
            <th style={{ ...truncateStyleProps }}>
              {props.responsive ? "U" : "URL"}
            </th>
            <th style={{ width: "10%", ...truncateStyleProps }}>
              {props.responsive ? "B" : "Branch"}
            </th>
            <th style={{ ...truncateStyleProps }}>
              {props.responsive ? "C" : "Commit"}
            </th>
            <th colSpan={2} style={{ width: "33%", ...truncateStyleProps }}>
              {props.responsive ? "J" : "Job"}
            </th>
            <th style={{ ...truncateStyleProps }}>
              {props.responsive ? "T" : "Tasks"}
            </th>
          </tr>
        </thead>
        <tbody>
          {executionsData?.objects?.length ? (
            executionsData?.objects
              ?.filter((execution) => {
                const rule = execution?.trigger_id_object?.rule_id_object;
                const repository = rule?.repository_id_object;
                const job = execution?.job_id_object;

                if (
                  props.repositoryId &&
                  repository?.id !== props?.repositoryId
                ) {
                  return false;
                }

                if (props.ruleId && rule?.id !== props?.ruleId) {
                  return false;
                }

                if (props.jobId && rule?.id !== job?.id) {
                  return false;
                }

                if (props.taskId) {
                  if (
                    execution?.build_output_id_object?.task_id &&
                    execution?.build_output_id_object?.task_id !== props.taskId
                  ) {
                    return false;
                  }

                  if (
                    execution?.test_output_id_object?.task_id &&
                    execution?.test_output_id_object?.task_id !== props.taskId
                  ) {
                    return false;
                  }

                  if (
                    execution?.publish_output_id_object?.task_id &&
                    execution?.publish_output_id_object?.task_id !==
                      props.taskId
                  ) {
                    return false;
                  }

                  if (
                    execution?.deploy_output_id_object?.task_id &&
                    execution?.deploy_output_id_object?.task_id !== props.taskId
                  ) {
                    return false;
                  }

                  if (
                    execution?.validate_output_id_object?.task_id &&
                    execution?.validate_output_id_object?.task_id !==
                      props.taskId
                  ) {
                    return false;
                  }
                  return false;
                }

                return true;
              })
              .map((execution, i) => {
                const rule = execution?.trigger_id_object?.rule_id_object;
                const repository = rule?.repository_id_object;
                const change = execution?.trigger_id_object?.change_id_object;
                const job = execution?.job_id_object;

                const outputs = [];

                if (execution?.build_output_id_object) {
                  const output = execution?.build_output_id_object;

                  outputs.push({
                    id: output?.id,
                    phase: "build",
                    status: output?.status,
                  });
                }

                if (execution?.test_output_id_object) {
                  const output = execution?.test_output_id_object;

                  outputs.push({
                    id: output?.id,
                    phase: "test",
                    status: output?.status,
                  });
                }

                if (execution?.publish_output_id_object) {
                  const output = execution?.publish_output_id_object;

                  outputs.push({
                    id: output?.id,
                    phase: "publish",
                    status: output?.status,
                  });
                }

                if (execution?.deploy_output_id_object) {
                  const output = execution?.deploy_output_id_object;

                  outputs.push({
                    id: output?.id,
                    phase: "deploy",
                    status: output?.status,
                  });
                }

                if (execution?.validate_output_id_object) {
                  const output = execution?.validate_output_id_object;

                  outputs.push({
                    id: output?.id,
                    phase: "validate",
                    status: output?.status,
                  });
                }

                const statusText =
                  (execution?.status || "").slice(0, 1).toUpperCase() +
                  (execution?.status || "").slice(1);

                let status = <Chip color={"danger"}>{statusText}</Chip>;

                if (execution?.status === "running") {
                  status = <Chip color={"warning"}>{statusText}</Chip>;
                } else if (execution?.status === "succeeded") {
                  status = <Chip color={"success"}>{statusText}</Chip>;
                }

                return (
                  <tr key={`execution-table-row-${execution.id}`}>
                    <td>
                      <Link
                        href={
                          (repository?.url || "") +
                          "/commit/" +
                          (change?.commit_hash || "")
                        }
                        target="_blank"
                        rel="noreferrer"
                      >
                        {repository?.url}
                      </Link>
                    </td>
                    <td>
                      <Link
                        href={
                          (repository?.url || "") +
                          "/tree/" +
                          (change?.branch_name || "")
                        }
                        target="_blank"
                        rel="noreferrer"
                      >
                        {change?.branch_name}
                      </Link>
                    </td>
                    <Tooltip
                      size={"sm"}
                      title={`${change?.authored_by} @ ${change?.authored_at}: ${change?.message}`}
                    >
                      <td>
                        <>
                          <Link
                            href={
                              (repository?.url || "") +
                              "/commit/" +
                              (change?.commit_hash || "")
                            }
                            target="_blank"
                            rel="noreferrer"
                          >
                            {change?.commit_hash}
                          </Link>
                        </>
                      </td>
                    </Tooltip>
                    <td>{job?.name}</td>
                    <td>{status}</td>
                    <td style={{ padding: 0, margin: 0 }}>
                      <Table size="sm" sx={{ p: 0, m: 0 }} borderAxis="y">
                        <tbody>
                          {outputs.map((output) => {
                            const statusText =
                              (output.status || "").slice(0, 1).toUpperCase() +
                              (output.status || "").slice(1);

                            let status = (
                              <Chip color={"danger"}>{statusText}</Chip>
                            );

                            if (output?.status === "running") {
                              status = (
                                <Chip color={"warning"}>{statusText}</Chip>
                              );
                            } else if (output?.status === "succeeded") {
                              status = (
                                <Chip color={"success"}>{statusText}</Chip>
                              );
                            }

                            return (
                              <tr>
                                <td style={{ width: "33%" }}>
                                  {output.phase.slice(0, 1).toUpperCase() +
                                    output.phase.slice(1)}
                                </td>
                                <td style={{ width: "33%" }}>{status}</td>
                                <td style={{ width: "33%" }}>
                                  <Button
                                    size={"sm"}
                                    variant="soft"
                                    color={"primary"}
                                    sx={{
                                      fontSize:
                                        "var(--joy-fontSize-xs, 0.75rem)",
                                    }}
                                    onClick={() => {
                                      setOutputId(output?.id);
                                      setShowLogModal(true);
                                    }}
                                  >
                                    Logs
                                  </Button>
                                </td>
                              </tr>
                            );
                          })}
                        </tbody>
                      </Table>
                    </td>
                  </tr>
                );
              })
          ) : (
            <tr>
              <td colSpan={6}>
                <Typography color={"neutral"}>
                  (No executions for the selected filters)
                </Typography>
              </td>
            </tr>
          )}
          <tr>
            <td colSpan={6} ref={ref}>
              <Typography color={"neutral"}> </Typography>
            </td>
          </tr>
        </tbody>
      </Table>
      <Modal
        sx={{ p: 0, m: 0 }}
        open={showLogModal}
        onClose={() => {
          setShowLogModal(false);
        }}
      >
        <ModalDialog
          variant="plain"
          size="sm"
          sx={{
            width: "97.5%",
            height: "97.5%",
            p: "1px",
            borderRadius: 1,
            m: 0,
          }}
        >
          <Logs outputId={outputId || ""} />
        </ModalDialog>
      </Modal>
    </>
  );
}
