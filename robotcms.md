# CloudMinds机器人CMS
- [ ] cms基本类  --Sunny
- [ ] cms基本编辑页面 --Sunny
- [ ] cms分类   --Sunny
- [ ] cms信息类内容访问接口   --Sunny
- [ ] cms系统部署   --Sunny
- [ ] cms音频开发  --JinLong
- [ ] cms音频访问接口  --JinLong
- [ ] cms视频开发  --JinLong
- [ ] cms视频访问接口  --JinLong
- [ ] RCU显示使用的页面适配  --JinLong

# 基本设想
CMS基本类、音频、视频的CURD。数据信息基本包括id、title、describe、keyword、content
音频、视频是id、title、keyword、url

## 基本类数据结构

|序号|字段名|类型|是否必填|解释|
|--:|:--|:--|:--:|:--|
|1|id|int| required|id|
|2|title|string|required|标题|
|3|describe|string|optional|内容摘要|
|4|keyword|string|required|关键词，主要搜索依靠keyword，关键词之间用空格分割|
|5|content|text|详细内容，此处为富文本。为了这次演示，本次都是上传一张图片来解决问题。|
|6|type|int|类型|数据类型，表示此文章属于那个分类。|

## 音频、视频数据结构（这里是两个表，但是结构相同）

|序号|字段名|类型|是否必填|解释|
|--:|:--|:--|:--:|:--|
|1|id|int| required|id|
|2|title|string|required|标题|
|3|describe|string|optional|内容摘要|
|4|keyword|string|required|关键词，主要搜索依靠keyword，关键词之间用空格分割|
|5|url|string|required|保存存储的视频地址|

# 返回接口设想

## 
