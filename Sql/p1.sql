

--查询用户右滑率和被右滑率
select
    e.uid as UID,
    f.like_rate as like_rate,
    e.liked_rate as beliked_rate

from(
        select
            receiver_user_id as uid,
            avg(b.is_like) as liked_rate
        from(
                select
                    receiver_user_id,
                    case when relationship_actor_action="like" then 1 else 0 end as is_like
                from dwd.dwd_ttx_eventlog_user_swipes_i_d
                where dt="2022-07-13"
                  and relationship_actor_action in ('like','dislike'))as b
        group by b.receiver_user_id
    )as e

        left join

    (select
         a.actor_user_id as uid,
         avg(a.is_like) as like_rate
     from(
             select
                 actor_user_id,
                 case when relationship_actor_action = "like" then 1 else 0 end as is_like
             from dwd.dwd_ttx_eventlog_user_swipes_i_d
             where dt="2022-07-13"
               and relationship_actor_action in ('like','dislike')
         )as a
     group by a.actor_user_id)as f

    on e.uid=f.uid


--查询用户右滑列表和被右滑列表，即用户喜欢和被喜欢的用户右哪些
select
    tb2.uid as uid,
    tb1.like_list as like_list,
    tb2.beliked_list as beliked_list


from(select
         a.actor_user_id as uid,
         collect_list(a.receiver_user_id) as like_list

     from
         (
             select
                 actor_user_id,
                 receiver_user_id,
                 relationship_actor_action
             from dwd.dwd_ttx_eventlog_user_swipes_i_d
             where dt="2022-07-13" and relationship_actor_action ="like"
                 limit 100
         )as a
     group by a.actor_user_id)as tb1

        inner join

    (select
         b.receiver_user_id as uid,
         collect_list(b.actor_user_id) as beliked_list

     from
         (
             select
                 actor_user_id,
                 receiver_user_id,
                 relationship_actor_action
             from dwd.dwd_ttx_eventlog_user_swipes_i_d
             where dt="2022-07-13" and relationship_actor_action ="like"
                 limit 100
         )as b
     group by b.receiver_user_id)as tb2

    on tb1.uid=tb2.uid








--个体的信息提取
select
    receiver_user_id,
    relationship_actor_action
from dwd.dwd_ttx_eventlog_user_swipes_i_d
where actor_user_id="16370810" and relationship_actor_action in ('like','dislike')

select
    user_id,
    school

from dwd.dwd_ttx_core_yay_user_infos_a_d
where user_id="16370810"

select
    school,
    is_985
from reco.school_to_feature_v2
where


select *
from
(
    select
    distinct a.receiver_user_id as uid,
    case when a.relationship_actor_action='like' then 1
        when a.relationship_actor_action='dislike' then 0
            end as like_list
    from dwd.dwd_ttx_eventlog_user_swipes_i_d
    where a.actor_user_id="16370810" and a.dt>="2022-07-01"
    and a.relationship_actor_action in ('like','dislike')
)as a
left join
(
    select
     school
    from dwd.dwd_ttx_core_yay_user_infos_a_d
        where dt='2022-07-13'
)as b
on a.receiver_user_id=b.user_id
left join
(
    select school,
               is_985
        from reco.school_to_feature_v2
) as c
on b.school=c.school






--    filter_condition = "gender = 'male' and height >= 175"

select
    actor_user_id,
    swipesN,
    swipesN2,
    round(likesN/swipesN, 2) as likeRate,
    if(likesN2 = 0, 0, round(likesN2/swipesN2, 2)) as likeRate2,
    if(likesN2 = 0, 0, round(likesN2/swipesN2, 2))-round(likesN/swipesN, 2) as delta
from
    (
        select
            actor_user_id,
            count(distinct receiver_user_id) as swipesN,
            count(distinct user_id) as swipesN2,
            count(distinct if(relationship_actor_action = 'like' or relationship_actor_action = 'flower' or relationship_actor_action = 'accept_flower', receiver_user_id, NULL)) as likesN,
            count(distinct if(relationship_actor_action = 'like' or relationship_actor_action = 'flower' or relationship_actor_action = 'accept_flower', user_id, NULL)) as likesN2
        from
            (
                select actor_user_id, receiver_user_id, relationship_actor_action
                from dwd.dwd_ttx_eventlog_user_swipes_i_d
                where dt >= '2022-06-14'
                    limit 10000
            )a
                left join
            (
                select *
                from dwd.dwd_qs_users_info_wide_table_a_d
                where dt = '2022-07-10'
                  and {0} --"gender = 'male' and height >= 175"
            )b
            on a.receiver_user_id = b.user_id
        group by actor_user_id

    )


 --   .format(filter_condition)
re = hc.sql(sql)
re.show(50)


select
    distinct a.receiver_user_id as uid,
    case when a.relationship_actor_action='like' then 1
        when a.relationship_actor_action='dislike' then 0
        end as like_list,
    b.education as education

from dwd.dwd_ttx_eventlog_user_swipes_i_d as a
        left join dwd.dwd_ttx_core_yay_user_infos_a_d as b
        on a.receiver_user_id=b.user_id

where a.actor_user_id="16370810" and a.dt>="2022-07-01"
and a.relationship_actor_action in ('like','dislike')
