SQL_SWIPE="""
select
    actor_user_id as actor_id,
    receiver_user_id as receiver_id,
    actor_user_gender,
    from_unixtime(int(timestamps*1.0/1000000000)) as swipe_timestamp,
    relationship_actor_action as action,
    relationship_actor_state as actor_state,
    relationship_receiver_state as receiver_state,
    get_json_object(relationship,"$.additional.swipeSource") as page,
    get_json_object(relationship,"$.additional.recommendRequestId") as request_id,
    case receiver_last_relationship_time when '-1' then null else from_unixtime(int(receiver_last_relationship_time*1.0/1000000000)) end as receiver_swipe_time,
    custom,
    get_json_object(custom,"$.had_clicked_the_headportrait") as had_clicked_the_headportrait,
    get_json_object(custom,"$.had_clicked_the_personal_id") as had_clicked_the_personal_id,
    get_json_object(custom,"$.had_rolled_the_card_to_end") as had_rolled_the_card_to_end,
    get_json_object(custom,"$.had_clicked_the_record_voice") as had_clicked_the_record_voice,
    get_json_object(custom,"$.had_rolled_the_card") as had_rolled_the_card,
    cast(get_json_object(custom,"$.total_duration_on_card") as bigint) as total_duration_on_card,
    cast(get_json_object(custom,"$.duration_of_reading_description") as bigint) as duration_of_reading_description
from dwd.dwd_ttx_eventlog_user_swipes_i_d
where dt='{0}'
and relationship_actor_action in ('like','dislike','flower')
and get_json_object(relationship,"$.additional.swipeSource") is null
and cast(get_json_object(custom,"$.total_duration_on_card") as bigint) < 60000
"""

df = hc.sql(SQL_SWIPE.format(DATE))

