CREATE SCHEMA IF NOT EXISTS public;

SET LOCAL search_path = public;

--
-- repository
--

CREATE TABLE
    public.repository (
        id uuid PRIMARY KEY NOT NULL UNIQUE DEFAULT gen_random_uuid (),
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_at timestamptz NOT NULL DEFAULT now(),
        deleted_at timestamptz NULL DEFAULT NULL,
        url text NOT NULL,
        name text NULL,
        handled_at timestamptz NOT NULL DEFAULT to_timestamp(0),
        repository_syncer_claimed_until timestamptz NOT NULL DEFAULT to_timestamp(0)
    );

ALTER TABLE public.repository OWNER TO postgres;

CREATE
OR REPLACE FUNCTION create_repository () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = now();
  NEW.updated_at = now();
  NEW.deleted_at = null;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_repository BEFORE INSERT ON public.repository FOR EACH ROW
EXECUTE PROCEDURE create_repository ();

CREATE
OR REPLACE FUNCTION update_repository () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = OLD.created_at;
  NEW.updated_at = now();
  IF OLD.deleted_at IS NOT null AND NEW.deleted_at IS NOT null THEN
    NEW.deleted_at = OLD.deleted_at;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_repository BEFORE
UPDATE ON public.repository FOR EACH ROW
EXECUTE PROCEDURE public.update_repository ();

CREATE RULE delete_repository AS ON DELETE TO public.repository
DO INSTEAD (
    UPDATE public.repository
    SET
        created_at = old.created_at,
        updated_at = now(),
        deleted_at = now()
    WHERE
        id = old.id
        AND deleted_at IS null
);

--
-- change
--

CREATE TABLE
    public.change (
        id uuid PRIMARY KEY NOT NULL UNIQUE DEFAULT gen_random_uuid (),
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_at timestamptz NOT NULL DEFAULT now(),
        deleted_at timestamptz NULL DEFAULT NULL,
        commit_hash text NOT NULL,
        branch text NOT NULL,
        tag text NULL,
        message text NOT NULL,
        authored_by text NOT NULL,
        authored_at timestamptz NOT NULL,
        committed_by text NOT NULL,
        committed_at timestamptz NOT NULL,
        handled_at timestamptz NULL,
        job_coordinator_claimed_until timestamptz NOT NULL DEFAULT to_timestamp(0)
    );

ALTER TABLE public.change OWNER TO postgres;

CREATE
OR REPLACE FUNCTION create_change () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = now();
  NEW.updated_at = now();
  NEW.deleted_at = null;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_change BEFORE INSERT ON public.change FOR EACH ROW
EXECUTE PROCEDURE create_change ();

CREATE
OR REPLACE FUNCTION update_change () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = OLD.created_at;
  NEW.updated_at = now();
  IF OLD.deleted_at IS NOT null AND NEW.deleted_at IS NOT null THEN
    NEW.deleted_at = OLD.deleted_at;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_change BEFORE
UPDATE ON public.change FOR EACH ROW
EXECUTE PROCEDURE public.update_change ();

CREATE RULE delete_change AS ON DELETE TO public.change
DO INSTEAD (
    UPDATE public.change
    SET
        created_at = old.created_at,
        updated_at = now(),
        deleted_at = now()
    WHERE
        id = old.id
        AND deleted_at IS null
);

--
-- job
--

CREATE TABLE
    public.job (
        id uuid PRIMARY KEY NOT NULL UNIQUE DEFAULT gen_random_uuid (),
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_at timestamptz NOT NULL DEFAULT now(),
        deleted_at timestamptz NULL DEFAULT NULL,
        name text NOT NULL,
        branches text NULL,
        tags text NULL
    );

ALTER TABLE public.job OWNER TO postgres;

CREATE
OR REPLACE FUNCTION create_job () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = now();
  NEW.updated_at = now();
  NEW.deleted_at = null;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_job BEFORE INSERT ON public.job FOR EACH ROW
EXECUTE PROCEDURE create_job ();

