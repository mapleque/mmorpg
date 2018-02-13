Model数据模型
========
- Avatar角色
    - +Display展示数据，可进一步获取模型及贴图等信息
    - +Position位置，可用于进一步获取地图信息
    - +GrowthProperty成长属性值
    - +SkillTree技能树
    - +BufStack增益状态栏
    - -DebufStack减益状态栏
    - +GearStack装备栏
    - +BagStack背包栏

- Display显示
    - name名称

地图系统
--------

- Direction方向
- Distance距离
- Position位置
- Scale位置范围
- Range距离范围

成长系统
--------
- Property成长属性
    - Level等级
    - Experience经验
    - MaxLevel最大等级
    - MaxExperience最大经验
    - Strength力量
    - Agility敏捷
    - Intelligence智力
    - Haste急速
    - Health生命
    - Mana魔法
    - MaxHealth最大生命
    - MaxMana最大魔法
    - RenewHealth恢复生命
    - RenewMana回复魔法

技能系统
--------

- Skill技能
    - +Display
    - +Property影响成长属性变化
    - +Skill后续触发技能
    - Range施放距离
    - Scale影响范围
    - Cast施法时间，0则为瞬发
    - Channel引导时间，0则不需要引导
    - Cooldown冷却时间
    - Duration持续时间
    - Period触发周期，0则一直触发

物品系统
--------

- Item物品
    - +Display
    - type类型：装备，消耗品
    - +Property使用时影响属性变化
    - +Skill使用时施放的技能

