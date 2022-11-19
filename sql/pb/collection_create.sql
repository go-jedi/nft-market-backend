CREATE OR REPLACE FUNCTION collection_create(js json, _uid character varying)
	RETURNS void
	LANGUAGE plpgsql
AS $function$
DECLARE
	_c collections;
BEGIN
	SELECT *
	FROM collections
	WHERE name = js->>'name'
	INTO _c;

	IF _c.id ISNULL THEN
		INSERT INTO collections (name, count, collection_uid)
		VALUES(js->>'name', (js->>'count')::INTEGER, _uid);
	ELSE
		RAISE EXCEPTION 'коллекция с таким именем уже существует';
	END IF;
END;
$function$