CREATE
OR REPLACE FUNCTION update_job () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = OLD.created_at;
  NEW.updated_at = now();
  IF OLD.deleted_at IS NOT null AND NEW.deleted_at IS NOT null THEN
    NEW.deleted_at = OLD.deleted_at;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_job BEFORE
UPDATE ON public.job FOR EACH ROW
EXECUTE PROCEDURE public.update_job ();

CREATE RULE delete_job AS ON DELETE TO public.job
DO INSTEAD (
    UPDATE public.job
    SET
        created_at = old.created_at,
        updated_at = now(),
        deleted_at = now()
    WHERE
        id = old.id
        AND deleted_at IS null
);

--
-- depends_on
--

CREATE TABLE
    public.depends_on (
        id uuid PRIMARY KEY NOT NULL UNIQUE DEFAULT gen_random_uuid (),
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_at timestamptz NOT NULL DEFAULT now(),
        deleted_at timestamptz NULL DEFAULT NULL
    );

ALTER TABLE public.depends_on OWNER TO postgres;

CREATE
OR REPLACE FUNCTION create_depends_on () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = now();
  NEW.updated_at = now();
  NEW.deleted_at = null;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_depends_on BEFORE INSERT ON public.depends_on FOR EACH ROW
EXECUTE PROCEDURE create_depends_on ();

CREATE
OR REPLACE FUNCTION update_depends_on () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = OLD.created_at;
  NEW.updated_at = now();
  IF OLD.deleted_at IS NOT null AND NEW.deleted_at IS NOT null THEN
    NEW.deleted_at = OLD.deleted_at;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_depends_on BEFORE
UPDATE ON public.depends_on FOR EACH ROW
EXECUTE PROCEDURE public.update_depends_on ();

CREATE RULE delete_depends_on AS ON DELETE TO public.depends_on
DO INSTEAD (
    UPDATE public.depends_on
    SET
        created_at = old.created_at,
        updated_at = now(),
        deleted_at = now()
    WHERE
        id = old.id
        AND deleted_at IS null
);

--
-- execution
--

CREATE TABLE
    public.execution (
        id uuid PRIMARY KEY NOT NULL UNIQUE DEFAULT gen_random_uuid (),
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_at timestamptz NOT NULL DEFAULT now(),
        deleted_at timestamptz NULL DEFAULT NULL,
        status text NOT NULL CHECK (status IN ('pending', 'running', 'succeeded', 'failing', 'failed', 'erroring', 'errored', 'skipped')),
        started_at timestamptz NULL DEFAULT NULL,
        ended_at timestamptz NULL DEFAULT NULL CHECK (ended_at IS null OR ended_at >= started_at),
        job_executor_claimed_until timestamptz NOT NULL DEFAULT to_timestamp(0)
    );

ALTER TABLE public.execution OWNER TO postgres;

CREATE
OR REPLACE FUNCTION create_execution () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = now();
  NEW.updated_at = now();
  NEW.deleted_at = null;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_execution BEFORE INSERT ON public.execution FOR EACH ROW
EXECUTE PROCEDURE create_execution ();

CREATE
OR REPLACE FUNCTION update_execution () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = OLD.created_at;
  NEW.updated_at = now();
  IF OLD.deleted_at IS NOT null AND NEW.deleted_at IS NOT null THEN
    NEW.deleted_at = OLD.deleted_at;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_execution BEFORE
UPDATE ON public.execution FOR EACH ROW
EXECUTE PROCEDURE public.update_execution ();

CREATE RULE delete_execution AS ON DELETE TO public.execution
DO INSTEAD (
    UPDATE public.execution
    SET
        created_at = old.created_at,
        updated_at = now(),
        deleted_at = now()
    WHERE
        id = old.id
        AND deleted_at IS null
);

--
-- task
--

