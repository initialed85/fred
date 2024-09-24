--
-- init
--
CREATE SCHEMA IF NOT EXISTS public;

ALTER SCHEMA public OWNER TO postgres;

COMMENT ON SCHEMA public IS 'standard public schema';

SET
    default_tablespace = '';

SET
    default_table_access_method = heap;

CREATE EXTENSION IF NOT EXISTS postgis SCHEMA public;

CREATE EXTENSION IF NOT EXISTS postgis_raster SCHEMA public;

SET
    postgis.gdal_enabled_drivers = 'ENABLE_ALL';

CREATE EXTENSION IF NOT EXISTS hstore SCHEMA public;

ALTER ROLE postgres
SET
    search_path TO public,
    postgis,
    hstore;

SET
    search_path TO public,
    postgis,
    hstore;

--
-- repository
--
DROP TABLE IF EXISTS public.repository CASCADE;

CREATE TABLE
    public.repository (
        id uuid PRIMARY KEY NOT NULL UNIQUE DEFAULT gen_random_uuid (),
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_at timestamptz NOT NULL DEFAULT now(),
        deleted_at timestamptz NULL DEFAULT NULL,
        last_synced timestamptz NOT NULL DEFAULT to_timestamp(0),
        url text NOT NULL CHECK (trim(url) != ''),
        username text NULL CHECK (
            username IS null
            OR trim(username) != ''
        ),
        password text NULL CHECK (
            password IS null
            OR trim(password) != ''
        ),
        ssh_key text NULL CHECK (
            ssh_key IS null
            OR trim(ssh_key) != ''
        )
    );

ALTER TABLE public.repository OWNER TO postgres;

CREATE UNIQUE INDEX repository_unique_url_not_deleted ON public.repository (url)
WHERE
    deleted_at IS null;

CREATE UNIQUE INDEX repository_unique_url_deleted ON public.repository (url, deleted_at)
WHERE
    deleted_at IS NOT null;

--
-- change
--
DROP TABLE IF EXISTS public.change CASCADE;

CREATE TABLE
    public.change (
        id uuid PRIMARY KEY NOT NULL UNIQUE DEFAULT gen_random_uuid (),
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_at timestamptz NOT NULL DEFAULT now(),
        deleted_at timestamptz NULL DEFAULT NULL,
        branch_name text NOT NULL,
        commit_hash text NOT NULL,
        message text NOT NULL,
        authored_by text NOT NULL,
        authored_at timestamptz NOT NULL,
        committed_by text NOT NULL,
        committed_at timestamptz NOT NULL,
        trigger_produced_at timestamptz NULL DEFAULT NULL,
        repository_id uuid NOT NULL REFERENCES public.repository (id)
    );

ALTER TABLE public.change OWNER TO postgres;

CREATE UNIQUE INDEX change_unique_branch_name_commit_hash_repository_id_not_deleted ON public.change (branch_name, commit_hash, repository_id)
WHERE
    deleted_at IS null;

CREATE UNIQUE INDEX change_unique_branch_name_commit_hash_repository_id_deleted ON public.change (branch_name, commit_hash, repository_id, deleted_at)
WHERE
    deleted_at IS NOT null;

--
-- rule
--
DROP TABLE IF EXISTS public.rule CASCADE;

CREATE TABLE
    public.rule (
        id uuid PRIMARY KEY NOT NULL UNIQUE DEFAULT gen_random_uuid (),
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_at timestamptz NOT NULL DEFAULT now(),
        deleted_at timestamptz NULL DEFAULT NULL,
        branch_name text NULL DEFAULT NULL CHECK (
            branch_name IS null
            OR trim(branch_name) != ''
        ),
        repository_id uuid NOT NULL REFERENCES public.repository (id)
    );

ALTER TABLE public.rule OWNER TO postgres;

CREATE UNIQUE INDEX rule_unique_branch_name_repository_id_not_deleted ON public.rule (branch_name, repository_id)
WHERE
    deleted_at IS null;

CREATE UNIQUE INDEX rule_unique_branch_name_repository_id_deleted ON public.rule (branch_name, repository_id, deleted_at)
WHERE
    deleted_at IS NOT null;

--
-- trigger
--
DROP TABLE IF EXISTS public.trigger CASCADE;

