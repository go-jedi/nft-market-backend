CREATE OR REPLACE FUNCTION token_create(js json, _uid character varying)
	RETURNS void
	LANGUAGE plpgsql
AS $function$
DECLARE
	_t tokens;
	_c collections;
BEGIN
	SELECT *
	FROM collections
	WHERE collection_uid = js->>'uid_collection'
	INTO _c;

	IF _c.id ISNULL THEN
		RAISE EXCEPTION 'коллекция с таким именем не существует';
	END IF;

	SELECT *
	FROM tokens
	WHERE name = js->>'name'
	INTO _t;

	IF _t.id ISNULL THEN
		INSERT INTO tokens(name, price, author, blockchain, uid_collection, token_uid)
		VALUES(js->>'name', (js->>'price')::NUMERIC, js->>'author', js->>'blockchain', js->>'uid_collection', _uid);

		UPDATE collections SET count = count+1 WHERE collection_uid = js->>'uid_collection';
	ELSE
		RAISE EXCEPTION 'токен с таким именем уже существует';
	END IF;
END;
$function$