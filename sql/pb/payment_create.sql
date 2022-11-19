CREATE OR REPLACE FUNCTION payment_create(js json)
    RETURNS void
    LANGUAGE plpgsql
AS $function$
DECLARE
    _p payments;
BEGIN
    SELECT *
    FROM payments
    WHERE name = js->>'name'
    INTO _p;

    IF _p.id ISNULL THEN
        INSERT INTO payments (name, value)
        VALUES(js->>'name', js->>'value');
    ELSE
        RAISE EXCEPTION 'платёжка с таким именем уже существует';
    END IF;
END;
$function$