CREATE TABLE
    public.trigger (
        id uuid PRIMARY KEY NOT NULL UNIQUE DEFAULT gen_random_uuid (),
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_at timestamptz NOT NULL DEFAULT now(),
        deleted_at timestamptz NULL DEFAULT NULL,
        job_executor_claimed_until timestamptz NOT NULL DEFAULT to_timestamp(0),
        job_execution_started_at timestamptz NULL DEFAULT NULL,
        rule_id uuid NOT NULL REFERENCES public.rule (id),
        change_id uuid NOT NULL REFERENCES public.change (id)
    );

ALTER TABLE public.trigger OWNER TO postgres;

--
-- task
--
DROP TABLE IF EXISTS public.task CASCADE;

CREATE TABLE
    public.task (
        id uuid PRIMARY KEY NOT NULL UNIQUE DEFAULT gen_random_uuid (),
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_at timestamptz NOT NULL DEFAULT now(),
        deleted_at timestamptz NULL DEFAULT NULL,
        name TEXT NOT NULL UNIQUE,
        platform text NOT NULL,
        image text NOT NULL,
        script text NOT NULL
    );

ALTER TABLE public.task OWNER TO postgres;

--
-- job
--
DROP TABLE IF EXISTS public.job CASCADE;

CREATE TABLE
    public.job (
        id uuid PRIMARY KEY NOT NULL UNIQUE DEFAULT gen_random_uuid (),
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_at timestamptz NOT NULL DEFAULT now(),
        deleted_at timestamptz NULL DEFAULT NULL,
        name TEXT NOT NULL UNIQUE,
        job_executor_claimed_until timestamptz NOT NULL DEFAULT to_timestamp(0),
        rule_id uuid NOT NULL REFERENCES public.rule (id),
        build_task_id uuid NULL REFERENCES public.task (id),
        test_task_id uuid NULL REFERENCES public.task (id),
        publish_task_id uuid NULL REFERENCES public.task (id),
        deploy_task_id uuid NULL REFERENCES public.task (id),
        validate_task_id uuid NULL REFERENCES public.task (id)
    );

ALTER TABLE public.job OWNER TO postgres;

--
-- rule_requires_job
--
DROP TABLE IF EXISTS public.rule_requires_job CASCADE;

CREATE TABLE
    public.rule_requires_job (
        id uuid PRIMARY KEY NOT NULL UNIQUE DEFAULT gen_random_uuid (),
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_at timestamptz NOT NULL DEFAULT now(),
        deleted_at timestamptz NULL DEFAULT NULL,
        rule_id uuid NOT NULL REFERENCES public.rule (id),
        job_id uuid NOT NULL REFERENCES public.job (id)
    );

ALTER TABLE public.rule_requires_job OWNER TO postgres;

--
-- output
--
DROP TABLE IF EXISTS public.output CASCADE;

CREATE TABLE
    public.output (
        id uuid PRIMARY KEY NOT NULL UNIQUE DEFAULT gen_random_uuid (),
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_at timestamptz NOT NULL DEFAULT now(),
        deleted_at timestamptz NULL DEFAULT NULL,
        status text NOT NULL,
        exit_status int NOT NULL,
        buffer text NOT NULL,
        error text NULL,
        task_id uuid NOT NULL REFERENCES public.task (id)
    );

ALTER TABLE public.output OWNER TO postgres;

--
-- execution
--
DROP TABLE IF EXISTS public.execution CASCADE;

CREATE TABLE
    public.execution (
        id uuid PRIMARY KEY NOT NULL UNIQUE DEFAULT gen_random_uuid (),
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_at timestamptz NOT NULL DEFAULT now(),
        deleted_at timestamptz NULL DEFAULT NULL,
        status text NOT NULL,
        trigger_id uuid NOT NULL REFERENCES public.trigger (id),
        build_output_id uuid NULL REFERENCES public.output (id),
        test_output_id uuid NULL REFERENCES public.output (id),
        publish_output_id uuid NULL REFERENCES public.output (id),
        deploy_output_id uuid NULL REFERENCES public.output (id),
        validate_output_id uuid NULL REFERENCES public.output (id),
        job_id uuid NOT NULL REFERENCES public.job (id)
    );

ALTER TABLE public.execution OWNER TO postgres;

--
-- trigger_has_execution
--
DROP TABLE IF EXISTS public.trigger_has_execution CASCADE;

