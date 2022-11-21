CREATE OR REPLACE FUNCTION get_token_by_uid(_uidt character varying)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		COALESCE(gtbu.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT t.id, t.name, t.price, t.author, t.blockchain, t.token_uid, c.collection_uid as uid_collection, c.name as name_collection
			FROM tokens t, collections c
			WHERE c.collection_uid = t.uid_collection
			AND t.token_uid = _uidt
		) ag
	) gtbu
	INTO _response;

	RETURN _response;
END;
$function$