CREATE OR REPLACE FUNCTION user_check_token(_tid BIGINT, _tkn character varying)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		COALESCE(uct.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT up.buy_nft
			FROM users_payment up
			WHERE tele_id = _tid
		) ag
	) uct
	INTO _response;

	RETURN _response;
END;
$function$