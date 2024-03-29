# 钱包管理

## 钱包功能点
- 批量生成账号(通用账号，专有账号)
- 批量导入账号
- 给账号打标签，在创建任务的时候，可以按照标签进行选择，进行批量操作
- 显示该账号的在各个链的余额
- 网络模块，要维持链的节点网络（比如说根据在eth上操作，要根据链ID（1）去找到对应要连接的节点）
- 批量转账【多对多，一对多】
- 任务模块

账号又分为通用账号和私链账号，


## 表数据结构
### wallet 【钱包表】
- id
- wallet_type
  
  钱包类型【通用型】【专用型】(通用型的钱包可以转化成专用型，但专用性不能转化成通用型)
- type_variables 

  如果钱包是【专用型】钱包，用来保存是哪个链的（比如eth，bitcoin，solana等私钥的钱包）
- pritekey
- publickey
- address
- notes
- create_at

### tag 【标签表】
> 用来给钱包打标签，相当于分组，做批量选取，一个钱包可以有多个标签
- id
- name

### wallet_tag_mapping 【钱包标签映射表】
- wallet_id
- tag_id

### wallet_generic_panel 【钱包通用面板表】
> 用来表示钱包应该显示哪些链的余额（bitcoin、eth、solana、tron等）

### task【任务表】
- id
- task_name
- create_at
- notes
- status
- task_type
    
    表示任务种类（比如批量转账、solana签到）

### task_tag_mapping 【任务标签映射表】
> 表示任务可以选择哪些标签（一个任务可以选取多个标签）
- task_id
- tag_id

### task_progress 【任务快照进度表】
> 用来表示一个任务选取了哪些用户，以及它们任务的完成情况
- task_id
- wallet_id
- completed

  是否完成
- completed_at
- create_at
