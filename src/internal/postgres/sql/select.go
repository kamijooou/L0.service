package sql

const SelectSQL = `
select
    order_uid,
    orders.track_number,
    entry,
    deliveries.name,
    phone,
    zip,
    city,
    address,
    region,
    email,
    transaction,
    request_id,
    currency,
    provider,
    amount,
    payment_dt,
    bank,
    delivery_cost,
    goods_total,
    custom_fee,
    chrt_id,
    items.track_number,
    price,
    rid,
    items.name,
    sale,
    size,
    total_price,
    nm_id,
    brand,
    status,
    locale,
    internal_signature,
    customer_id,
    delivery_service,
    shardkey,
    sm_id,
    date_created,
    oof_shard
from orders
inner join deliveries
    on orders.id = deliveries.order_id
inner join payments
    on orders.order_uid = payments.transaction
inner join items
    on orders.track_number = items.track_number
order by orders.order_uid;
`