CREATE TABLE
    public.task (
        id uuid PRIMARY KEY NOT NULL UNIQUE DEFAULT gen_random_uuid (),
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_at timestamptz NOT NULL DEFAULT now(),
        deleted_at timestamptz NULL DEFAULT NULL,
        name TEXT NOT NULL CHECK (name ~* '^[A-Za-z0-9-]+$'),
        index integer NOT NULL,
        platform text NOT NULL CHECK (platform ~* '^.*/.*$'),
        image text NOT NULL CHECK (image ~* '^(?:(?=[^:\/]{1,253})(?!-)[a-zA-Z0-9-]{1,63}(?<!-)(?:\.(?!-)[a-zA-Z0-9-]{1,63}(?<!-))*(?::[0-9]{1,5})?/)?((?![._-])(?:[a-z0-9._-]*)(?<![._-])(?:/(?![._-])[a-z0-9._-]*(?<![._-]))*)(?::(?![.-])[a-zA-Z0-9_.-]{1,128})?$'),
        script text NOT NULL
    );

ALTER TABLE public.task OWNER TO postgres;

CREATE
OR REPLACE FUNCTION create_task () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = now();
  NEW.updated_at = now();
  NEW.deleted_at = null;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_task BEFORE INSERT ON public.task FOR EACH ROW
EXECUTE PROCEDURE create_task ();

CREATE
OR REPLACE FUNCTION update_task () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = OLD.created_at;
  NEW.updated_at = now();
  IF OLD.deleted_at IS NOT null AND NEW.deleted_at IS NOT null THEN
    NEW.deleted_at = OLD.deleted_at;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_task BEFORE
UPDATE ON public.task FOR EACH ROW
EXECUTE PROCEDURE public.update_task ();

CREATE RULE delete_task AS ON DELETE TO public.task
DO INSTEAD (
    UPDATE public.task
    SET
        created_at = old.created_at,
        updated_at = now(),
        deleted_at = now()
    WHERE
        id = old.id
        AND deleted_at IS null
);

--
-- output
--

CREATE TABLE
    public.output (
        id uuid PRIMARY KEY NOT NULL UNIQUE DEFAULT gen_random_uuid (),
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_at timestamptz NOT NULL DEFAULT now(),
        deleted_at timestamptz NULL DEFAULT NULL,
        status text NOT NULL CHECK (status IN ('pending', 'running', 'succeeded', 'failing', 'failed', 'erroring', 'errored', 'skipped')),
        started_at timestamptz NULL DEFAULT NULL,
        ended_at timestamptz NULL DEFAULT NULL,
        exit_status int NOT NULL,
        error text NULL
    );

ALTER TABLE public.output OWNER TO postgres;

CREATE
OR REPLACE FUNCTION create_output () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = now();
  NEW.updated_at = now();
  NEW.deleted_at = null;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_output BEFORE INSERT ON public.output FOR EACH ROW
EXECUTE PROCEDURE create_output ();

CREATE
OR REPLACE FUNCTION update_output () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = OLD.created_at;
  NEW.updated_at = now();
  IF OLD.deleted_at IS NOT null AND NEW.deleted_at IS NOT null THEN
    NEW.deleted_at = OLD.deleted_at;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_output BEFORE
UPDATE ON public.output FOR EACH ROW
EXECUTE PROCEDURE public.update_output ();

CREATE RULE delete_output AS ON DELETE TO public.output
DO INSTEAD (
    UPDATE public.output
    SET
        created_at = old.created_at,
        updated_at = now(),
        deleted_at = now()
    WHERE
        id = old.id
        AND deleted_at IS null
);

--
-- log
--

CREATE TABLE
    public.log (
        id uuid PRIMARY KEY NOT NULL UNIQUE DEFAULT gen_random_uuid (),
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_at timestamptz NOT NULL DEFAULT now(),
        deleted_at timestamptz NULL DEFAULT NULL,
        buffer bytea NOT NULL
    );

ALTER TABLE public.log OWNER TO postgres;

CREATE
OR REPLACE FUNCTION create_log () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = now();
  NEW.updated_at = now();
  NEW.deleted_at = null;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_log BEFORE INSERT ON public.log FOR EACH ROW
EXECUTE PROCEDURE create_log ();

