CREATE OR REPLACE FUNCTION depot_create(js json)
	RETURNS void
	LANGUAGE plpgsql
AS $function$
DECLARE
	_u users;
	_d depots;
BEGIN
	SELECT *
	FROM users
	WHERE tele_id = (js->>'mammoth_id')::BIGINT
	INTO _u;

	IF _u.id ISNULL THEN
		RAISE EXCEPTION 'пользователь не найден';
	END IF;

	INSERT INTO depots(mammoth_id, mammoth_username, worker_id, worker_username, amount, is_show_name) VALUES((js->>'mammoth_id')::BIGINT, js->>'mammoth_username', (js->>'worker_id')::BIGINT, js->>'worker_username', (js->>'amount')::NUMERIC, (js->>'is_show_name')::BOOLEAN);
END;
$function$