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
import { Execution } from "./Execution";
import { Logs } from "./Logs";

const defaultLimit = 10;

export interface BuildTableProps {
  responsive: boolean;
  portrait: boolean;
  windowWidth: number;
  windowHeight: number;
  repositoryId: string | undefined;
  branchFilter: string | undefined;
  ruleId: string | undefined;
  jobId: string | undefined;
  taskId: string | undefined;
}

export function BuildTable(props: BuildTableProps) {
  const [outputId, setOutputId] = useState<string | undefined>(undefined);
  const [showLogModal, setShowLogModal] = useState(false);

  const [ref, inView] = useInView();

  const queryHash = JSON.stringify(props);

  const relevantLimit = defaultLimit;

  const {
    data: infiniteChangesData,
    hasNextPage,
    fetchNextPage,
  } = useInfiniteQuery({
    queryKey: ["changes"],
    queryHash: queryHash,
    queryFn: async ({ pageParam = 0 }) => {
      const res = await clientForReactQuery.GET("/api/changes", {
        params: {
          query: {
            repository_id__eq: props.repositoryId,
            branch__ilike: props.branchFilter?.trim() ? props.branchFilter?.trim() : undefined,
            committed_at__desc: "",
            limit: relevantLimit,
            offset: pageParam,
            repository__load: "",
            referenced_by_execution__load: "",
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
        return undefined;
      }

      return (lastPage?.offset || 0) + relevantLimit;
    },
  });

  const changesData: {
    objects: components["schemas"]["Change"][];
  } = {
    objects: [],
  };

  const ids = new Set<string>();
  const createdAts: string[] = [];

  infiniteChangesData?.pages.forEach((page) => {
    page?.objects?.forEach((object) => {
      if (!object?.id) {
        return;
      }

      if (ids.has(object.id)) {
        return;
      }

      ids.add(object.id);
      if (object.created_at) {
        createdAts.push(object.created_at);
      }

      changesData?.objects.push(object);
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
            <th style={{ width: "200px", ...truncateStyleProps }}>{props.responsive ? "W" : "When"}</th>
            <th style={{ width: "250px", ...truncateStyleProps }}>{props.responsive ? "U" : "URL"}</th>
            <th style={{ width: "150px", ...truncateStyleProps }}>{props.responsive ? "B" : "Branch"}</th>
            <th style={{ width: "325px", ...truncateStyleProps }}>{props.responsive ? "C" : "Commit"}</th>
            <th style={{ ...truncateStyleProps }}>{props.responsive ? "E" : "Executions"}</th>
          </tr>
        </thead>
        <tbody>
          {changesData?.objects?.length ? (
            changesData?.objects.map((change, i) => {
              const repository = change?.repository_id_object;

              return (
                <tr key={`execution-table-row-${change.id}`}>
                  <td>{change?.authored_at}</td>
                  <td>
                    <Link
                      href={(repository?.url || "") + "/commit/" + (change?.commit_hash || "")}
                      target="_blank"
                      rel="noreferrer"
                    >
                      {repository?.url}
                    </Link>
                  </td>
                  <td>
                    <Link
                      href={(repository?.url || "") + "/tree/" + (change?.branch || "")}
                      target="_blank"
                      rel="noreferrer"
                    >
                      {change?.branch}
                    </Link>
                  </td>
                  <Tooltip size={"sm"} title={`${change?.authored_by} @ ${change?.authored_at}: ${change?.message}`}>
                    <td>
                      <>
                        <Link
                          href={(repository?.url || "") + "/commit/" + (change?.commit_hash || "")}
                          target="_blank"
                          rel="noreferrer"
                        >
                          {change?.commit_hash}
                        </Link>
                      </>
                    </td>
                  </Tooltip>
                  <td>
                    <Table size="sm" sx={{ p: 0, m: 0 }} borderAxis="y">
                      <tbody>
                        {change?.referenced_by_execution_change_id_objects?.map(
                          (execution) =>
                            execution.id && (
                              <Execution
                                executionId={execution.id}
                                setOutputId={setOutputId}
                                setShowLogModal={setShowLogModal}
                                key={execution.id}
                              />
                            )
                        )}
                      </tbody>
                    </Table>
                  </td>
                </tr>
              );
            })
          ) : (
            <tr>
              <td colSpan={5}>
                <Typography color={"neutral"}>(No executions for the selected filters)</Typography>
              </td>
            </tr>
          )}
          <tr>
            <td colSpan={5} ref={ref}>
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