CREATE TABLE
    public.trigger_has_execution (
        id uuid PRIMARY KEY NOT NULL UNIQUE DEFAULT gen_random_uuid (),
        created_at timestamptz NOT NULL DEFAULT now(),
        updated_at timestamptz NOT NULL DEFAULT now(),
        deleted_at timestamptz NULL DEFAULT NULL,
        trigger_id uuid NOT NULL REFERENCES public.trigger (id),
        execution_id uuid NOT NULL REFERENCES public.execution (id)
    );

ALTER TABLE public.trigger_has_execution OWNER TO postgres;

--
-- triggers for repository
--
CREATE
OR REPLACE FUNCTION create_repository () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = now();
  NEW.updated_at = now();
  NEW.deleted_at = null;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_repository BEFORE INSERT ON repository FOR EACH ROW
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
UPDATE ON repository FOR EACH ROW
EXECUTE PROCEDURE update_repository ();

CREATE RULE "delete_repository" AS ON DELETE TO "repository"
DO INSTEAD (
    UPDATE repository
    SET
        created_at = old.created_at,
        updated_at = now(),
        deleted_at = now()
    WHERE
        id = old.id
        AND deleted_at IS null
);

CREATE RULE "delete_repository_cascade_to_change" AS ON DELETE TO "repository"
DO ALSO (
    DELETE FROM change
    WHERE
        repository_id = old.id
        AND deleted_at IS null
);

CREATE RULE "delete_repository_cascade_to_rule" AS ON DELETE TO "repository"
DO ALSO (
    DELETE FROM rule
    WHERE
        repository_id = old.id
        AND deleted_at IS null
);

--
-- triggers for change
--
CREATE
OR REPLACE FUNCTION create_change () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = now();
  NEW.updated_at = now();
  NEW.deleted_at = null;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_change BEFORE INSERT ON change FOR EACH ROW
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
UPDATE ON change FOR EACH ROW
EXECUTE PROCEDURE update_change ();

CREATE RULE "delete_change" AS ON DELETE TO "change"
DO INSTEAD (
    UPDATE change
    SET
        created_at = old.created_at,
        updated_at = now(),
        deleted_at = now()
    WHERE
        id = old.id
        AND deleted_at IS null
);

CREATE RULE "delete_change_cascade_to_trigger" AS ON DELETE TO "change"
DO ALSO (
    DELETE FROM trigger
    WHERE
        change_id = old.id
        AND deleted_at IS null
);

--
-- triggers for rule
--
CREATE
OR REPLACE FUNCTION create_rule () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = now();
  NEW.updated_at = now();
  NEW.deleted_at = null;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_rule BEFORE INSERT ON rule FOR EACH ROW
EXECUTE PROCEDURE create_rule ();

CREATE
OR REPLACE FUNCTION update_rule () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = OLD.created_at;
  NEW.updated_at = now();
  IF OLD.deleted_at IS NOT null AND NEW.deleted_at IS NOT null THEN
    NEW.deleted_at = OLD.deleted_at;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_rule BEFORE
UPDATE ON rule FOR EACH ROW
EXECUTE PROCEDURE update_rule ();

CREATE RULE "delete_rule" AS ON DELETE TO "rule"
DO INSTEAD (
    UPDATE rule
    SET
        created_at = old.created_at,
        updated_at = now(),
        deleted_at = now()
    WHERE
        id = old.id
        AND deleted_at IS null
);

CREATE RULE "delete_rule_cascade_to_trigger" AS ON DELETE TO "rule"
DO ALSO (
    DELETE FROM trigger
    WHERE
        rule_id = old.id
        AND deleted_at IS null
);

CREATE RULE "delete_rule_cascade_to_rule_requires_job" AS ON DELETE TO "rule"
DO ALSO (
    DELETE FROM rule_requires_job
    WHERE
        rule_id = old.id
        AND deleted_at IS null
);

CREATE RULE "delete_rule_cascade_to_job" AS ON DELETE TO "rule"
DO ALSO (
    DELETE FROM job
    WHERE
        rule_id = old.id
        AND deleted_at IS null
);

--
-- triggers for trigger
--
CREATE
OR REPLACE FUNCTION create_trigger () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = now();
  NEW.updated_at = now();
  NEW.deleted_at = null;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_trigger BEFORE INSERT ON trigger FOR EACH ROW
EXECUTE PROCEDURE create_trigger ();

CREATE
OR REPLACE FUNCTION update_trigger () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = OLD.created_at;
  NEW.updated_at = now();
  IF OLD.deleted_at IS NOT null AND NEW.deleted_at IS NOT null THEN
    NEW.deleted_at = OLD.deleted_at;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_trigger BEFORE
