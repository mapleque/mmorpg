Api接口
========

- target 是一个 Avatar 实例对象

Map地图系统
--------
- get_position:target
    - Position

- get_map:Position
    - list:
        - target
        - Position
        - Distance

- move_direction:target,Direction,Distance
    - list:
        - Distance
        - Position

- move_to:target,Position

Growth成长系统
--------
- add_experience:target,Experience
- add_level:target,Level
- update_max_experience:target,Experience
- update max_level:target,Level
- update_strength:target,value
- update_agility:target,value
- update_intelligence:target,value
- update_haste:target,value
- update_health:target,Health
- update_mana:target,Mana
- update_max_health:target,Health
- update_max_mana:target,Mana
- update_renew_health:target,Health
- update_renew_mana:target,Mana
- run_renew_health:target
- run_renew_mana:target

Skill技能系统
--------
- get_tree:target
    - list:
        - parent_skill
        - Skill
- learn_skill:target,Skill
- forget_skill:target,Skill
- cast_target:target_self,Skill,target
- cast_scale:target_self,Skill,scale

Item物品系统
--------
- bag_items:target
    - list:
        - bag_id
        - Item
- gear_items:target
    - list:
        - gear_id
        - Item
- use_item:target,Item
- add_bag_count:target,count
- add_to_bag:target,Item,bag_id
- bag_to_gear:target,Item,gear_id
- gear_to_bag:target,Item,bag_id
- remove_from_gear:target,Item,gear_id
- remove_from_bag:target,Item,bag_id

