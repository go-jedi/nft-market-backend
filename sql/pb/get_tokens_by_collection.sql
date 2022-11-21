CREATE OR REPLACE FUNCTION get_tokens_by_collection(_uidc character varying)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		COALESCE(gtbc.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT t.id, c.name as name_collection, c.count, t.name as name_token, t.price as price_token, t.token_uid
			FROM collections c, tokens t
			WHERE c.collection_uid = t.uid_collection
			AND t.uid_collection = _uidc
		) ag
	) gtbc
	INTO _response;

	RETURN _response;
END;
$function$