UPDATE ON trigger FOR EACH ROW
EXECUTE PROCEDURE update_trigger ();

CREATE RULE "delete_trigger" AS ON DELETE TO "trigger"
DO INSTEAD (
    UPDATE trigger
    SET
        created_at = old.created_at,
        updated_at = now(),
        deleted_at = now()
    WHERE
        id = old.id
        AND deleted_at IS null
);

CREATE RULE "delete_trigger_cascade_to_trigger_has_execution" AS ON DELETE TO "trigger"
DO ALSO (
    DELETE FROM trigger_has_execution
    WHERE
        trigger_id = old.id
        AND deleted_at IS null
);

CREATE RULE "delete_trigger_cascade_to_execution" AS ON DELETE TO "trigger"
DO ALSO (
    DELETE FROM execution
    WHERE
        trigger_id = old.id
        AND deleted_at IS null
);

--
-- triggers for task
--
CREATE
OR REPLACE FUNCTION create_task () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = now();
  NEW.updated_at = now();
  NEW.deleted_at = null;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_task BEFORE INSERT ON task FOR EACH ROW
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
UPDATE ON task FOR EACH ROW
EXECUTE PROCEDURE update_task ();

CREATE RULE "delete_task" AS ON DELETE TO "task"
DO INSTEAD (
    UPDATE task
    SET
        created_at = old.created_at,
        updated_at = now(),
        deleted_at = now()
    WHERE
        id = old.id
        AND deleted_at IS null
);

CREATE RULE "delete_task_cascade_to_output" AS ON DELETE TO "task"
DO ALSO (
    DELETE FROM output
    WHERE
        task_id = old.id
        AND deleted_at IS null
);

CREATE RULE "delete_task_cascade_to_job_build_task" AS ON DELETE TO "task"
DO ALSO (
    UPDATE job
    SET
        build_task_id = null
    WHERE
        build_task_id = old.id
        AND deleted_at IS null
);

CREATE RULE "delete_task_cascade_to_job_test_task" AS ON DELETE TO "task"
DO ALSO (
    UPDATE job
    SET
        test_task_id = null
    WHERE
        test_task_id = old.id
        AND deleted_at IS null
);

CREATE RULE "delete_task_cascade_to_job_publish_task" AS ON DELETE TO "task"
DO ALSO (
    UPDATE job
    SET
        publish_task_id = null
    WHERE
        publish_task_id = old.id
        AND deleted_at IS null
);

CREATE RULE "delete_task_cascade_to_job_deploy_task" AS ON DELETE TO "task"
DO ALSO (
    UPDATE job
    SET
        deploy_task_id = null
    WHERE
        deploy_task_id = old.id
        AND deleted_at IS null
);

CREATE RULE "delete_task_cascade_to_job_validate_task" AS ON DELETE TO "task"
DO ALSO (
    UPDATE job
    SET
        validate_task_id = null
    WHERE
        validate_task_id = old.id
        AND deleted_at IS null
);

--
-- triggers for job
--
CREATE
OR REPLACE FUNCTION create_job () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = now();
  NEW.updated_at = now();
  NEW.deleted_at = null;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_job BEFORE INSERT ON job FOR EACH ROW
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
UPDATE ON job FOR EACH ROW
EXECUTE PROCEDURE update_job ();

CREATE RULE "delete_job" AS ON DELETE TO "job"
DO INSTEAD (
    UPDATE job
    SET
        created_at = old.created_at,
        updated_at = now(),
        deleted_at = now()
    WHERE
        id = old.id
        AND deleted_at IS null
);

CREATE RULE "delete_job_cascade_to_rule_requires_job" AS ON DELETE TO "job"
DO ALSO (
    DELETE FROM rule_requires_job
    WHERE
        job_id = old.id
        AND deleted_at IS null
);

--
-- triggers for rule_requires_job
--
CREATE
OR REPLACE FUNCTION create_rule_requires_job () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = now();
  NEW.updated_at = now();
  NEW.deleted_at = null;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_rule_requires_job BEFORE INSERT ON rule_requires_job FOR EACH ROW
EXECUTE PROCEDURE create_rule_requires_job ();