CREATE
OR REPLACE FUNCTION update_log () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = OLD.created_at;
  NEW.updated_at = now();
  IF OLD.deleted_at IS NOT null AND NEW.deleted_at IS NOT null THEN
    NEW.deleted_at = OLD.deleted_at;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_log BEFORE
UPDATE ON public.log FOR EACH ROW
EXECUTE PROCEDURE public.update_log ();

CREATE RULE delete_log AS ON DELETE TO public.log
DO INSTEAD (
    UPDATE public.log
    SET
        created_at = old.created_at,
        updated_at = now(),
        deleted_at = now()
    WHERE
        id = old.id
        AND deleted_at IS null
);

--
-- repository -> change (one-to-many)
--

ALTER TABLE public.change
ADD COLUMN repository_id uuid NOT NULL REFERENCES public.repository (id);

--
-- repository -> job (one-to-many)
--

ALTER TABLE public.job
ADD COLUMN repository_id uuid NOT NULL REFERENCES public.repository (id);

--
-- change -> execution (one-to-many)
--

ALTER TABLE public.execution
ADD COLUMN change_id uuid NOT NULL REFERENCES public.change (id);

--
-- job -> execution (one-to-many)
--

ALTER TABLE public.execution
ADD COLUMN job_id uuid NOT NULL REFERENCES public.job (id);

--
-- job -> task (one-to-many)
--

ALTER TABLE public.task
ADD COLUMN job_id uuid NOT NULL REFERENCES public.job (id);

--
-- job -> depends_on (one-to-many)
--

ALTER TABLE public.depends_on
ADD COLUMN source_job_id uuid NOT NULL REFERENCES public.job (id);

--
-- job -> depends_on (one-to-many)
--

ALTER TABLE public.depends_on
ADD COLUMN sink_job_id uuid NOT NULL REFERENCES public.job (id);

--
-- execution -> output (one-to-many)
--

ALTER TABLE public.output
ADD COLUMN execution_id uuid NOT NULL REFERENCES public.execution (id);

--
-- task -> output (one-to-many)
--

ALTER TABLE public.output
ADD COLUMN task_id uuid NOT NULL REFERENCES public.task (id);

--
-- output -> log (one-to-one)
--

ALTER TABLE public.log
ADD COLUMN output_id uuid NOT NULL REFERENCES public.output (id) DEFERRABLE INITIALLY DEFERRED;

CREATE UNIQUE INDEX log_unique_id_not_deleted ON public.log (id)
WHERE
    deleted_at IS null;

CREATE UNIQUE INDEX log_unique_id_deleted ON public.log (id, deleted_at)
WHERE
    deleted_at IS NOT null;

ALTER TABLE public.output
ADD COLUMN log_id uuid NOT NULL REFERENCES public.log (id) DEFERRABLE INITIALLY DEFERRED;

CREATE UNIQUE INDEX output_unique_id_not_deleted ON public.output (id)
WHERE
    deleted_at IS null;

CREATE UNIQUE INDEX output_unique_id_deleted ON public.output (id, deleted_at)
WHERE
    deleted_at IS NOT null;

--
-- change unique on (commit_hash, repository_id)
--

CREATE UNIQUE INDEX change_unique_commit_hash_repository_id_not_deleted ON public.change (commit_hash, repository_id)
WHERE
    deleted_at IS null;

CREATE UNIQUE INDEX change_unique_commit_hash_repository_id_deleted ON public.change (commit_hash, repository_id, deleted_at)
WHERE
    deleted_at IS NOT null;


--
-- job unique on (name)
--

CREATE UNIQUE INDEX job_unique_name_not_deleted ON public.job (name)
WHERE
    deleted_at IS null;

CREATE UNIQUE INDEX job_unique_name_deleted ON public.job (name, deleted_at)
WHERE
    deleted_at IS NOT null;


--
-- task unique on (name, job_id)
--

CREATE UNIQUE INDEX task_unique_name_job_id_not_deleted ON public.task (name, job_id)
WHERE
    deleted_at IS null;

CREATE UNIQUE INDEX task_unique_name_job_id_deleted ON public.task (name, job_id, deleted_at)
WHERE
    deleted_at IS NOT null;