CREATE
OR REPLACE FUNCTION update_rule_requires_job () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = OLD.created_at;
  NEW.updated_at = now();
  IF OLD.deleted_at IS NOT null AND NEW.deleted_at IS NOT null THEN
    NEW.deleted_at = OLD.deleted_at;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_rule_requires_job BEFORE
UPDATE ON rule_requires_job FOR EACH ROW
EXECUTE PROCEDURE update_rule_requires_job ();

CREATE RULE "delete_rule_requires_job" AS ON DELETE TO "rule_requires_job"
DO INSTEAD (
    UPDATE rule_requires_job
    SET
        created_at = old.created_at,
        updated_at = now(),
        deleted_at = now()
    WHERE
        id = old.id
        AND deleted_at IS null
);

--
-- triggers for output
--
CREATE
OR REPLACE FUNCTION create_output () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = now();
  NEW.updated_at = now();
  NEW.deleted_at = null;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_output BEFORE INSERT ON output FOR EACH ROW
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
UPDATE ON output FOR EACH ROW
EXECUTE PROCEDURE update_output ();

CREATE RULE "delete_output" AS ON DELETE TO "output"
DO INSTEAD (
    UPDATE output
    SET
        created_at = old.created_at,
        updated_at = now(),
        deleted_at = now()
    WHERE
        id = old.id
        AND deleted_at IS null
);

CREATE RULE "delete_output_cascade_to_execution_build_output" AS ON DELETE TO "output"
DO ALSO (
    UPDATE execution
    SET
        build_output_id = null
    WHERE
        build_output_id = old.id
        AND deleted_at IS null
);

CREATE RULE "delete_output_cascade_to_execution_test_output" AS ON DELETE TO "output"
DO ALSO (
    UPDATE execution
    SET
        test_output_id = null
    WHERE
        test_output_id = old.id
        AND deleted_at IS null
);

CREATE RULE "delete_output_cascade_to_execution_publish_output" AS ON DELETE TO "output"
DO ALSO (
    UPDATE execution
    SET
        publish_output_id = null
    WHERE
        publish_output_id = old.id
        AND deleted_at IS null
);

CREATE RULE "delete_output_cascade_to_execution_deploy_output" AS ON DELETE TO "output"
DO ALSO (
    UPDATE execution
    SET
        deploy_output_id = null
    WHERE
        deploy_output_id = old.id
        AND deleted_at IS null
);

CREATE RULE "delete_output_cascade_to_execution_validate_output" AS ON DELETE TO "output"
DO ALSO (
    UPDATE execution
    SET
        validate_output_id = null
    WHERE
        validate_output_id = old.id
        AND deleted_at IS null
);

--
-- triggers for execution
--
CREATE
OR REPLACE FUNCTION create_execution () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = now();
  NEW.updated_at = now();
  NEW.deleted_at = null;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_execution BEFORE INSERT ON execution FOR EACH ROW
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
UPDATE ON execution FOR EACH ROW
EXECUTE PROCEDURE update_execution ();

CREATE RULE "delete_execution" AS ON DELETE TO "execution"
DO INSTEAD (
    UPDATE execution
    SET
        created_at = old.created_at,
        updated_at = now(),
        deleted_at = now()
    WHERE
        id = old.id
        AND deleted_at IS null
);

--
-- triggers for trigger_has_execution
--
CREATE
OR REPLACE FUNCTION create_trigger_has_execution () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = now();
  NEW.updated_at = now();
  NEW.deleted_at = null;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_trigger_has_execution BEFORE INSERT ON trigger_has_execution FOR EACH ROW
EXECUTE PROCEDURE create_trigger_has_execution ();

CREATE
OR REPLACE FUNCTION update_trigger_has_execution () RETURNS TRIGGER AS $$
BEGIN
  NEW.created_at = OLD.created_at;
  NEW.updated_at = now();
  IF OLD.deleted_at IS NOT null AND NEW.deleted_at IS NOT null THEN
    NEW.deleted_at = OLD.deleted_at;
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_trigger_has_execution BEFORE
UPDATE ON trigger_has_execution FOR EACH ROW
EXECUTE PROCEDURE update_trigger_has_execution ();

CREATE RULE "delete_trigger_has_execution" AS ON DELETE TO "trigger_has_execution"
DO INSTEAD (
    UPDATE trigger_has_execution
    SET
        created_at = old.created_at,
        updated_at = now(),
        deleted_at = now()
    WHERE
        id = old.id
        AND deleted_at IS null
);
